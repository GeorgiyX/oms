package config

import (
	"log"
	"os"

	"route256/libs/config"
)

const configPath = "CONFIG"

type Config struct {
	DSN               string   `yaml:"dsn"`
	NotificationTopic string   `yaml:"notification_topic"`
	Brokers           []string `yaml:"brokers"`
	Services          struct {
		LomsHTTP string `yaml:"loms_http"`
		LomsGRPC string `yaml:"loms_grpc"`
		Jaeger   string `yaml:"jaeger"`
	} `yaml:"services"`
	Debug bool `json:"debug"`
}

var Instance Config

func Init() {
	var err error
	Instance, err = config.Read[Config](os.Getenv(configPath))
	if err != nil {
		log.Fatal(err)
	}
}
