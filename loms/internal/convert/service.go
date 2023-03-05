package convert

import (
	"route256/loms/internal/model"
	desc "route256/loms/pkg/loms"
)

func ToOrderItemsToCreate(in *desc.CreateOrderRequest) []model.OrderItemToCreate {
	items := make([]model.OrderItemToCreate, 0, len(in.GetItems()))
	for _, item := range in.GetItems() {
		items = append(items, model.OrderItemToCreate{
			SKU:   item.GetSku(),
			Count: item.GetCount(),
		})
	}
	return items
}

func ToOrderStatus(in model.OrderStatus) desc.OrderStatus {
	switch in {
	case model.New:
		return desc.OrderStatus_NEW
	case model.Failed:
		return desc.OrderStatus_FAILED
	case model.AwaitingPayment:
		return desc.OrderStatus_AWAITING_PAYMENT
	case model.Payed:
		return desc.OrderStatus_PAYED
	case model.Cancelled:
		return desc.OrderStatus_CANCELLED
	}
	return desc.OrderStatus_INVALID
}

func ToListOrderResponse(in *model.Order) *desc.ListOrderResponse {
	out := &desc.ListOrderResponse{
		Status:  ToOrderStatus(in.Status),
		User:    in.User,
		Items:   make([]*desc.ListOrderResponse_Item, 0, len(in.Items)),
		OrderId: in.OrderID,
	}

	for _, item := range in.Items {
		out.Items = append(out.Items, &desc.ListOrderResponse_Item{
			Sku:   item.SKU,
			Count: item.Count,
		})
	}

	return out
}

func ToStocksResponse(in []model.StocksItemInfo) *desc.StocksResponse {
	out := &desc.StocksResponse{
		Items: make([]*desc.StocksResponse_Item, 0, len(in)),
	}

	for _, info := range in {
		out.Items = append(out.Items, &desc.StocksResponse_Item{
			WarehouseId: info.WarehouseID,
			Count:       info.Count,
		})
	}

	return out
}
