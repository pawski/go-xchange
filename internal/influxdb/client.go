package influxdb

import (
	"sync"

	"github.com/influxdata/influxdb/client/v2"

	"github.com/pawski/go-xchange/internal/configuration"
	"github.com/pawski/go-xchange/internal/logger"
)

var dbClient client.Client
var once sync.Once

func Get() client.Client {
	once.Do(func() {
		// Create a new HTTPClient
		var err error
		dbClient, err = client.NewHTTPClient(client.HTTPConfig{
			Addr:     configuration.Get().InfluxDbHost,
			Username: configuration.Get().InfluxDbUser,
			Password: configuration.Get().InfluxDbPassword,
		})

		if err != nil {
			logger.Get().Fatal(err)
		}
	})

	return dbClient
}
