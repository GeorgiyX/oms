package usecase

import (
	"context"

	"route256/loms/internal/model"
	"route256/loms/internal/repositories/order"
	"route256/loms/internal/repositories/warehouse"
)

var _ UseCase = (*useCase)(nil)

type UseCase interface {
	CancelOrder(ctx context.Context, orderID int64) error
	OrderPayed(ctx context.Context, orderID int64) error
	Stock(ctx context.Context, sku uint32) ([]model.StocksItemInfo, error)
	CreateOrder(ctx context.Context, user int64, items []model.OrderItemToCreate) (int64, error)
	ListOrder(ctx context.Context, orderID int64) (model.Order, error)
}

type useCase struct {
	warehouseRepo warehouse.Repository
	orderRepo     order.Repository
}

type Config struct {
	WarehouseRepository warehouse.Repository
	OrderRepository     order.Repository
}

func New(config Config) *useCase {
	return &useCase{
		warehouseRepo: config.WarehouseRepository,
		orderRepo:     config.OrderRepository,
	}
}
