package command

import (
	"fmt"
	"github.com/pawski/go-xchange/configuration"
	"github.com/pawski/go-xchange/logger"
	"github.com/pawski/go-xchange/walutomat"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var shouldRun = true

func RunExecute() {
	registerSigTerm()
	logger.Get().Info("Trying to get Direct Rates balance")

	apiClient := walutomat.NewApiClient(configuration.Get().WalutomatApiHost, configuration.Get().WalutomatApiKey)

	for shouldRun {
		rateResponse, _ := apiClient.GetDirectRates(walutomat.EURPLN)
		logger.Get().Infof("%v: Buy %v, Sell %v", rateResponse.Result.CurrencyPair, rateResponse.Result.BuyRate, rateResponse.Result.SellRate)
		time.Sleep(time.Second * 2)
	}
}

func registerSigTerm() {
	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt)
	signal.Notify(s, syscall.SIGTERM)
	go func() {
		<-s
		fmt.Println("Shutting down gracefully.")
		shouldRun = false
	}()
}
