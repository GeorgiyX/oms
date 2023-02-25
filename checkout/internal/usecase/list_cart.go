package usecase

import (
	"context"

	"route256/checkout/internal/model"
)

func (u *useCase) ListCart(ctx context.Context, user int64) (model.Cart, error) {
	return model.Cart{}, nil
}
