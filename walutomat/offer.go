package walutomat

type Offer struct {
	Pair      string
	Buy       float32
	BuyOld    float32
	CountBuy  int
	Sell      float32
	SellOld   float32
	CountSell int
	Avg       float32
	AvgOld    float32
}
