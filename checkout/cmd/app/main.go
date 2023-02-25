package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"route256/checkout/internal/clients/http/product_service"
	"route256/checkout/internal/handlers"

	"route256/checkout/internal/clients/http/loms"
	"route256/checkout/internal/config"
	"route256/checkout/internal/usecase"
	"route256/libs/httpaux"
)

func main() {
	config.Init()

	useCaseConfig := usecase.Config{
		StocksChecker: loms.New(config.Instance.Services.Loms),
		SkuResolver:   product_service.New(config.Instance.Services.ProductService),
	}
	useCaseInstance := usecase.New(useCaseConfig)
	handlerInstance := handlers.New(useCaseInstance)

	router := httprouter.New()
	router.Handler(http.MethodPost, "/handlerInstance", httpaux.New(handlerInstance.AddToCart))
	router.Handler(http.MethodPost, "/deleteFromCart", httpaux.New(handlerInstance.DeleteFromCart))
	router.Handler(http.MethodPost, "/listCart", httpaux.New(handlerInstance.ListCart))
	router.Handler(http.MethodPost, "/purchase", httpaux.New(handlerInstance.Purchase))

	log.Printf("start checkout service at %s\n", config.Instance.Services.Checkout)
	log.Fatal(http.ListenAndServe(config.Instance.Services.Checkout, router))
}
