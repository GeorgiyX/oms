package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"

	"route256/checkout/internal/clients/http/loms"
	"route256/checkout/internal/config"
	"route256/checkout/internal/domain"
	"route256/checkout/internal/handlers/addtocart"
	"route256/libs/srvwrapper"
)

func main() {
	config.Init()
	lomsClient := loms.New(config.Instance.Services.Loms)
	businessLogic := domain.New(lomsClient)
	addToCart := addtocart.New(businessLogic)

	router := httprouter.New()
	router.Handler(http.MethodPost, "/createOrder", srvwrapper.New(addToCart.Handle))

	log.Printf("start checkout service at %s\n", config.Instance.Services.Checkout)
	log.Fatal(http.ListenAndServe(config.Instance.Services.Checkout, router))
}
