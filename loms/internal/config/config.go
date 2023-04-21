package config

import (
	"log"
	"route256/libs/config"
)

type Config struct {
	DSN               string   `yaml:"dsn"`
	NotificationTopic string   `yaml:"notification_topic"`
	Brokers           []string `yaml:"brokers"`
	Services          struct {
		LomsHTTP string `yaml:"loms_http"`
		LomsGRPC string `yaml:"loms_grpc"`
	} `yaml:"services"`
	Debug bool `json:"debug"`
}

var Instance Config

func Init() {
	var err error
	Instance, err = config.Read[Config]("loms/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
}
