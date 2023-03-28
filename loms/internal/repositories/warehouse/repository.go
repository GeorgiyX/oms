package warehouse

//go:generate mockery --case underscore --name Repository --with-expecter

import (
	"context"
	"route256/libs/db"
	"route256/loms/internal/model"
)

var _ Repository = (*repository)(nil)

type Repository interface {
	SkuStock(ctx context.Context, sku uint32) ([]model.Warehouse, error)
	IsEnough(ctx context.Context, sku uint32, count uint32) (bool, error)
	ReserveNext(ctx context.Context, sku uint32, count uint32, order int64) (uint32, error)
	CancelReserves(ctx context.Context, order []int64) error
}

type repository struct {
	db db.TxDB
}

func New(db db.TxDB) *repository {
	return &repository{
		db: db,
	}
}
