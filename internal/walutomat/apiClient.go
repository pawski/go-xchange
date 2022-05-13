package walutomat

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pawski/go-xchange/internal/logger"
	"github.com/pawski/go-xchange/internal/walutomat/resources/account"
	"github.com/pawski/go-xchange/internal/walutomat/resources/direct"
)

type ApiClient struct {
	host   string
	apiKey string
}

func NewApiClient(host string, apiKey string) *ApiClient {
	return &ApiClient{host, apiKey}
}

func (api *ApiClient) GetDirectRates(pair CurrencyPair) (*direct.RatesResponse, error) {
	httpClient := &http.Client{}
	httpRequest, httpBuildError := http.NewRequest("GET", api.host+direct.ResourcePath, nil)

	var ratesResponse *direct.RatesResponse

	if httpBuildError != nil {
		logger.Get().Error(httpBuildError.Error())

		return ratesResponse, httpBuildError
	}

	httpRequest.Header.Add("X-API-Key", api.apiKey)

	httpQuery := httpRequest.URL.Query()
	httpQuery.Add("currencyPair", fmt.Sprint(pair))
	httpRequest.URL.RawQuery = httpQuery.Encode()

	response, httpError := httpClient.Do(httpRequest)

	if httpError != nil {
		logger.Get().Error(httpError)

		return ratesResponse, httpBuildError
	}

	defer response.Body.Close()
	body, responseReadError := ioutil.ReadAll(response.Body)

	if nil != responseReadError {
		logger.Get().Error(responseReadError)

		return ratesResponse, responseReadError
	}

	json.Unmarshal(body, &ratesResponse)

	logger.Get().Debugf("%d bytes", len(body))
	logger.Get().Debugf("Account balance: %s", body)

	return ratesResponse, nil
}

func (api *ApiClient) GetAccountBalance() (*account.BalanceResponse, error) {

	httpClient := &http.Client{}
	httpRequest, httpBuildError := http.NewRequest("GET", api.host+account.ResourcePath, nil)

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
