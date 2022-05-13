package command

import (
	"github.com/pawski/go-xchange/internal/configuration"
	"github.com/pawski/go-xchange/internal/logger"
	walutomat2 "github.com/pawski/go-xchange/internal/walutomat"
)

func DirectRatesExecute() (err error) {

	logger.Get().Info("Trying to get Direct Rates balance")

	apiClient := walutomat2.NewApiClient(configuration.Get().WalutomatApiHost, configuration.Get().WalutomatApiKey)

	ratesResponse, err := apiClient.GetDirectRates(walutomat2.EURPLN)

	if err != nil {
		logger.Get().Error(err)

		return
	}

	logger.Get().Printf("%+v\n", ratesResponse)

	return
}
