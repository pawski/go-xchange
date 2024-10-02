package walutomat_v1

type Offer struct {
	Pair      string
	Buy       float64
	BuyOld    float64
	CountBuy  int
	Sell      float64
	SellOld   float64
	CountSell int
	Avg       float64
	AvgOld    float64
}
