package wallet

import (
	"testing"

	"github.com/onsi/gomega"
)

func TestCanRecreateExchangeOfferFromValue(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	offer := RecreateExchangeOfferFromValue(Pair(EURPLN), RecreateCurrencyRate(EURPLN, Rate(4.5500)), RecreateCurrencyRate(EURPLN, Rate(4.4500)))

	g.Expect(offer).To(gomega.Equal(ExchangeOffer{
		pair:     EURPLN,
		sellRate: CurrencyRate{pair: EURPLN, rate: 4.5500},
		buyRate:  CurrencyRate{pair: EURPLN, rate: 4.4500},
	}))
}
