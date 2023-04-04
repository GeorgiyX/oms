package notification_outbox

import (
	"context"

	"github.com/pkg/errors"
	"route256/loms/internal/model"
)

func (r *notificationOutbox) GetPendingNotifications(ctx context.Context, offset, limit uint64) ([]model.StatusChangeDatabase, error) {
	const query = `SELECT fk_order_info_id, created_at, send_status, notification_status, count(*) over () as total 
	FROM notification_outbox WHERE send_status = send_status_pending() ORDER BY created_at LIMIT $1 OFFSET $2;`

	var notifications []model.StatusChangeDatabase
	err := r.db.Select(ctx, &notifications, query, limit, offset)
	if err != nil {
		return nil, errors.Wrap(err, "cannot fetch notifications")
	}

	return notifications, nil
}

func (r *notificationOutbox) ScheduleNotification(ctx context.Context, notification model.StatusChangeDatabase) error {
	const query = `INSERT INTO notification_outbox (fk_order_info_id, created_at, send_status, notification_status) VALUES ($1, DEFAULT, $2, $3);`

	resp, err := r.db.Exec(ctx, query, notification.OrderID, notification.SendStatus, notification.NotificationStatus)
	if err != nil {
		return errors.Wrap(err, "cannot schedule notification")
	}

	affected := resp.RowsAffected()
	if affected != 1 {
		return errors.New("nothing inserted to notification_outbox")
	}

	return nil
}

func (r *notificationOutbox) SetStatus(ctx context.Context, orderID int64, notificationStatus model.NotificationStatus) error {
	const query = `UPDATE notification_outbox SET send_status = $2 WHERE fk_order_info_id = $1;`
	resp, err := r.db.Exec(ctx, query, orderID, notificationStatus)
	if err != nil {
		return errors.Wrap(err, "cannot update notification_outbox")
	}

	affected := resp.RowsAffected()
	if affected != 1 {
		return errors.New("nothing updated to notification_outbox")
	}

	return nil
}
