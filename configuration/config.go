package configuration

import (
	"io/ioutil"
	"sync"

	"gopkg.in/yaml.v2"

	"github.com/pawski/go-xchange/logger"
)

type Config struct {
	InfluxDbHost          string `yaml:"influx_host"`
	InfluxDbDatabase      string `yaml:"influx_database"`
	RabbitMqUrl           string `yaml:"rabbitmq_url"`
	WalutomatUrl          string `yaml:"walutomat_url"`
	CollectUpdateInterval int64  `yaml:"collect_update_interval"`
}

var cfg Config
var once sync.Once

func Get() Config {
	once.Do(func() {
		cfg = loadConfiguration()
	})

	return cfg
}

func loadConfiguration() Config {
	source, err := ioutil.ReadFile("config.yml")

	if err != nil {
		logger.Get().Fatalf("error: %v", err)
	}

	var config Config

	err = yaml.Unmarshal([]byte(source), &config)

	if err != nil {
		panic(err)
	}

	return config
}
