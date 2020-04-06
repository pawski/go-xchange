package command

import (
	"github.com/pawski/go-xchange/procctl"
	"github.com/pawski/go-xchange/configuration"
	"time"
	"github.com/pawski/go-xchange/logger"
	"github.com/pawski/go-xchange/http"
	"github.com/pawski/go-xchange/rabbitmq"
)

func FetchExecute() (err error) {
	go procctl.RegisterSigTerm()
	rabbitmq.Connect()

	url := configuration.Get().WalutomatUrl
	interval := time.Second * time.Duration(configuration.Get().CollectUpdateInterval)
	ticker := time.NewTicker(interval)

	go func() {
		logger.Get().Println("Start at", time.Now())
		handleResponseMessage(http.GetUrl(url))
		for t := range ticker.C {
			logger.Get().Println("Start at", t)
			response := http.GetUrl(url)
			handleResponseMessage(response)
		}
	}()

	block := make(chan bool, 1)
	<-block

	return
}

func handleResponseMessage(response []byte) {

	rabbitmq.PublishToQueue(response)

	return
}
