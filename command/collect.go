package command

import (
	"time"
	"github.com/pawski/go-xchange/procctl"
	"github.com/pawski/go-xchange/walutomat"
	"github.com/pawski/go-xchange/http"
	"github.com/influxdata/influxdb/client/v2"
	"github.com/pawski/go-xchange/configuration"
	"github.com/pawski/go-xchange/logger"
	"github.com/pawski/go-xchange/influxdb"
)

func CollectExecute() (err error) {
	go procctl.RegisterSigTerm()

	url := configuration.Get().WalutomatUrl
	interval := time.Second * time.Duration(configuration.Get().CollectUpdateInterval)
	ticker := time.NewTicker(interval)

	go func() {
		logger.Get().Println("Start at", time.Now())
		handleResponse(http.GetUrl(url))
		for t := range ticker.C {
			logger.Get().Println("Start at", t)
			response := http.GetUrl(url)
			handleResponse(response)
		}
	}()

	block := make(chan bool, 1)
	<-block

	return
}

func handleResponse(response []byte) {

	defer func() {
		if err := recover(); err != nil {
			logger.Get().Error("Sending data to influx failed:", err)
			http.FlushBufferToFile()
		}
	}()

	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database: configuration.Get().InfluxDbDatabase,
		Precision: "s",
	})

	if err != nil {
		http.FlushBufferToFile()
		logger.Get().Fatal(err)
	}

	for index, offer := range walutomat.Convert(response) {
		logger.Get().Info(index, offer)

		// Create a point and add to batch
		tags := map[string]string{"pair": offer.Pair}
		fields := map[string]interface{}{
			"Buy": offer.Buy,
			"BuyOld": offer.BuyOld,
			"CountBuy": offer.CountBuy,
			"Sell" : offer.Sell,
			"SellOld" : offer.SellOld,
			"CountSell" : offer.CountSell,
			"Avg" : offer.Avg,
			"AvgOld" : offer.AvgOld,
		}

		pt, err := client.NewPoint("offers", tags, fields, time.Now())
		if err != nil {
			http.FlushBufferToFile()
			logger.Get().Fatal(err)
		}
		bp.AddPoint(pt)
	}

	// Write the batch
	if err := influxdb.Get().Write(bp); err != nil {
		panic(err)
	}
}
