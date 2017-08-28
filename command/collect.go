package command

import (
	"log"
	"time"
	"github.com/pawski/go-xchange/procctl"
	"github.com/pawski/go-xchange/walutomat"
	"github.com/pawski/go-xchange/http"
	"github.com/influxdata/influxdb/client/v2"
	"fmt"
	"github.com/pawski/go-xchange/configuration"
)

func CollectExecute() (err error) {
	go procctl.RegisterSigTerm()

	configuration := configuration.Configuration()

	url := configuration.WalutomatUrl
	ticker := time.NewTicker(time.Second * 60)

	// Create a new HTTPClient
	influxdb, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     configuration.InfluxDbHost,
		Username: "",
		Password: "",
	})

	if err != nil {
		log.Fatal(err)
	}

	go func() {
		log.Println("Start at", time.Now())
		handleResponse(http.GetUrl(url), influxdb, configuration.InfluxDbDatabase)
		for t := range ticker.C {
			log.Println("Start at", t)
			handleResponse(http.GetUrl(url), influxdb, configuration.InfluxDbDatabase)
		}
	}()

	block := make(chan bool, 1)
	<-block

	return
}

func handleResponse(response []byte, influxdb client.Client, db string) {

	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  db,
		Precision: "s",
	})
	if err != nil {
		http.FlushBufferToFile()
		log.Fatal(err)
	}

	for index, offer := range walutomat.Convert(response) {
		fmt.Println(index, offer)

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
			log.Fatal(err)
		}
		bp.AddPoint(pt)
	}

	// Write the batch
	if err := influxdb.Write(bp); err != nil {
		http.FlushBufferToFile()
		log.Fatal(err)
	}
}
