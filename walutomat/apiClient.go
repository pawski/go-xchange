package walutomat

import (
	"encoding/json"
	"github.com/pawski/go-xchange/logger"
	"github.com/pawski/go-xchange/walutomat/resources/account"
	"io/ioutil"
	"net/http"
)

type ApiClient struct {
	host       string
	apiKey     string
}

func NewApiClient(host string, apiKey string) *ApiClient {
	return &ApiClient{host, apiKey}
}

func (api *ApiClient) GetAccountBalance() (*account.BalanceResponse, error) {

 	httpClient := &http.Client{}
	httpRequest, httpBuildError := http.NewRequest("GET", api.host + account.ResourcePath, nil)

	var balanceResponse *account.BalanceResponse

	if httpBuildError != nil {
		logger.Get().Error(httpBuildError.Error())

		return balanceResponse, httpBuildError
	}

	httpRequest.Header.Add("X-API-Key", api.apiKey)

	response, httpError := httpClient.Do(httpRequest)

	if httpError != nil {
		logger.Get().Error(httpError)

		return balanceResponse, httpBuildError
	}

	defer response.Body.Close()
	body, responseReadError := ioutil.ReadAll(response.Body)

	if nil != responseReadError {
		logger.Get().Error(responseReadError)

		return balanceResponse, responseReadError
	}

	json.Unmarshal(body, &balanceResponse)

	logger.Get().Debugf("%d bytes", len(body))
	logger.Get().Debugf("Account balance: %s", body)

	return balanceResponse, nil
}
