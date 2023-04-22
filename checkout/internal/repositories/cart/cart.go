package cart

import (
	"context"
	"route256/checkout/internal/model"

	"github.com/pkg/errors"
)

func (r *repository) Add(ctx context.Context, user int64, sku uint32, count uint32) error {
	const query = `INSERT INTO cart(user_id, sku, count) VALUES ($1, $2, $3)
	ON CONFLICT (user_id, sku) DO UPDATE
    SET count = cart.count + excluded.count;`

	_, err := r.db.Exec(ctx, query, user, sku, count)
	if err != nil {
		return errors.Wrap(err, "cannot add sku to cart")
	}

	return nil
}

func (r *repository) Delete(ctx context.Context, user int64, sku uint32, count uint32) error {
	const query = `WITH upd AS (UPDATE cart SET count = cart.count - $3 WHERE user_id = $1 AND sku = $2)
	DELETE FROM cart WHERE user_id = $1 AND sku = $2 AND count <= 0;`

	_, err := r.db.Exec(ctx, query, user, sku, count)
	if err != nil {
		return errors.Wrap(err, "cannot delete sku from cart")
	}

	return nil
}

func (r *repository) List(ctx context.Context, user int64) ([]model.CartItemDB, error) {
	const query = `SELECT user_id, sku, count FROM cart WHERE user_id = $1;`

	var items []model.CartItemDB
	err := r.db.Select(ctx, &items, query, user)
	if err != nil {
		return nil, errors.Wrap(err, "cannot fetch user cart items")
	}

	return items, nil
}

func (r *repository) RemoveByUser(ctx context.Context, user int64) error {
	const query = `DELETE FROM cart WHERE user_id = $1`

	_, err := r.db.Exec(ctx, query, user)
	if err != nil {
		return errors.Wrap(err, "cannot remove user cart items")
	}

	return nil
}
