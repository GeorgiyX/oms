package model

type StocksRequest struct {
	SKU uint32 `json:"sku"`
}

type StocksItem struct {
	WarehouseID int64  `json:"warehouseID"`
	Count       uint64 `json:"count"`
}

type StocksResponse struct {
	Stocks []StocksItem `json:"stocks"`
}

type Stock struct {
	WarehouseID int64
	Count       uint64
}

type ProductRequest struct {
	Token string `json:"token"`
	SKU   uint32 `json:"sku"`
}

type ProductResponse struct {
	Name  string `json:"name"`
	Price uint32 `json:"price"`
}

type Product struct {
	Name  string
	Price uint32
}

type CreateOrderItem struct {
	SKU   uint32
	Count uint16
}

type CreateOrderRequestItem struct {
	SKU   uint32 `json:"sku"`
	Count uint16 `json:"count"`
}

type CreateOrderRequest struct {
	User  int64                    `json:"user"`
	Items []CreateOrderRequestItem `json:"items"`
}

type CreateOrderResponse struct {
	OrderID int64 `json:"orderID"`
}
