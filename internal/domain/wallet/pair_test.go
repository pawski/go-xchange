package wallet

import (
	"github.com/onsi/gomega"
	"testing"
)

func TestCanConvertPair(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	pair := Pair(EURPLN)

	g.Expect(pair.AsString()).To(gomega.Equal("EURPLN"))
	g.Expect(pair.From()).To(gomega.Equal("EUR"))
	g.Expect(pair.To()).To(gomega.Equal("PLN"))
}
