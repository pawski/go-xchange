package command

import (
	"sync"
	"time"

	"github.com/influxdata/influxdb/client/v2"
	"github.com/streadway/amqp"

	"github.com/pawski/go-xchange/configuration"
	"github.com/pawski/go-xchange/influxdb"
	"github.com/pawski/go-xchange/logger"
	"github.com/pawski/go-xchange/procctl"
	"github.com/pawski/go-xchange/rabbitmq"
	"github.com/pawski/go-xchange/walutomat"
)

func CollectExecute() (err error) {
	var shouldRun = true
	procctl.RegisterSigTerm(&shouldRun)

	wg := sync.WaitGroup{}

	logger.Get().Info("Setting up")

	go rabbitmq.ConsumeFromQueue(func(deliveries <-chan amqp.Delivery) {
		wg.Add(1)
		defer wg.Done()
		for d := range deliveries {
			logger.Get().Info("Received a message")
			handleMessageBody(d.Body)
			d.Ack(false)
			if shouldRun == false {
				return
			}
		}
	})

	logger.Get().Infof(" [*] Waiting for messages. To exit press CTRL+C")
	wg.Wait()

	return
}

func handleMessageBody(response []byte) {

	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  configuration.Get().InfluxDbDatabase,
		Precision: "s",
	})

	if err != nil {
		logger.Get().Fatal(err)
	}

	for index, offer := range walutomat.Convert(response) {
		logger.Get().Info(index, offer)

		// Create a point and add to batch
		tags := map[string]string{"pair": offer.Pair}
		fields := map[string]interface{}{
			"Buy":       offer.Buy,
			"BuyOld":    offer.BuyOld,
			"CountBuy":  offer.CountBuy,
			"Sell":      offer.Sell,
			"SellOld":   offer.SellOld,
			"CountSell": offer.CountSell,
			"Avg":       offer.Avg,
			"AvgOld":    offer.AvgOld,
		}

		pt, err := client.NewPoint("offers", tags, fields, time.Now())
		if err != nil {
			logger.Get().Fatal(err)
		}
		bp.AddPoint(pt)
	}

	// Write the batch
	if err := influxdb.Get().Write(bp); err != nil {
		panic(err)
	}
}
