package direct

import (
	"encoding/json"
	"github.com/pawski/go-xchange/walutomat/resources"
)

const ResourcePath = "/api/v2.0.0/direct_fx/rates"

type RatesResponse struct {
	Success bool `json:"success"`
	Errors  []resources.Error
	Result  struct {
		Ts           string      `json:"ts"`
		CurrencyPair string      `json:"currencyPair"`
		BuyRate      json.Number `json:"buyRate"`
		SellRate     json.Number `json:"sellRate"`
	}
}
