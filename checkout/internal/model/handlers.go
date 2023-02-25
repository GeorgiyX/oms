package model

type AddToCartRequest struct {
	User  int64  `json:"user" validate:"required"`
	Sku   uint32 `json:"sku" validate:"required"`
	Count uint16 `json:"count" validate:"required"`
}

type AddToCartResponse struct {
}

type DeleteFromCartRequest struct {
	User  int64  `json:"user" validate:"required"`
	Sku   uint32 `json:"sku" validate:"required"`
	Count uint16 `json:"count" validate:"required"`
}

type DeleteFromCartResponse struct {
}

type PurchaseRequest struct {
	User int64 `json:"user" validate:"required"`
}

type PurchaseResponse struct {
	OrderID int64 `json:"orderID"`
}

type ListCartRequest struct {
	User int64 `json:"user" validate:"required"`
}

type CartItemResponse struct {
	Sku   uint32 `json:"sku"`
	Count uint16 `json:"count"`
	Name  string `json:"name"`
	Price uint32 `json:"price"`
}

type CartResponse struct {
	Items      []*CartItemResponse `json:"items"`
	TotalPrice uint32              `json:"totalPrice"`
}
