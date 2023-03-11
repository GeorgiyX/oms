package config

import (
	"log"
	"route256/libs/config"
)

type Config struct {
	Token    string `yaml:"token"`
	DSN      string `yaml:"dsn"`
	Services struct {
		Checkout       string `yaml:"checkout"`
		Loms           string `yaml:"loms"`
		ProductService string `yaml:"product_service"`
	} `yaml:"services"`
}

var Instance Config

func Init() {
	var err error
	Instance, err = config.Read[Config]("checkout/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
}
