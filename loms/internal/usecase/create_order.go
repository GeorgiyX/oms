package usecase

import (
	"context"
	"route256/loms/internal/model"

	"github.com/brianvoe/gofakeit/v6"
)

func (u *useCase) CreateOrder(ctx context.Context, user int64, items []model.OrderItemToCreate) (int64, error) {
	return gofakeit.Int64(), nil
}
