package walutomat

type OffersResponse struct {
	Offers []struct {
		Buy       float32 `json:"buy"`
		Sell      float32 `json:"sell"`
		BuyOld    float32 `json:"buy_old"`
		SellOld   float32 `json:"sell_old"`
		CountSell int     `json:"count_sell"`
		CountBuy  int     `json:"count_buy"`
		Pair      string  `json:"pair"`
		Avg       float32 `json:"avg"`
		AvgOld    float32 `json:"avg_old"`
	} `json:"offers"`
}
