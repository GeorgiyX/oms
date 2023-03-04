package main

import (
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"route256/checkout/internal/clients/loms"
	productService "route256/checkout/internal/clients/product_service"
	"route256/checkout/internal/config"
	"route256/checkout/internal/service"
	"route256/checkout/internal/usecase"
	desc "route256/checkout/pkg/checkout"
	descProductService "route256/checkout/pkg/product-service"
	"route256/libs/middleware"
	descLoms "route256/loms/pkg/loms"
)

func main() {
	config.Init()

	opts := grpc.WithTransportCredentials(insecure.NewCredentials())

	connLoms, err := grpc.Dial(config.Instance.Services.Loms, opts)
	if err != nil {
		log.Fatalf("failed to connect to loms: %v", err)
	}
	defer connLoms.Close()
	lomsClient := descLoms.NewLomsClient(connLoms)

	connProduct, err := grpc.Dial(config.Instance.Services.ProductService, opts)
	if err != nil {
		log.Fatalf("failed to connect to product service: %v", err)
	}
	defer connLoms.Close()
	productServiceClient := descProductService.NewProductServiceClient(connProduct)

	useCaseConfig := usecase.Config{
		StocksChecker: loms.New(lomsClient),
		SkuResolver:   productService.New(productServiceClient),
	}
	useCaseInstance := usecase.New(useCaseConfig)
	serviceInstance := service.New(useCaseInstance)

	s := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				middleware.LoggingInterceptor,
			),
		),
	)

	reflection.Register(s)
	desc.RegisterCheckoutServer(s, serviceInstance)

	log.Printf("start \"checkout\" service at %s\n", config.Instance.Services.Checkout)
	lis, err := net.Listen("tcp", config.Instance.Services.Checkout)
	if err != nil {
		log.Fatalf("create tcp listener: %v", err)
	}
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
