package wallet

import (
	"testing"

	"github.com/onsi/gomega"
)

func TestMoneyAmountFromFloat64(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	t.Run("33.9876", func(t *testing.T) {
		m, err := NewAmount(33.98, EUR)
		g.Expect(err).To(gomega.BeNil())

		g.Expect(m.precision).To(gomega.Equal(2))
		g.Expect(m.amount).To(gomega.Equal(int64(3398)))
		g.Expect(m.currency).To(gomega.Equal(EUR))
	})

	t.Run("3042.1838", func(t *testing.T) {
		m, err := NewAmount(3042.18, EUR)
		g.Expect(err).To(gomega.BeNil())

		g.Expect(m.precision).To(gomega.Equal(2))
		g.Expect(m.amount).To(gomega.Equal(int64(304218)))
		g.Expect(m.currency).To(gomega.Equal(EUR))
	})

	t.Run("1", func(t *testing.T) {
		m, err := NewAmount(1, EUR)
		g.Expect(err).To(gomega.BeNil())

		g.Expect(m.precision).To(gomega.Equal(2))
		g.Expect(m.amount).To(gomega.Equal(int64(100)))
		g.Expect(m.currency).To(gomega.Equal(EUR))
	})

	t.Run("1.01", func(t *testing.T) {
		m, err := NewAmount(1.01, EUR)
		g.Expect(err).To(gomega.BeNil())

		g.Expect(m.precision).To(gomega.Equal(2))
		g.Expect(m.amount).To(gomega.Equal(int64(101)))
		g.Expect(m.currency).To(gomega.Equal(EUR))
	})
	t.Run("1.04", func(t *testing.T) {
		m, err := NewAmount(1.04, EUR)
		g.Expect(err).To(gomega.BeNil())

		g.Expect(m.precision).To(gomega.Equal(2))
		g.Expect(m.amount).To(gomega.Equal(int64(104)))
		g.Expect(m.currency).To(gomega.Equal(EUR))
	})
	t.Run("1.05", func(t *testing.T) {
		m, err := NewAmount(1.05, EUR)
		g.Expect(err).To(gomega.BeNil())

		g.Expect(m.precision).To(gomega.Equal(2))
		g.Expect(m.amount).To(gomega.Equal(int64(105)))
		g.Expect(m.currency).To(gomega.Equal(EUR))
	})

	t.Run("1.06", func(t *testing.T) {
		m, err := NewAmount(1.06, EUR)
		g.Expect(err).To(gomega.BeNil())

		g.Expect(m.precision).To(gomega.Equal(2))
		g.Expect(m.amount).To(gomega.Equal(int64(106)))
		g.Expect(m.currency).To(gomega.Equal(EUR))
	})
	t.Run("1.09", func(t *testing.T) {
		m, err := NewAmount(1.09, EUR)
		g.Expect(err).To(gomega.BeNil())

		g.Expect(m.precision).To(gomega.Equal(2))
		g.Expect(m.amount).To(gomega.Equal(int64(109)))
		g.Expect(m.currency).To(gomega.Equal(EUR))
	})

	t.Run("999999.01", func(t *testing.T) {
		m, err := NewAmount(999999.01, EUR)
		g.Expect(err).To(gomega.BeNil())

		g.Expect(m.precision).To(gomega.Equal(2))
		g.Expect(m.amount).To(gomega.Equal(int64(99999901)))
		g.Expect(m.currency).To(gomega.Equal(EUR))
	})
}

func TestMoneyExchange(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	t.Run("EUR to PLN", func(t *testing.T) {
		deposit, err := NewAmount(33.98, EUR)
		offer := RecreateExchangeOfferFromValue(EURPLN, RecreateCurrencyRate(EURPLN, Rate(1.1)), RecreateCurrencyRate(EURPLN, Rate(4.4868)))
		g.Expect(err).To(gomega.BeNil())

		money, err := deposit.ExchangeTo(offer)
		g.Expect(err).To(gomega.BeNil())

		g.Expect(money.precision).To(gomega.Equal(2))
		g.Expect(money.amount).To(gomega.Equal(int64(15246)))
		g.Expect(money.currency).To(gomega.Equal(PLN))
		g.Expect(money.AsString()).To(gomega.Equal("152.46 PLN"))
	})

	t.Run("PLN to EUR", func(t *testing.T) {
		deposit, err := NewAmount(93264, PLN)
		g.Expect(err).To(gomega.BeNil())

		offer := RecreateExchangeOfferFromValue(EURPLN, RecreateCurrencyRate(EURPLN, Rate(4.6110)), RecreateCurrencyRate(EURPLN, Rate(1.1)))
		money, err := deposit.ExchangeTo(offer)
		g.Expect(err).To(gomega.BeNil())

		g.Expect(money.precision).To(gomega.Equal(2))
		g.Expect(money.amount).To(gomega.Equal(int64(2022642)))
		g.Expect(money.currency).To(gomega.Equal(EUR))
		g.Expect(money.AsString()).To(gomega.Equal("20226.42 EUR"))
	})

	t.Run("EUR to PLN - same amount", func(t *testing.T) {
		deposit, err := NewAmount(10000, EUR)
		g.Expect(err).To(gomega.BeNil())

		offer := RecreateExchangeOfferFromValue(EURPLN, RecreateCurrencyRate(EURPLN, Rate(4.5001)), RecreateCurrencyRate(EURPLN, Rate(4.5001)))
		money, err := deposit.ExchangeTo(offer)
		g.Expect(err).To(gomega.BeNil())

		g.Expect(money.precision).To(gomega.Equal(2))
		g.Expect(money.amount).To(gomega.Equal(int64(4500100)))
		g.Expect(money.currency).To(gomega.Equal(PLN))
		g.Expect(money.AsString()).To(gomega.Equal("45001.00 PLN"))
	})

	t.Run("EUR to PLN", func(t *testing.T) {
		deposit, err := NewAmount(10000, EUR)
		g.Expect(err).To(gomega.BeNil())

		offer := RecreateExchangeOfferFromValue(EURPLN, RecreateCurrencyRate(EURPLN, Rate(4.2000)), RecreateCurrencyRate(EURPLN, Rate(4.8000)))
		exchangedMoney, err := deposit.ExchangeTo(offer)
		g.Expect(err).To(gomega.BeNil())

		expectedMoney := RecreateAmountFromCents(4800000, PLN, MoneyPrecision)
		g.Expect(exchangedMoney).To(gomega.Equal(expectedMoney))
		g.Expect(exchangedMoney.AsString()).To(gomega.Equal("48000.00 PLN"))
	})
}
