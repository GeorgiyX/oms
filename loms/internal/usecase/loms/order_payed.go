package loms

import (
	"context"
	"route256/loms/internal/model"

	"github.com/pkg/errors"
)

func (u *useCase) OrderPayed(ctx context.Context, orderID int64) error {
	err := u.orderRepo.SetOrderStatuses(ctx, []int64{orderID}, model.Payed)
	if err != nil {
		return errors.Wrap(err, "set order status")
	}

	err = u.notifier.SendNotification(orderID, model.Payed)
	if err != nil {
		return errors.Wrap(err, "send notification")
	}

	return nil
}
