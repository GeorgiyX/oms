package order

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/pkg/errors"
	"route256/loms/internal/model"
)

func (r *repository) Add(ctx context.Context, user int64, sku uint32, count uint32) error {
	const query = `INSERT INTO warehouse(user_id, sku, count) VALUES ($1, $2, $3)
	ON CONFLICT (user_id, sku) DO UPDATE
    SET count = warehouse.count + excluded.count;`

	_, err := r.db.Exec(ctx, query, user, sku, count)
	if err != nil {
		return errors.Wrap(err, "cannot add sku to warehouse")
	}

	return nil
}

func (r *repository) Delete(ctx context.Context, user int64, sku uint32, count uint32) error {
	const query = `WITH upd AS (UPDATE warehouse SET count = warehouse.count - $3 WHERE user_id = $1 AND sku = $2)
	DELETE FROM warehouse WHERE user_id = $1 AND sku = $2 AND count <= 0;`

	_, err := r.db.Exec(ctx, query, user, sku, count)
	if err != nil {
		return errors.Wrap(err, "cannot delete sku from warehouse")
	}

	return nil
}

func (r *repository) List(ctx context.Context, user int64) ([]model.CartItemDB, error) {
	const query = `SELECT user_id, sku, count FROM warehouse WHERE user_id = $1;`

	var items []model.CartItemDB
	err := pgxscan.Select(ctx, r.db, &items, query, user)
	if err != nil {
		return nil, errors.Wrap(err, "cannot fetch user warehouse items")
	}

	return items, nil
}

func (r *repository) RemoveByUser(ctx context.Context, user int64) error {
	const query = `DELETE FROM warehouse WHERE user_id = $1`

	resp, err := r.db.Exec(ctx, query, user)
	if err != nil {
		return errors.Wrap(err, "cannot remove user warehouse items")
	}

	affected := resp.RowsAffected()
	if affected != 1 {
		return errors.New("nothing was deleted")
	}

	return nil
}
