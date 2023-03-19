package order

import (
	"context"
	"route256/libs/db"
	"route256/loms/internal/model"
)

var _ Repository = (*repository)(nil)

type Repository interface {
	CreateOrder(ctx context.Context, user int64) (int64, error)
	AddToOrder(ctx context.Context, items []model.OrderItemDB, order int64) error
	SetOrderStatuses(ctx context.Context, order []int64, status model.OrderStatus) error
	GetOrderInfo(ctx context.Context, order int64) (model.OrderInfo, error)
	GetOrderItems(ctx context.Context, order int64) ([]model.OrderItemDB, error)
	GetExpiredPaymentOrders(ctx context.Context) ([]int64, error)
}

type repository struct {
	db db.TxDB
}

func New(db db.TxDB) *repository {
	return &repository{
		db: db,
	}
}
