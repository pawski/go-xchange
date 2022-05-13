package configuration

import (
	"io/ioutil"
	"sync"

	"gopkg.in/yaml.v2"

	"github.com/pawski/go-xchange/logger"
)

type config struct {
	InfluxDbHost          string `yaml:"influx_host"`
	InfluxDbUser          string `yaml:"influx_user"`
	InfluxDbPassword      string `yaml:"influx_password"`
	InfluxDbDatabase      string `yaml:"influx_database"`
	RabbitMqUrl           string `yaml:"rabbitmq_url"`
	WalutomatUrl          string `yaml:"walutomat_url"`
	CollectUpdateInterval int64  `yaml:"collect_update_interval"`
	WalutomatApiKey       string `yaml:"walutomat_api_key"`
	WalutomatApiHost      string `yaml:"walutomat_api_host"`
}

var cfg *config
var once sync.Once

func Get() *config {
	once.Do(func() {
		loadConfiguration()
	})

	return cfg
}

func loadConfiguration() {
	source, err := ioutil.ReadFile("config.yml")

	if err != nil {
		logger.Get().Fatalf("error: %v", err)
	}

	err = yaml.Unmarshal(source, &cfg)

	if err != nil {
		panic(err)
	}
}
