package model

type CreateOrderRequestItem struct {
	SKU   uint32 `json:"sku"`
	Count uint16 `json:"count"`
}

type CreateOrderRequest struct {
	User  int64                    `json:"user" validate:"required"`
	Items []CreateOrderRequestItem `json:"items" validate:"required"`
}

type CreateOrderResponse struct {
	OrderID int64 `json:"orderID"`
}

type CancelOrderRequest struct {
	OrderID int64 `json:"orderID" validate:"required"`
}

type CancelOrderResponse struct {
}

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

type OrderPayedRequest struct {
	OrderID int64 `json:"orderID" validate:"required"`
}

type OrderPayedResponse struct {
}

type ListOrderRequest struct {
	OrderID int64 `json:"orderID" validate:"required"`
}

type ListOrderItemResponse struct {
	SKU   uint32 `json:"sku"`
	Count uint16 `json:"count"`
}

type ListOrderResponse struct {
	Status  string                  `json:"status"`
	User    int64                   `json:"user"`
	Items   []ListOrderItemResponse `json:"items"`
	OrderID int64                   `json:"orderID"`
}
