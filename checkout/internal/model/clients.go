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
	SKU uint32 `json:"sku"`
}

type ProductResponse struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Product struct {
	Name  string
	Price float64
}
