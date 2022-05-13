package command

import (
	"os"
	"sync"
	"time"

	"github.com/pawski/go-xchange/internal/configuration"
	"github.com/pawski/go-xchange/internal/http"
	"github.com/pawski/go-xchange/internal/logger"
	"github.com/pawski/go-xchange/internal/rabbitmq"
)

func FetchExecute() error {
	rabbitConnection, err := rabbitmq.NewConnection(configuration.Get().RabbitMqUrl)
	if err != nil {
		return err
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

	wait := make(chan os.Signal, 1)
	url := configuration.Get().WalutomatUrl
	interval := time.Second * time.Duration(configuration.Get().CollectUpdateInterval)
	ticker := time.NewTicker(interval)

	wg := sync.WaitGroup{}
	go func() {
		wg.Add(1)
		defer wg.Done()
		for t := range ticker.C {
			logger.Get().Println("Start at", t)
			err := rabbitmq.Publish(channel, rabbitQueue.Name, http.GetUrl(url))
			if err != nil {
				logger.Get().Error(err)
			}
		}
	}()

	<-wait
	ticker.Stop()
	wg.Wait()

	return nil
}
