package influxdb

import (
	"github.com/pawski/go-xchange/configuration"
	"github.com/pawski/go-xchange/logger"
	"github.com/influxdata/influxdb/client/v2"
	"sync"
)

var dbClient client.Client
var once sync.Once

func Get() client.Client  {
	once.Do(func(){
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
