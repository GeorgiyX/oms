package model

import "time"

type OrderInfo struct {
	ID        int64       `db:"id"`
	UserID    int64       `db:"user_id"`
	CreatedAt time.Time   `db:"created_at"`
	Status    OrderStatus `db:"status"`
}

type OrderItemDB struct {
	Sku         uint32 `db:"sku"`
	Count       uint32 `db:"count"`
	OrderInfoID int64  `db:"fk_order_info_id"`
}

type Warehouse struct {
	WarehouseID      int64  `db:"warehouse_id"`
	Sku              uint32 `db:"sku"`
	AvailableToOrder uint64 `db:"available_to_order"`
}

type Reserve struct {
	Sku         uint32 `db:"sku"`
	WarehouseID int64  `db:"warehouse_id"`
	Count       uint32 `db:"count"`
	OrderInfoID int64  `db:"fk_order_info_id"`
}

type StatusChangeDatabase struct {
	OrderID            int64              `db:"fk_order_info_id"`
	CreatedAt          time.Time          `db:"created_at"`
	SendStatus         int16              `db:"send_status"`
	NotificationStatus NotificationStatus `db:"notification_status"`
	Total              int64              `db:"total"`
}
