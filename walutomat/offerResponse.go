package walutomat

type OffersResponse struct {
	Offers []struct {
		Buy       string `json:"buy"`
		Sell      string `json:"sell"`
		BuyOld    string `json:"buy_old"`
		SellOld   string `json:"sell_old"`
		CountSell int    `json:"count_sell"`
		CountBuy  int    `json:"count_buy"`
		Pair      string `json:"pair"`
		Avg       string `json:"avg"`
		AvgOld    string `json:"avg_old"`
	} `json:"offers"`
}
