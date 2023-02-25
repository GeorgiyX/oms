package main

import (
	"log"
	"net/http"
	"route256/libs/httpaux"
	"route256/loms/internal/config"
	"route256/loms/internal/handlers"
	"route256/loms/internal/usecase"

	"github.com/julienschmidt/httprouter"
)

func main() {
	config.Init()

	useCaseInstance := usecase.New()
	handlerInstance := handlers.New(useCaseInstance)

	router := httprouter.New()
	router.Handler(http.MethodPost, "/createOrder", httpaux.New(handlerInstance.CreateOrder))
	router.Handler(http.MethodPost, "/listOrder", httpaux.New(handlerInstance.ListOrder))
	router.Handler(http.MethodPost, "/orderPayed", httpaux.New(handlerInstance.OrderPayed))
	router.Handler(http.MethodPost, "/cancelOrder", httpaux.New(handlerInstance.CancelOrder))
	router.Handler(http.MethodPost, "/stocks", httpaux.New(handlerInstance.Stock))

	log.Printf("start \"loms\" service at %s\n", config.Instance.Services.Loms)
	log.Fatal(http.ListenAndServe(config.Instance.Services.Loms, router))
}
