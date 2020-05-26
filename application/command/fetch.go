package command

import (
	"github.com/pawski/go-xchange/configuration"
	"github.com/pawski/go-xchange/http"
	"github.com/pawski/go-xchange/logger"
	"github.com/pawski/go-xchange/procctl"
	"github.com/pawski/go-xchange/rabbitmq"
	"sync"
	"time"
)

func FetchExecute() (err error) {
	var shouldRun = true
	procctl.RegisterSigTerm(&shouldRun)
	rabbitmq.Connect()

	url := configuration.Get().WalutomatUrl
	interval := time.Second * time.Duration(configuration.Get().CollectUpdateInterval)
	ticker := time.NewTicker(interval)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		logger.Get().Println("Start at", time.Now())
		handleResponseMessage(http.GetUrl(url))
		for t := range ticker.C {
			if !shouldRun {
				break
			}

			logger.Get().Println("Start at", t)
			response := http.GetUrl(url)
			handleResponseMessage(response)
		}

		wg.Done()
	}()

	wg.Wait()

	return
}

func handleResponseMessage(response []byte) {

	rabbitmq.PublishToQueue(response)

	return
}
