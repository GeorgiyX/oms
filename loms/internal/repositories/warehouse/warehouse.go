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

func (r *repository) IsEnough(ctx context.Context, sku uint32, count uint32) (bool, error) {
	const query = `SELECT sum(available_to_order) >= $2 FROM warehouse WHERE sku = $1;`

	var isEnough bool
	err := pgxscan.Get(ctx, r.db, &isEnough, query, sku, count)
	if err != nil {
		return false, errors.Wrap(err, "check is enough")
	}

	return isEnough, nil
}

// ReserveNext trying to reserve sku at warehouse with greatest sku count, returns remain sku count to reserve if on found warehouse not enough sku
func (r *repository) ReserveNext(ctx context.Context, sku uint32, count uint32, order int64) (uint32, error) {
	const query = `
	WITH to_decrement AS (
    SELECT
        warehouse_id,
        sku,
        CASE WHEN $2 > available_to_order THEN $2 - available_to_order ELSE 0 END AS remain,
        CASE WHEN $2 > available_to_order THEN 0 ELSE available_to_order - $2 END AS new_available_to_order,
        CASE WHEN $2 > available_to_order THEN available_to_order ELSE $2 END AS reserved
        FROM warehouse WHERE sku = $1 AND available_to_order > 0 ORDER BY available_to_order DESC LIMIT 1),
    to_reserve AS (
        UPDATE warehouse w SET available_to_order = td.new_available_to_order
        FROM to_decrement td WHERE w.warehouse_id = td.warehouse_id AND w.sku = td.sku
        RETURNING td.sku, td.warehouse_id, td.remain, td.reserved)
	INSERT INTO reserve(sku, warehouse_id, count, fk_order_info_id)
		   SELECT sku, warehouse_id, reserved, $3 FROM to_reserve
	RETURNING remain;`

	var remain uint32
	err := pgxscan.Get(ctx, r.db, &remain, query, sku, count)
	if err != nil {
		return 0, errors.Wrap(err, "cannot reserve sku")
	}

	return remain, nil
}

func (r *repository) CancelReserve(ctx context.Context, order int64) error {
	const query = `
	WITH cancel_reserve AS (
		UPDATE warehouse w SET available_to_order = w.available_to_order + r.count
		FROM reserve r WHERE w.warehouse_id = r.warehouse_id AND w.sku = r.sku AND r.fk_order_info_id = $1
	)
	DELETE FROM reserve WHERE fk_order_info_id = $1;`

	resp, err := r.db.Exec(ctx, query, order)
	if err != nil {
		return errors.Wrap(err, "cannot cancel reserve")
	}

	affected := resp.RowsAffected()
	if affected != 1 {
		return errors.New("nothing was changed while cancel reserve")
	}

	return nil
}
