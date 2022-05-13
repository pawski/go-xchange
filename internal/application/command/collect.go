package command

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/influxdata/influxdb/client/v2"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/streadway/amqp"

	"github.com/pawski/go-xchange/internal/configuration"
	"github.com/pawski/go-xchange/internal/influxdb"
	"github.com/pawski/go-xchange/internal/logger"
	"github.com/pawski/go-xchange/internal/rabbitmq"
	"github.com/pawski/go-xchange/internal/walutomat"
)

func CollectExecute() error {
	// TODO Normalise bootstrap
	rabbitConnection, err := rabbitmq.NewConnection(configuration.Get().RabbitMqUrl)
	if err != nil {
		return errors.Wrap(err, "Cannot connect to rabbitMQ")
	}

	channel, err := rabbitConnection.Channel()
	if err != nil {
		return err
	}
	defer func() {
		if err := channel.Close(); err != nil {
			logger.Get().Error(err)
		}
	}()

	rabbitQueue, err := channel.QueueDeclare(
		"rates_queue", // name
		true,          // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		return err
	}

	reg := prometheus.NewRegistry()
	reg.MustRegister(collectors.NewBuildInfoCollector())
	reg.MustRegister(collectors.NewGoCollector())

	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.HandlerFor(
		reg,
		promhttp.HandlerOpts{},
	))

	srv := &http.Server{
		Addr:    "localhost:8000",
		Handler: mux,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != http.ErrServerClosed {
			fmt.Println()
		}
	}()

	logger.Get().Infof("Running...")

	wait := make(chan os.Signal, 1)
	logger.Get().Info("Setting up")

	wg := sync.WaitGroup{}
	shouldRun := true
	go rabbitmq.ConsumeFromQueue(
		channel,
		rabbitQueue.Name,
		func(deliveries <-chan amqp.Delivery) {
			wg.Add(1)
			defer wg.Done()
			for d := range deliveries {
				logger.Get().Info("Received a message")
				handleMessageBody(d.Body)
				err := d.Ack(false)
				if err != nil {
					logger.Get().Error(err)
				}
				if shouldRun == false {
					return
				}
			}
		})

	logger.Get().Infof(" [*] Waiting for messages. To exit press CTRL+C")
	<-wait
	shouldRun = false
	wg.Wait()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Get().Fatalf("Server Shutdown Failed:%+v", err)
	}

	return nil
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

	for _, offer := range walutomat.Convert(response) {
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
