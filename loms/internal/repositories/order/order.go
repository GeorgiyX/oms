package order

import (
	"context"

	"github.com/pkg/errors"
	"route256/loms/internal/model"
)

func (r *repository) CreateOrder(ctx context.Context, user int64) (int64, error) {
	const query = `INSERT INTO order_info(user_id, status) VALUES ($1, status_new()) RETURNING id;`

	var id int64
	err := r.db.Get(ctx, &id, query, user)
	if err != nil {
		return 0, errors.Wrap(err, "cannot create order info")
	}

	return id, nil
}

func (r *repository) AddToOrder(ctx context.Context, items []model.OrderItemDB, order int64) error {
	const query = `INSERT INTO order_item(sku, fk_order_info_id, count) SELECT UNNEST($1::INTEGER[]), $2, UNNEST($3::INTEGER[]);`

	skus := make([]uint32, 0, len(items))
	count := make([]uint32, 0, len(items))
	for _, item := range items {
		skus = append(skus, item.Sku)
		count = append(count, item.Count)
	}

	_, err := r.db.Exec(ctx, query, skus, order, count)
	if err != nil {
		return errors.Wrap(err, "cannot add sku to order")
	}

	return nil
}

func (r *repository) SetOrderStatus(ctx context.Context, order int64, status model.OrderStatus) error {
	const query = `UPDATE order_info SET status = $2 WHERE id = $1;`

	resp, err := r.db.Exec(ctx, query, order, status)
	if err != nil {
		return errors.Wrap(err, "cannot change order status")
	}

	affected := resp.RowsAffected()
	if affected != 1 {
		return errors.New("nothing was updated")
	}

	return nil
}

func (r *repository) GetOrderInfo(ctx context.Context, order int64) (model.OrderInfo, error) {
	const query = `SELECT id, user_id, created_at, status FROM order_info WHERE id = $1;`

	var info model.OrderInfo
	err := r.db.Get(ctx, &info, query, order)
	if err != nil {
		return model.OrderInfo{}, errors.Wrap(err, "cannot fetch order info")
	}

	return info, nil
}

func (r *repository) GetOrderItems(ctx context.Context, order int64) ([]model.OrderItemDB, error) {
	const query = `SELECT sku, fk_order_info_id, count FROM order_item WHERE fk_order_info_id = $1;`

	var items []model.OrderItemDB
	err := r.db.Select(ctx, &items, query, order)
	if err != nil {
		return nil, errors.Wrap(err, "cannot fetch order items")
	}

	return items, nil
}
