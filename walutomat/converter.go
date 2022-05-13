package walutomat

import (
	"encoding/json"
	"strconv"

	"github.com/pawski/go-xchange/logger"
)

func Convert(response []byte) []Offer {
	var offerResponseDecoded OffersResponse

	err := json.Unmarshal(response, &offerResponseDecoded)

	if err != nil {
		logger.Get().Fatal(err)
	}

	offers := offerResponseToOffer(offerResponseDecoded)

	logger.Get().Printf("Offers in response: %d", len(offers))

	return offers
}

func offerResponseToOffer(offerResponse OffersResponse) []Offer {
	var offers []Offer

	for _, offerResponse := range offerResponse.Offers {
		offers = append(offers, Offer{
			Pair:      offerResponse.Pair,
			Buy:       offerResponse.Buy,
			BuyOld:    offerResponse.BuyOld,
			Sell:      offerResponse.Sell,
			SellOld:   offerResponse.SellOld,
			CountBuy:  offerResponse.CountBuy,
			CountSell: offerResponse.CountSell,
			Avg:       offerResponse.Avg,
			AvgOld:    offerResponse.AvgOld,
		})
	}

	return offers
}

func stringToFloat(floatAsString string) float32 {
	result, err := strconv.ParseFloat(floatAsString, 32)

	if nil != err {
		logger.Get().Printf("Conversion error %s\n", err)
	}

	return float32(result)
}
