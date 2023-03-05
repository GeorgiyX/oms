package model

type StocksItemInfo struct {
	WarehouseID int64
	Count       uint64
}

type OrderItem struct {
	SKU   uint32
	Count uint16
}

type Order struct {
	Status  string
	User    int64
	Items   []OrderItem
	OrderID int64
}

type OrderItemToCreate struct {
	SKU   uint32 `json:"sku"`
	Count uint16 `json:"count"`
}
