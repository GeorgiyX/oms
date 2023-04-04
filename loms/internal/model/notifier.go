package model

type StatusChange struct {
	OrderID int64
	Status  OrderStatus
}

type StatusChangeKafka struct {
	OrderID int64 `json:"order_id"`
	Status  int16 `json:"status"`
}
