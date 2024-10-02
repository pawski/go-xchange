package wallet

type Currency string

const (
	EUR Currency = "EUR"
	PLN Currency = "PLN"
)

func (c Currency) AsString() string {
	return string(c)
}
