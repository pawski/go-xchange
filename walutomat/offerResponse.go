package walutomat

import "encoding/json"

type OffersResponse struct {
	Offers []struct {
		Buy       json.Number `json:"buy"`
		Sell      json.Number `json:"sell"`
		BuyOld    json.Number `json:"buy_old"`
		SellOld   json.Number `json:"sell_old"`
		CountSell int    `json:"count_sell"`
		CountBuy  int    `json:"count_buy"`
		Pair      string `json:"pair"`
		Avg       json.Number `json:"avg"`
		AvgOld    json.Number `json:"avg_old"`
	} `json:"offers"`
}
