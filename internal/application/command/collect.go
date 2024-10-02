package command

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/influxdata/influxdb/client/v2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/pawski/go-xchange/internal/configuration"
	"github.com/pawski/go-xchange/internal/influxdb"
	"github.com/pawski/go-xchange/internal/infrastructure/cex/walutomat_v1"
	"github.com/pawski/go-xchange/internal/logger"
)

func CollectExecute(ctx context.Context, dryRun bool) error {

	reg := prometheus.NewRegistry()
	reg.MustRegister(collectors.NewBuildInfoCollector())
	reg.MustRegister(collectors.NewGoCollector())

	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.HandlerFor(
		reg,
		promhttp.HandlerOpts{},
	))

	httpServer := &http.Server{
		Addr:    "localhost:8000",
		Handler: mux,
	}

	go func() {
		err := httpServer.ListenAndServe()
		if err != http.ErrServerClosed {
			fmt.Println(err)
		}
	}()

	logger.Get().Infof("Running...")

	wait := make(chan os.Signal, 1)
	logger.Get().Info("Setting up")

	interval := time.Second * time.Duration(configuration.Get().CollectUpdateInterval)
	ticker := time.NewTicker(interval)

	httpClient := &http.Client{
		Timeout: time.Second * time.Duration(10),
	}
	bestOffersCollector := walutomat_v1.NewBestOfferCollector(httpClient, configuration.Get().WalutomatUrl, logger.Get())

	wg := sync.WaitGroup{}
	go func() {
		wg.Add(1)
		defer wg.Done()
		for t := range ticker.C {
			logger.Get().Println("Start at", t)
			offers, err := bestOffersCollector.FindRecent(ctx)
			if err != nil {
				logger.Get().Error(err)
				continue
			}

			if dryRun {
				fmt.Println(offers, err)
			} else {
				err = sendToStore(offers)
				if err != nil {
					logger.Get().Error(err)
					continue
				}
			}
		}
	}()

	<-wait
	ticker.Stop()
	wg.Wait()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := httpServer.Shutdown(ctx); err != nil {
		logger.Get().Fatalf("Server Shutdown Failed:%+v", err)
	}

	return nil
}

func sendToStore(offers []walutomat_v1.Offer) error {
	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  configuration.Get().InfluxDbDatabase,
		Precision: "s",
	})

	if err != nil {
		return err
	}

	for _, offer := range offers {
		// Create a point and add to batch
		tags := map[string]string{"pair": offer.Pair}
		fields := map[string]interface{}{
			"Buy":       offer.Buy,
			"BuyOld":    offer.BuyOld,
			"CountBuy":  offer.CountBuy,
			"Sell":      offer.Sell,
			"SellOld":   offer.SellOld,
			"CountSell": offer.CountSell,
			"Avg":       offer.Avg,
			"AvgOld":    offer.AvgOld,
		}

		pt, err := client.NewPoint("offers", tags, fields, time.Now())
		if err != nil {
			return err
		}
		bp.AddPoint(pt)
	}

	// Write the batch
	if err := influxdb.Get().Write(bp); err != nil {
		return err
	}

	return nil
}
