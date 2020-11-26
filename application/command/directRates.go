package command

import (
	"github.com/pawski/go-xchange/configuration"
	"github.com/pawski/go-xchange/logger"
	"github.com/pawski/go-xchange/walutomat"
)

func DirectRatesExecute() (err error) {

	logger.Get().Info("Trying to get Direct Rates balance")

	apiClient := walutomat.NewApiClient(configuration.Get().WalutomatApiHost, configuration.Get().WalutomatApiKey)

	ratesResponse, err := apiClient.GetDirectRates(walutomat.EURPLN)

	if err != nil {
		logger.Get().Error(err)

		return
	}

	logger.Get().Printf("%+v\n", ratesResponse)

	return
}
