package wallet

import (
	"testing"

	"github.com/onsi/gomega"
)

func TestCanRecreateDeposit(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	t.Run("PLN with EUR equivalent", func(t *testing.T) {
		money, err := NewAmount(9232.80, PLN)
		g.Expect(err).To(gomega.BeNil())

		deposit := RecreateDeposit(
			money,
			RecreateEquivalent(
				EURPLN, Rate(4.6164),
				RecreateAmountFromCents(200000, EUR, MoneyPrecision),
			),
		)
		g.Expect(deposit.Money().AsString()).To(gomega.Equal("9232.80 PLN"))
		g.Expect(deposit.Equivalent().Money().AsString()).To(gomega.Equal("2000.00 EUR"))
		g.Expect(deposit.Equivalent().Ratio().AsFloat64()).To(gomega.Equal(4.6164))
		g.Expect(deposit.Equivalent().Pair().AsString()).To(gomega.Equal(string(EURPLN)))

	})

	t.Run("EUR with PLN equivalent", func(t *testing.T) {
		money, err := NewAmount(9232.80, EUR)
		g.Expect(err).To(gomega.BeNil())

		deposit := RecreateDeposit(
			money,
			RecreateEquivalent(EURPLN, Rate(4.6164), RecreateAmountFromCents(200000, PLN, MoneyPrecision)),
		)
		g.Expect(deposit.Money().AsString()).To(gomega.Equal("9232.80 EUR"))
		g.Expect(deposit.Equivalent().Money().AsString()).To(gomega.Equal("2000.00 PLN"))
		g.Expect(deposit.Equivalent().Ratio().AsFloat64()).To(gomega.Equal(4.6164))
		g.Expect(deposit.Equivalent().Pair().AsString()).To(gomega.Equal(string(EURPLN)))
	})
}
