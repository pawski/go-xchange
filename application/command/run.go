package command

import (
	"github.com/pawski/go-xchange/configuration"
	"github.com/pawski/go-xchange/logger"
	"github.com/pawski/go-xchange/procctl"
	"github.com/pawski/go-xchange/walutomat"
	"time"
)


func RunExecute() {
	var shouldRun = true
	procctl.RegisterSigTerm(&shouldRun)
	logger.Get().Info("Trying to get Direct Rates balance")

	apiClient := walutomat.NewApiClient(configuration.Get().WalutomatApiHost, configuration.Get().WalutomatApiKey)

	for shouldRun {
		rateResponse, _ := apiClient.GetDirectRates(walutomat.EURPLN)
		logger.Get().Infof("%v: Buy %v, Sell %v", rateResponse.Result.CurrencyPair, rateResponse.Result.BuyRate, rateResponse.Result.SellRate)
		time.Sleep(time.Second * 2)
	}
}
