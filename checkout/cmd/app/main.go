package main

import (
	"context"
	"fmt"
	"net"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"route256/checkout/internal/app/checkout"
	"route256/checkout/internal/clients/loms"
	productService "route256/checkout/internal/clients/product_service"
	"route256/checkout/internal/config"
	"route256/checkout/internal/repositories/cart"
	checkout2 "route256/checkout/internal/usecases/checkout"
	desc "route256/checkout/pkg/checkout"
	descProductService "route256/checkout/pkg/product-service"
	"route256/libs/db"
	"route256/libs/logger"
	"route256/libs/middleware"
	descLoms "route256/loms/pkg/loms"
)

func main() {
	config.Init()
	logger.Init(config.Instance.Debug)

	lg, err := logger.New(config.Instance.Debug)
	if err != nil {
		logger.Fatal("create logger", zap.Error(err))
	}

	ctx := context.Background()
	pool, err := pgxpool.Connect(ctx, config.Instance.DSN)
	if err != nil {
		logger.Fatal("failed to connect to DB", zap.Error(err))
	}
	defer pool.Close()

	err = pool.Ping(ctx)
	if err != nil {
		logger.Fatal("unsuccessful db ping", zap.Error(err))
	}

	txDB := db.NewPgxPoolDB(pool)

	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	interceptor := grpc.WithChainUnaryInterceptor(middleware.RateLimiterInterceptor())

	connLoms, err := grpc.Dial(config.Instance.Services.Loms, opts, interceptor)
	if err != nil {
		logger.Fatal("failed to connect to loms", zap.Error(err))
	}
	defer connLoms.Close()
	lomsClient := descLoms.NewLomsClient(connLoms)

	connProduct, err := grpc.Dial(config.Instance.Services.ProductService, opts)
	if err != nil {
		logger.Fatal("failed to connect to product service", zap.Error(err))
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
				middleware.LoggingInterceptor(lg),
			),
		),
	)

	reflection.Register(s)
	desc.RegisterCheckoutServer(s, serviceInstance)

	logger.Info(fmt.Sprintf("start \"checkout\" checkout at %s\n", config.Instance.Services.Checkout), zap.Error(err))
	lis, err := net.Listen("tcp", config.Instance.Services.Checkout)
	if err != nil {
		logger.Fatal("create tcp listener", zap.Error(err))
	}
	if err = s.Serve(lis); err != nil {
		logger.Fatal("failed to serve", zap.Error(err))
	}
}
