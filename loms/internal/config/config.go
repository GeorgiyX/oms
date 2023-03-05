package config

import (
	"log"
	"route256/libs/config"
)

type Config struct {
	Services struct {
		Loms string `yaml:"loms"`
	} `yaml:"services"`
}

var Instance Config

func Init() {
	var err error
	Instance, err = config.Read[Config]("loms/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
}
