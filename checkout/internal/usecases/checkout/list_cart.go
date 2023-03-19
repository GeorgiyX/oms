package checkout

import (
	"context"
	"route256/checkout/internal/model"
	"route256/libs/workerpool"
	"sync"

	"github.com/pkg/errors"
)

const maxWorkers = 5

func (u *useCase) ListCart(ctx context.Context, user int64) (model.Cart, error) {
	items, err := u.repo.List(ctx, user)
	if err != nil {
		return model.Cart{}, errors.Wrap(err, "fetch cart items")
	}

	cart := model.Cart{}
	cart.Items = make([]*model.CartItem, 0, len(items))
	cartMutex := sync.Mutex{}
	pool := workerpool.New(maxWorkers, len(items))

	for _, item := range items {
		pool.Schedule(ctx, func(ctx context.Context) error {
			product, err := u.productResolver.Resolve(ctx, item.Sku)
			if err != nil {
				return errors.WithMessage(err, "sku resolve")
			}

			cartMutex.Lock()
			defer cartMutex.Unlock()
			cart.Items = append(cart.Items, &model.CartItem{
				Sku:   item.Sku,
				Count: item.Count,
				Name:  product.Name,
				Price: product.Price,
			})
			cart.TotalPrice += product.Price * item.Count
			return nil
		})

	}
	return cart, nil
}
