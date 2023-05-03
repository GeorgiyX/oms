package checkout

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"route256/checkout/internal/convert"

	"github.com/pkg/errors"
)

func (u *useCase) Purchase(ctx context.Context, user int64) (int64, error) {
	items, err := u.repo.List(ctx, user)
	if err != nil {
		return 0, errors.Wrap(err, "fetch cart items")
	}

	if len(items) == 0 {
		return 0, status.Error(codes.InvalidArgument, "user with empty cart")
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
