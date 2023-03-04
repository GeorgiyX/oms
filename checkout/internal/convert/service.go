package convert

import (
	"route256/checkout/internal/model"
	desc "route256/checkout/pkg/checkout"
)

func ToCartResponse(cart model.Cart) *desc.ListCartResponse {
	out := &desc.ListCartResponse{}
	out.TotalPrice = cart.TotalPrice
	out.Items = make([]*desc.ListCartResponse_CartItemResponse, 0, len(cart.Items))
	for _, item := range cart.Items {
		out.Items = append(out.Items, &desc.ListCartResponse_CartItemResponse{
			Sku:   item.Sku,
			Count: item.Count,
			Name:  item.Name,
			Price: item.Price,
		})
	}
	return out
}

func ToPurchaseResponse(orderID int64) *desc.PurchaseResponse {
	return &desc.PurchaseResponse{
		OrderId: orderID,
	}
}
