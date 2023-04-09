package config

import (
	"log"
	"route256/libs/config"
)

type Config struct {
	NotificationTopic string   `yaml:"notification_topic"`
	Brokers           []string `yaml:"brokers"`
}

var Instance Config

func Init() {
	var err error
	Instance, err = config.Read[Config]("notifications/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
}
