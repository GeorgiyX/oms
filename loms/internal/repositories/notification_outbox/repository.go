package notification_outbox

//go:generate mockery --case underscore --name Repository --with-expecter

import (
	"context"
	"route256/libs/db"
	"route256/loms/internal/model"
)

var _ Repository = (*notificationOutbox)(nil)

type Repository interface {
	GetPendingNotifications(ctx context.Context, offset, limit uint64) ([]model.StatusChangeDatabase, error)
	ScheduleNotification(ctx context.Context, orderID int64, status model.OrderStatus) error
	SetStatus(ctx context.Context, orderID int64, notificationStatus model.NotificationStatus) error
}

type notificationOutbox struct {
	db db.TxDB
}

func New(db db.TxDB) *notificationOutbox {
	return &notificationOutbox{
		db: db,
	}
}
