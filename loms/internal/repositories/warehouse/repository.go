package warehouse

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"route256/loms/internal/model"
)

var _ Repository = (*repository)(nil)

type Repository interface {
	SkuStock(ctx context.Context, sku uint32) ([]model.Warehouse, error)
	Reserve(ctx context.Context, sku uint32, count uint32) error
}

type repository struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *repository {
	return &repository{
		db: db,
	}
}
