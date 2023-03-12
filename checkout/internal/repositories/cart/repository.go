package cart

import (
	"context"
	"route256/libs/db"

	"route256/checkout/internal/model"
)

var _ Repository = (*repository)(nil)

type Repository interface {
	Add(ctx context.Context, user int64, sku uint32, count uint32) error
	Delete(ctx context.Context, user int64, sku uint32, count uint32) error
	List(ctx context.Context, user int64) ([]model.CartItemDB, error)
	RemoveByUser(ctx context.Context, user int64) error
}

type repository struct {
	db db.TxDB
}

func New(db db.TxDB) *repository {
	return &repository{
		db: db,
	}
}
