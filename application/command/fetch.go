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

	logger.Get().Infof("Fetching %v", url)

	go func() {
		handleResponseMessage(http.GetUrl(url))
		for t := range ticker.C {
			if !shouldRun {
				break
			}

			logger.Get().Infof("Starting at %v", t.Format("2006-01-02 15:04:05"))
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
