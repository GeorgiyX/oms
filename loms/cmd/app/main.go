package main

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"net"
	"route256/libs/cron"
	"route256/libs/db"
	"route256/libs/logger"
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
		logger.Fatal("failed to сreate notifier", zap.Error(err))
	}
	defer notifierInstance.Close()
	useCaseInstance.SetNotifier(notifierInstance)

	serviceInstance := loms.New(useCaseInstance)

	s := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpcMiddleware.ChainUnaryServer(
				middleware.LoggingInterceptor(lg),
			),
		),
	)

	c := cron.New()
	defer s.Stop()
	c.Add(useCaseInstance.GetCancelOrdersByTimeoutCron())
	c.Add(useCaseInstance.GetNotifyCron())

	reflection.Register(s)
	desc.RegisterLomsServer(s, serviceInstance)

	logger.Info(fmt.Sprintf("start \"loms\" checkout at %s\n", config.Instance.Services.Loms), zap.Error(err))
	lis, err := net.Listen("tcp", config.Instance.Services.Loms)
	if err != nil {
		logger.Fatal("create tcp listener", zap.Error(err))
	}
	if err = s.Serve(lis); err != nil {
		logger.Fatal("failed to serve", zap.Error(err))
	}
}
