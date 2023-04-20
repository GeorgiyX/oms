package config

import (
	"log"
	"route256/libs/config"
)

type Config struct {
	Token    string `yaml:"token"`
	DSN      string `yaml:"dsn"`
	Services struct {
		CheckoutHTTP   string `yaml:"checkout_http"`
		CheckoutGRPC   string `yaml:"checkout_grpc"`
		Loms           string `yaml:"loms"`
		ProductService string `yaml:"product_service"`
	} `yaml:"services"`
	Debug bool `json:"debug"`
}

var Instance Config

func Init() {
	var err error
	Instance, err = config.Read[Config]("checkout/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
}
