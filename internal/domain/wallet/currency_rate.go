package wallet

const RatePrecision = 4

type CurrencyRate struct {
	pair Pair
	rate Rate
}

func RecreateCurrencyRate(pair Pair, ratio Rate) CurrencyRate {
	return CurrencyRate{
		pair: pair,
		rate: ratio,
	}
}

func (c CurrencyRate) Pair() Pair {
	return c.pair
}

func (c CurrencyRate) Rate() Rate {
	return c.rate
}
