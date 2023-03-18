package convert

import "route256/loms/internal/model"

func ToOrder(order model.OrderInfo, orderItems []model.OrderItemDB) model.Order {
	out := model.Order{
		Status:  order.Status,
		User:    order.UserID,
		Items:   make([]model.OrderItem, 0, len(orderItems)),
		OrderID: order.ID,
	}

	for _, item := range orderItems {
		out.Items = append(out.Items, model.OrderItem{
			SKU:   item.Sku,
			Count: item.Count,
		})
	}

	return out
}
