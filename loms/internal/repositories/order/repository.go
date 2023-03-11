package order

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"route256/loms/internal/model"
)

var _ Repository = (*repository)(nil)

type Repository interface {
	CreateOrder(ctx context.Context, user int64) (int64, error)
	AddToOrder(ctx context.Context, items []model.OrderItemDB) error
	SetOrderStatus(ctx context.Context, order int64, status model.OrderStatus) error
	GetOrderInfo(ctx context.Context, order int64) (model.OrderInfo, error)
	GetOrderItems(ctx context.Context, order int64) ([]model.OrderItemDB, error)
}

type repository struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *repository {
	return &repository{
		db: db,
	}
}
