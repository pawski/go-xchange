package wallet

type Rate float64

func (r Rate) AsFloat64() float64 {
	return float64(r)
}
