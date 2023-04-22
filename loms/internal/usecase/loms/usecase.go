package loms

import (
	"context"

	"route256/libs/cron"
	"route256/libs/db"
	"route256/loms/internal/model"
	"route256/loms/internal/notifier"
	notificationOutbox "route256/loms/internal/repositories/notification_outbox"
	"route256/loms/internal/repositories/order"
	"route256/loms/internal/repositories/warehouse"
)

var _ UseCase = (*useCase)(nil)
var _ notifier.NotificationCallBack = (*useCase)(nil)

type UseCase interface {
	CancelOrder(ctx context.Context, orderID int64) error
	OrderPayed(ctx context.Context, orderID int64) error
	Stock(ctx context.Context, sku uint32) ([]model.StocksItemInfo, error)
	CreateOrder(ctx context.Context, user int64, items []model.OrderItemToCreate) (int64, error)
	ListOrder(ctx context.Context, orderID int64) (model.Order, error)
	CancelOrdersByTimeout(ctx context.Context) error
	GetCancelOrdersByTimeoutCron() cron.TaskDescriptor
	SetNotifier(notifier notifier.Notifier)
}

type useCase struct {
	warehouseRepo      warehouse.Repository
	orderRepo          order.Repository
	notifierOutboxRepo notificationOutbox.Repository
	notifier           notifier.Notifier
	db                 db.TxDB
}

type Config struct {
	WarehouseRepository warehouse.Repository
	OrderRepository     order.Repository
	NotifierOutboxRepo  notificationOutbox.Repository
	Notifier            notifier.Notifier
	TxDB                db.TxDB
}

func New(config Config) *useCase {
	return &useCase{
		warehouseRepo:      config.WarehouseRepository,
		orderRepo:          config.OrderRepository,
		notifierOutboxRepo: config.NotifierOutboxRepo,
		notifier:           config.Notifier,
		db:                 config.TxDB,
	}
}
