package walutomat_v1

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/pkg/errors"
)

type BestOfferCollector struct {
	httpClient *http.Client
	offersUrl  string
	logger     *logrus.Logger
}

func NewBestOfferCollector(httpClient *http.Client, offersUrl string, logger *logrus.Logger) *BestOfferCollector {
	return &BestOfferCollector{
		httpClient: httpClient,
		offersUrl:  offersUrl,
		logger:     logger,
	}
}

func (c *BestOfferCollector) FindRecent(_ context.Context) ([]Offer, error) {
	var offers []Offer

	resp, err := c.httpClient.Get(c.offersUrl)

	if err != nil {
		return offers, err
	}

	if resp.StatusCode != 200 {
		return offers, errors.Errorf(
			"Walutomat Api returned non 200 HTTP response code %d, %s",
			resp.StatusCode,
			resp.Status,
		)
	}

	response, err := ioutil.ReadAll(resp.Body)
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			c.logger.Warnf("Walutomat response body close issue: %s", err)
		}
	}()

	var offerResponseDecoded OffersResponse
	err = json.Unmarshal(response, &offerResponseDecoded)

	if err != nil {
		return offers, err
	}

	var walutomatOffers []Offer
	for _, offerResponse := range offerResponseDecoded.Offers {

		buy, err := offerResponse.Buy.Float64()
		if err != nil {
			return offers, errors.WithMessagef(err, "ExchangeOffer.Buy (%s)", offerResponse.Buy.String())
		}

		buyOld, err := offerResponse.BuyOld.Float64()
		if err != nil {
			return offers, errors.WithMessagef(err, "ExchangeOffer.BuyOld (%s)", offerResponse.BuyOld.String())
		}

		sell, err := offerResponse.Sell.Float64()
		if err != nil {
			return offers, errors.WithMessagef(err, "ExchangeOffer.Sell (%s)", offerResponse.Sell.String())
		}

		sellOld, err := offerResponse.SellOld.Float64()
		if err != nil {
			return offers, errors.WithMessagef(err, "ExchangeOffer.SellOld (%s)", offerResponse.SellOld.String())
		}

		avg, err := offerResponse.Avg.Float64()
		if err != nil {
			if offerResponse.Avg.String() == "" {
				avg = 0.0
			} else {
				return offers, errors.WithMessagef(err, "ExchangeOffer.Avg (%s)", offerResponse.Avg.String())
			}
		}

		avgOld, err := offerResponse.AvgOld.Float64()
		if err != nil {
			if offerResponse.AvgOld.String() == "" {
				avgOld = 0.0
			} else {
				return offers, errors.WithMessagef(err, "ExchangeOffer.AvgOld (%s)", offerResponse.AvgOld.String())
			}
		}

		walutomatOffers = append(walutomatOffers, Offer{
			Pair:      offerResponse.Pair,
			Buy:       buy,
			BuyOld:    buyOld,
			Sell:      sell,
			SellOld:   sellOld,
			CountBuy:  offerResponse.CountBuy,
			CountSell: offerResponse.CountSell,
			Avg:       avg,
			AvgOld:    avgOld,
		})
	}

	return walutomatOffers, nil
}
