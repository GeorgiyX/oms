package checkout

import (
	"context"
	"github.com/pkg/errors"
)

func (u *useCase) DeleteFromCart(ctx context.Context, user int64, sku uint32, count uint32) error {
	err := u.repo.Delete(ctx, user, sku, count)
	if err != nil {
		return errors.Wrap(err, "remove sku")
	}
	return nil
}
