package account

import (
	"encoding/json"
	"github.com/pawski/go-xchange/walutomat/resources"
)

const ResourcePath = "/api/v2.0.0/account/balances"

type BalanceResponse struct {
	Success bool `json:"success"`
	Errors  []resources.Error
	Result  []struct {
		Currency         string      `json:"currency"`
		BalanceTotal     json.Number `json:"balanceTotal"`
		BalanceAvailable json.Number `json:"balanceAvailable"`
		BalanceReserved  json.Number `json:"balanceReserved"`
	}
}
