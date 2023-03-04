package usecase

import (
	"context"
	"route256/checkout/internal/model"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/pkg/errors"
)

var cartItems = []model.CreateOrderItem{
	{
		SKU:   gofakeit.Uint32(),
		Count: gofakeit.Uint32(),
	},
	{
		SKU:   gofakeit.Uint32(),
		Count: gofakeit.Uint32(),
	},
	{
		SKU:   gofakeit.Uint32(),
		Count: gofakeit.Uint32(),
	},
}

func (u *useCase) Purchase(ctx context.Context, user int64) (int64, error) {
	orderID, err := u.stocksChecker.CreateOrder(ctx, user, cartItems)
	if err != nil {
		return 0, errors.WithMessage(err, "purchase")
	}
	return orderID, nil
}
