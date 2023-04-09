package main

import (
	"context"
	"log"
	"net"
	"route256/libs/cron"
	"route256/libs/db"
	"route256/libs/middleware"
	"route256/loms/internal/app/loms"
	"route256/loms/internal/config"
	"route256/loms/internal/notifier"
	"route256/loms/internal/repositories/order"
	"route256/loms/internal/repositories/warehouse"
	loms2 "route256/loms/internal/usecase/loms"
	desc "route256/loms/pkg/loms"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
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

	useCaseInstance := loms2.New(
		loms2.Config{
			WarehouseRepository: warehouse.New(txDB),
			OrderRepository:     order.New(txDB),
			Notifier:            nil,
			TxDB:                txDB,
		},
	)

	notifierInstance, err := notifier.NewNotifier(useCaseInstance)
	if err != nil {
		log.Fatalf("create notifier: %v", err)
	}
	defer notifierInstance.Close()
	useCaseInstance.SetNotifier(notifierInstance)

	serviceInstance := loms.New(useCaseInstance)

	s := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpcMiddleware.ChainUnaryServer(
				middleware.LoggingInterceptor,
			),
		),
	)

	c := cron.New()
	defer s.Stop()
	c.Add(useCaseInstance.GetCancelOrdersByTimeoutCron())
	c.Add(useCaseInstance.GetNotifyCron())

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
