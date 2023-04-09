package loms

import (
	"context"
	"route256/libs/db"
	"route256/loms/internal/model"

	"github.com/pkg/errors"
)

func (u *useCase) OrderPayed(ctx context.Context, orderID int64) error {
	err := u.db.InTx(ctx, db.RepeatableRead, func(ctxTx context.Context) (err error) {
		errIn := u.orderRepo.SetOrderStatuses(ctxTx, []int64{orderID}, model.Payed)
		if errIn != nil {
			return errors.Wrap(errIn, "set order status")
		}

		errIn = u.notifierOutboxRepo.ScheduleNotification(ctxTx, orderID, model.Payed)
		if errIn != nil {
			return errors.Wrap(errIn, "schedule notification")
		}

		errIn = u.notifier.SendNotification(orderID, model.Payed)
		if errIn != nil {
			return errors.Wrap(errIn, "send notification")
		}

		return nil
	})

	if err != nil {
		return errors.Wrap(err, "mark order payed")
	}

	return nil
}
