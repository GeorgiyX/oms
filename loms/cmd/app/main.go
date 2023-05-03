package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"route256/libs/cron"
	"route256/libs/db"
	grpcServer "route256/libs/grpc/server"
	"route256/libs/logger"
	"route256/libs/middleware"
	"route256/loms/internal/app/loms"
	"route256/loms/internal/config"
	"route256/loms/internal/notifier"
	notificationOutbox "route256/loms/internal/repositories/notification_outbox"
	"route256/loms/internal/repositories/order"
	"route256/loms/internal/repositories/warehouse"
	loms2 "route256/loms/internal/usecase/loms"
	desc "route256/loms/pkg/loms"
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
			NotifierOutboxRepo:  notificationOutbox.New(txDB),
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
	c := cron.New()
	defer c.Stop()
	c.Add(useCaseInstance.GetCancelOrdersByTimeoutCron())
	c.Add(useCaseInstance.GetNotifyCron())

	server, err := grpcServer.NewServer("loms", config.Instance.Services.Jaeger, grpc.ChainUnaryInterceptor(
		middleware.LoggingInterceptor(lg)),
	)
	if err != nil {
		logger.Fatal("create server", zap.Error(err))
	}

	server.RegisterService(&desc.Loms_ServiceDesc, serviceInstance)

	logger.Info(fmt.Sprintf("start \"loms\" at %s\n", config.Instance.Services.LomsGRPC))
	err = server.Serve(config.Instance.Services.LomsGRPC, config.Instance.Services.LomsHTTP)
	if err != nil {
		logger.Fatal("failed to serve", zap.Error(err))
	}
}
