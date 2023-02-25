package usecase

import (
	"context"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/pkg/errors"
	"route256/checkout/internal/model"
)

var cartItems = []model.CreateOrderItem{
	{
		SKU:   gofakeit.Uint32(),
		Count: gofakeit.Uint16(),
	},
	{
		SKU:   gofakeit.Uint32(),
		Count: gofakeit.Uint16(),
	},
	{
		SKU:   gofakeit.Uint32(),
		Count: gofakeit.Uint16(),
	},
}

func (u *useCase) Purchase(ctx context.Context, user int64) (int64, error) {
	orderID, err := u.stocksChecker.CreateOrder(ctx, user, cartItems)
	if err != nil {
		return 0, errors.WithMessage(err, "purchase")
	}
	return orderID, nil
}
