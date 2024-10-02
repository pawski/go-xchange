package wallet

type ExchangeOffer struct {
	pair     Pair
	sellRate CurrencyRate
	buyRate  CurrencyRate
}

func RecreateExchangeOfferFromValue(pair Pair, sellRate CurrencyRate, buyRate CurrencyRate) ExchangeOffer {
	return ExchangeOffer{
		pair:     pair,
		sellRate: sellRate,
		buyRate:  buyRate,
	}
}

func (o ExchangeOffer) Pair() Pair {
	return o.pair
}

func (o ExchangeOffer) SellRatio() CurrencyRate {
	return o.sellRate
}

func (o ExchangeOffer) BuyRatio() CurrencyRate {
	return o.buyRate
}
