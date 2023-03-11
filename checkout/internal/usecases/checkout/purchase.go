package checkout

import (
	"context"

	"github.com/pkg/errors"
	"route256/checkout/internal/convert"
)

func (u *useCase) Purchase(ctx context.Context, user int64) (int64, error) {
	items, err := u.repo.List(ctx, user)
	if err != nil {
		return 0, errors.Wrap(err, "fetch cart items")
	}

	orderID, err := u.stocksChecker.CreateOrder(ctx, user, convert.ToCreateOrderItems(items))
	if err != nil {
		return 0, errors.Wrap(err, "on create order")
	}

	// todo: тут нужна распределенная транзакция
	err = u.repo.RemoveByUser(ctx, user)
	if err != nil {
		return 0, errors.Wrap(err, "remove by user")
	}

	return orderID, nil
}
