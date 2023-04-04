package loms

import (
	"context"
	"database/sql"
	"fmt"
	"route256/libs/cron"
	"time"

	"github.com/pkg/errors"
	"route256/libs/db"
	"route256/loms/internal/model"
)

func (u *useCase) CancelOrdersByTimeout(ctx context.Context) error {
	return u.db.InTx(ctx, db.RepeatableRead, func(ctxTx context.Context) error {
		orders, err := u.orderRepo.GetExpiredPaymentOrders(ctx)
		if err != nil {
			if err != sql.ErrNoRows {
				return errors.Wrap(err, "fetch expired orders")
			}
			return nil
		}
		err = u.warehouseRepo.CancelReserves(ctxTx, orders)
		if err != nil {
			return errors.Wrap(err, "cancel reservation")
		}

		err = u.orderRepo.SetOrderStatuses(ctxTx, orders, model.Cancelled)
		if err != nil {
			return errors.Wrap(err, "set order status")
		}

		for _, orderID := range orders {
			err = u.notifier.SendNotification(orderID, model.Cancelled)
			if err != nil {
				return errors.Wrap(err, "send notification")
			}
		}

		return nil
	})
}

func (u *useCase) GetCancelOrdersByTimeoutCron() cron.TaskDescriptor {
	return cron.TaskDescriptor{
		Period: time.Minute,
		Task: func(ctx context.Context) error {
			return u.CancelOrdersByTimeout(ctx)
		},
		ErrCB: func(err error) {
			fmt.Printf("Fail cancel orders by timeout: %v", err)
		},
		RetryPolicy: cron.ByScheduleAfterError,
	}
}
