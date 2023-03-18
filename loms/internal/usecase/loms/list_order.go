package loms

import (
	"context"
	"route256/loms/internal/convert"
	"route256/loms/internal/model"

	"github.com/pkg/errors"
)

func (u *useCase) ListOrder(ctx context.Context, orderID int64) (model.Order, error) {
	order, err := u.orderRepo.GetOrderInfo(ctx, orderID)
	if err != nil {
		return model.Order{}, errors.Wrap(err, "get order info")
	}

	orderItems, err := u.orderRepo.GetOrderItems(ctx, orderID)
	if err != nil {
		return model.Order{}, errors.Wrap(err, "get order items")
	}

	return convert.ToOrder(order, orderItems), nil
}
