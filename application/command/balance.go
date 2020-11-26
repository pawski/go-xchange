package command

import (
	"github.com/pawski/go-xchange/configuration"
	"github.com/pawski/go-xchange/logger"
	"github.com/pawski/go-xchange/walutomat"
)

func BalanceExecute() (err error) {

	logger.Get().Info("Trying to get Account balance")

	apiClient := walutomat.NewApiClient(configuration.Get().WalutomatApiHost, configuration.Get().WalutomatApiKey)

	balanceResponse, err := apiClient.GetAccountBalance()

	if err != nil {
		logger.Get().Error(err)

		return
	}

	logger.Get().Printf("%+v\n", balanceResponse)

	return
}
