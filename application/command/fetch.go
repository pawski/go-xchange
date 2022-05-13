package command

import (
	"sync"
	"time"

	"github.com/pawski/go-xchange/configuration"
	"github.com/pawski/go-xchange/http"
	"github.com/pawski/go-xchange/logger"
	"github.com/pawski/go-xchange/procctl"
	"github.com/pawski/go-xchange/rabbitmq"
)

func FetchExecute() (err error) {
	var shouldRun = true
	procctl.RegisterSigTerm(&shouldRun)

	url := configuration.Get().WalutomatUrl
	interval := time.Second * time.Duration(configuration.Get().CollectUpdateInterval)
	ticker := time.NewTicker(interval)

	wg := sync.WaitGroup{}
	go func() {
		wg.Add(1)
		defer wg.Done()
		logger.Get().Println("Start at", time.Now())
		handleResponseMessage(http.GetUrl(url))
		for t := range ticker.C {
			logger.Get().Println("Start at", t)
			response := http.GetUrl(url)
			handleResponseMessage(response)
			if shouldRun == false {
				ticker.Stop()
			}
		}
	}()

	wg.Wait()

	return
}

func handleResponseMessage(response []byte) {

	rabbitmq.PublishToQueue(response)

	return
}
