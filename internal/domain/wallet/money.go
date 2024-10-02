package wallet

import (
	"fmt"
	"math"
	"math/big"

	"github.com/pkg/errors"
)

const MoneyPrecision = 2

type Money struct {
	currency  Currency
	amount    int64
	precision int
}

func RecreateAmountFromCents(amount int64, currency Currency, precision int) Money {
	return Money{
		amount:    amount,
		currency:  currency,
		precision: precision,
	}
}

func NewAmount(amount float64, currency Currency) (Money, error) {
	m := withDefaultPrecision()
	m.currency = currency

	fromAmount := math.Round(amount * math.Pow10(m.precision))
	toAmount := fromAmount / math.Pow10(m.precision)

	if amount-toAmount != 0 {
		return m, errors.Errorf("provided amount %f requires greater precision than default %d", amount, m.precision)
	}
	m.amount = int64(fromAmount)
	return m, nil
}

func withDefaultPrecision() Money {
	return Money{
		precision: MoneyPrecision,
	}
}

func (m Money) ExchangeTo(offer ExchangeOffer) (Money, error) {

	if m.currency.AsString() != offer.Pair().From() && m.currency.AsString() != offer.Pair().To() {
		return Money{}, errors.Errorf("offer doesn't include exchange for %s", m.currency.AsString())
	}

	var amount float64
	var c Currency

	if m.currency.AsString() == offer.Pair().To() {
		c = Currency(offer.Pair().From())
		amount = m.AsFloat64() / offer.SellRatio().Rate().AsFloat64()
	} else if m.currency.AsString() == offer.Pair().From() {
		c = Currency(offer.Pair().To())
		amount = m.AsFloat64() * offer.BuyRatio().Rate().AsFloat64()
	} else {
		panic("unsupported money exchange")
	}

	pow := math.Pow10(m.precision)

	return NewAmount(math.Round(amount*pow)/pow, c)
}

func (m Money) Currency() Currency {
	return m.currency
}

func (m Money) AsString() string {
	return fmt.Sprintf("%s %s", big.NewFloat(m.AsFloat64()).Text('f', m.precision), m.currency)
}

func (m Money) AsFloat64() float64 {
	return float64(m.amount) / math.Pow10(m.precision)
}

func (m Money) Extract(nm Money) (Money, error) {
	if m.currency != nm.currency || m.precision != nm.precision {
		return Money{}, errors.New("money has to be in same currency and with same precision")
	}

	return NewAmount(float64(m.amount-nm.amount)/math.Pow10(m.precision), m.currency)
}

func (m Money) Percent(percent float64) (Money, error) {
	pow := math.Pow10(MoneyPrecision)

	return NewAmount(math.Round(m.AsFloat64()*percent*pow)/pow, m.Currency())
}
