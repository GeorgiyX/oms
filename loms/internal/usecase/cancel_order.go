package usecase

import (
	"context"

	"github.com/pkg/errors"
	"route256/libs/db"
	"route256/loms/internal/model"
)

func (u *useCase) CancelOrder(ctx context.Context, orderID int64) error {
	return u.db.InTx(ctx, db.RepeatableRead, func(ctxTx context.Context) error {
		err := u.warehouseRepo.CancelReserve(ctxTx, orderID)
		if err != nil {
			return errors.Wrap(err, "cancel reservation")
		}

		err = u.orderRepo.SetOrderStatus(ctxTx, orderID, model.Cancelled)
		if err != nil {
			return errors.Wrap(err, "set order status")
		}

		return nil
	})
}
