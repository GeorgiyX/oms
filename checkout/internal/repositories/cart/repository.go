package cart

//go:generate mockery --case underscore --name Repository --with-expecter

import (
	"context"
	"route256/checkout/internal/model"
	"route256/libs/db"
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
