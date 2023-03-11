package warehouse

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/pkg/errors"
	"route256/loms/internal/model"
)

func (r *repository) SkuStock(ctx context.Context, sku uint32) ([]model.Warehouse, error) {
	const query = `SELECT warehouse_id, sku, available_to_order FROM warehouse WHERE sku = $1;`

	var items []model.Warehouse
	err := pgxscan.Select(ctx, r.db, &items, query, sku)
	if err != nil {
		return nil, errors.Wrap(err, "cannot fetch sku stock")
	}

	return items, nil
}
