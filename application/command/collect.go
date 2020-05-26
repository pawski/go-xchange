package command

import (
	"github.com/influxdata/influxdb/client/v2"
	"github.com/pawski/go-xchange/configuration"
	"github.com/pawski/go-xchange/influxdb"
	"github.com/pawski/go-xchange/logger"
	"github.com/pawski/go-xchange/rabbitmq"
	"github.com/pawski/go-xchange/walutomat"
	"github.com/streadway/amqp"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func CollectExecute() (err error) {
	wg := sync.WaitGroup{}

	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt)
	signal.Notify(s, syscall.SIGTERM)

	logger.Get().Info("Setting up")

	wg.Add(1)
	go rabbitmq.ConsumeFromQueue(func(deliveries <-chan amqp.Delivery) {
		for {
			select {
			case d := <-deliveries:
				logger.Get().Info("Received a message")
				handleMessageBody(d.Body)
				d.Ack(false)
			case <-s:
				logger.Get().Info("Shutting down...")
				wg.Done()
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
		Database: configuration.Get().InfluxDbDatabase,
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
			"Buy": offer.Buy,
			"BuyOld": offer.BuyOld,
			"CountBuy": offer.CountBuy,
			"Sell" : offer.Sell,
			"SellOld" : offer.SellOld,
			"CountSell" : offer.CountSell,
			"Avg" : offer.Avg,
			"AvgOld" : offer.AvgOld,
		}

		logger.Get().Debug(fields)
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
