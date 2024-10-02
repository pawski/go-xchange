package wallet

type Equivalent struct {
	pair  Pair
	ratio Rate
	money Money
}

func RecreateEquivalent(pair Pair, ratio Rate, money Money) Equivalent {
	return Equivalent{
		pair:  pair,
		ratio: ratio,
		money: money,
	}
}

func (e Equivalent) Pair() Pair {
	return e.pair
}

func (e Equivalent) Ratio() Rate {
	return e.ratio
}

func (e Equivalent) Money() Money {
	return e.money
}
