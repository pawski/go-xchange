package command

import (
	"fmt"
	"os"
	"time"

	"github.com/pawski/go-xchange/internal/configuration"
	"github.com/pawski/go-xchange/internal/logger"
	walutomat2 "github.com/pawski/go-xchange/internal/walutomat"
)

func RunExecute() {
	var shouldRun = true
	s := make(chan os.Signal, 1)
	go func(shouldRun *bool) {
		<-s
		fmt.Println("Finishing remaining tasks...")
		*shouldRun = false
	}(&shouldRun)

	logger.Get().Info("Trying to get Direct Rates balance")

	apiClient := walutomat2.NewApiClient(configuration.Get().WalutomatApiHost, configuration.Get().WalutomatApiKey)

	for shouldRun {
		rateResponse, _ := apiClient.GetDirectRates(walutomat2.EURPLN)
		logger.Get().Infof("%v: Buy %v, Sell %v", rateResponse.Result.CurrencyPair, rateResponse.Result.BuyRate, rateResponse.Result.SellRate)
		time.Sleep(time.Second * 2)
	}
}
