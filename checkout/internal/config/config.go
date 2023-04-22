package config

import (
	"log"
	"os"
	"route256/libs/config"
)

const configPath = "CONFIG"

type Config struct {
	Token    string `yaml:"token"`
	DSN      string `yaml:"dsn"`
	Services struct {
		CheckoutHTTP   string `yaml:"checkout_http"`
		CheckoutGRPC   string `yaml:"checkout_grpc"`
		Loms           string `yaml:"loms"`
		ProductService string `yaml:"product_service"`
		Jaeger         string `yaml:"jaeger"`
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
