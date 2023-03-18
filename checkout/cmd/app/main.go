package main

import (
	"context"
	"log"
	"net"
	"route256/checkout/internal/app/checkout"
	"route256/checkout/internal/clients/loms"
	productService "route256/checkout/internal/clients/product_service"
	"route256/checkout/internal/config"
	"route256/checkout/internal/repositories/cart"
	checkout2 "route256/checkout/internal/usecases/checkout"
	desc "route256/checkout/pkg/checkout"
	descProductService "route256/checkout/pkg/product-service"
	"route256/libs/db"
	"route256/libs/middleware"
	descLoms "route256/loms/pkg/loms"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func main() {
	config.Init()

	ctx := context.Background()
	pool, err := pgxpool.Connect(ctx, config.Instance.DSN)
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}
	defer pool.Close()

	err = pool.Ping(ctx)
	if err != nil {
		log.Fatalf("unsuccess db ping: %v", err)
	}

	txDB := db.NewPgxPoolDB(pool)

	opts := grpc.WithTransportCredentials(insecure.NewCredentials())

	connLoms, err := grpc.Dial(config.Instance.Services.Loms, opts)
	if err != nil {
		log.Fatalf("failed to connect to loms: %v", err)
	}
	defer connLoms.Close()
	lomsClient := descLoms.NewLomsClient(connLoms)

	connProduct, err := grpc.Dial(config.Instance.Services.ProductService, opts)
	if err != nil {
		log.Fatalf("failed to connect to product checkout: %v", err)
	}
	defer connLoms.Close()
	productServiceClient := descProductService.NewProductServiceClient(connProduct)

	useCaseConfig := checkout2.Config{
		StocksChecker: loms.New(lomsClient),
		SkuResolver:   productService.New(productServiceClient),
		Repository:    cart.New(txDB),
	}
	useCaseInstance := checkout2.New(useCaseConfig)
	serviceInstance := checkout.New(useCaseInstance)

	s := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpcMiddleware.ChainUnaryServer(
				middleware.LoggingInterceptor,
			),
		),
	)

	reflection.Register(s)
	desc.RegisterCheckoutServer(s, serviceInstance)

	log.Printf("start \"checkout\" checkout at %s\n", config.Instance.Services.Checkout)
	lis, err := net.Listen("tcp", config.Instance.Services.Checkout)
	if err != nil {
		log.Fatalf("create tcp listener: %v", err)
	}
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
