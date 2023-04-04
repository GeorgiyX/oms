package loms

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"route256/libs/cron"
	"route256/loms/internal/model"
)

const batchSize = 500

func (u *useCase) MarkAsSent(ctx context.Context, orderID int64) error {
	err := u.notifierOutboxRepo.SetStatus(ctx, orderID, model.PendingNotificationStatus)
	if err != nil {
		return errors.Wrap(err, "update notification fail")
	}
	return nil
}

func (u *useCase) MarkAsPending(ctx context.Context, orderID int64) error {
	err := u.notifierOutboxRepo.SetStatus(ctx, orderID, model.SendNotificationStatus)
	if err != nil {
		return errors.Wrap(err, "update notification fail")
	}
	return nil
}

func (u *useCase) NotifyPending(ctx context.Context) error {
	current := uint64(0)
	total := current + 1 // get right value later

	for current < total {
		statuses, err := u.notifierOutboxRepo.GetPendingNotifications(ctx, current, batchSize)
		total = uint64(statuses[0].Total)
		if err != nil {
			return errors.Wrap(err, "fail get pending notifications")
		}

		for _, status := range statuses {
			errIn := u.notifier.SendNotification(status.OrderID, model.OrderStatus(status.NotificationStatus))
			if errIn != nil {
				return errors.Wrap(err, "fail send notification")
			}

			errIn = u.notifierOutboxRepo.SetStatus(ctx, status.OrderID, model.WaitConfirmationNotificationStatus)
			if errIn != nil {
				return errors.Wrap(err, "fail update status")
			}
		}

		current += uint64(len(statuses))
	}

	return nil
}

func (u *useCase) GetNotifyCron() cron.TaskDescriptor {
	return cron.TaskDescriptor{
		Period: time.Minute,
		Task: func(ctx context.Context) error {
			return u.NotifyPending(ctx)
		},
		ErrCB: func(err error) {
			fmt.Printf("Fail notify in: %v", err)
		},
		RetryPolicy: cron.ByScheduleAfterError,
	}
}
