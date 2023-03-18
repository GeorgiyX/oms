package loms

import (
	"context"
	"fmt"
	"route256/libs/db"
	"route256/loms/internal/convert"
	"route256/loms/internal/model"

	"github.com/pkg/errors"
)

var (
	ErrCantReserveSku = errors.New("not enough sku in warehouse")
)

func (u *useCase) CreateOrder(ctx context.Context, user int64, items []model.OrderItemToCreate) (int64, error) {
	var orderID int64
	err := u.db.InTx(ctx, db.RepeatableRead, func(ctxTx context.Context) (err error) {
		orderID, err = u.orderRepo.CreateOrder(ctxTx, user)
		if err != nil {
			return errors.Wrap(err, "create order")
		}

		err = u.orderRepo.AddToOrder(ctxTx, convert.ToOrderItemsDB(orderID, items), orderID)
		if err != nil {
			return errors.Wrap(err, "add to order")
		}

		for _, item := range items {
			isEnough, err := u.warehouseRepo.IsEnough(ctxTx, item.SKU, item.Count)
			if err != nil {
				return errors.Wrap(err, "check is enough")
			}
			if !isEnough {
				return ErrCantReserveSku
			}

			remainToReserve := item.Count
			for remainToReserve > 0 && err == nil {
				remainToReserve, err = u.warehouseRepo.ReserveNext(ctxTx, item.SKU, item.Count, orderID)
			}

			if err != nil {
				fmt.Printf("can't reserve sku: %v", err)
				break
			}
		}

		status := model.AwaitingPayment
		if err != nil {
			status = model.Failed
		}

		err = u.orderRepo.SetOrderStatus(ctxTx, orderID, status)
		if err != nil {
			return errors.Wrap(err, "set order status")
		}

		return nil
	})

	if err != nil {
		return 0, errors.Wrap(err, "create order")
	}

	return orderID, nil
}
