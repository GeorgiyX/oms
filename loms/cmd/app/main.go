package main

import (
	"context"
	"log"
	"net"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"route256/libs/db"
	"route256/libs/middleware"
	"route256/loms/internal/app/loms"
	"route256/loms/internal/config"
	"route256/loms/internal/repositories/order"
	"route256/loms/internal/repositories/warehouse"
	"route256/loms/internal/usecase"
	desc "route256/loms/pkg/loms"
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

	useCaseInstance := usecase.New(
		usecase.Config{
			WarehouseRepository: warehouse.New(txDB),
			OrderRepository:     order.New(txDB),
			TxDB:                txDB,
		},
	)
	serviceInstance := loms.New(useCaseInstance)

	s := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpcMiddleware.ChainUnaryServer(
				middleware.LoggingInterceptor,
			),
		),
	)

	reflection.Register(s)
	desc.RegisterLomsServer(s, serviceInstance)

	log.Printf("start \"loms\" checkout at %s\n", config.Instance.Services.Loms)
	lis, err := net.Listen("tcp", config.Instance.Services.Loms)
	if err != nil {
		log.Fatalf("create tcp listener: %v", err)
	}
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
