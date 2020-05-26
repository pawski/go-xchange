package walutomat

import (
	"encoding/json"
	"github.com/pawski/go-xchange/logger"
	"strconv"
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
			Buy:       stringToFloat(offerResponse.Buy.String()),
			BuyOld:    stringToFloat(offerResponse.BuyOld.String()),
			Sell:      stringToFloat(offerResponse.Sell.String()),
			SellOld:   stringToFloat(offerResponse.SellOld.String()),
			CountBuy:  offerResponse.CountBuy,
			CountSell: offerResponse.CountSell,
			Avg:       stringToFloat(offerResponse.Avg.String()),
			AvgOld:    stringToFloat(offerResponse.AvgOld.String()),
		})
	}

	return offers
}

func stringToFloat(floatAsString string) float32 {
	if floatAsString == "" {
		return 0.0
	}

	result, err := strconv.ParseFloat(floatAsString, 32)

	if nil != err {
		logger.Get().Errorf("Conversion error %s\n", err)
		return 0.0
	}

	return float32(result)
}
