package usecase

import (
	"context"
	"route256/loms/internal/model"

	"github.com/brianvoe/gofakeit/v6"
)

const ordersSize = 5

func (u *useCase) ListOrder(ctx context.Context, orderID int64) (model.Order, error) {
	order := model.Order{
		Status:  "new",
		User:    gofakeit.Int64(),
		Items:   nil,
		OrderID: gofakeit.Int64(),
	}
	order.Items = make([]model.OrderItem, 0, ordersSize)
	for i := 0; i < warehousesCount; i++ {
		order.Items = append(order.Items, model.OrderItem{
			SKU:   gofakeit.Uint32(),
			Count: gofakeit.Uint16(),
		})
	}
	return order, nil
}
