package usecase

import (
	"context"
)

func (u *useCase) DeleteFromCart(ctx context.Context, user int64, sku uint32, count uint32) error {
	return nil
}
