package checkout

import (
	"context"
	"route256/checkout/internal/model"

	"github.com/pkg/errors"
)

func (u *useCase) ListCart(ctx context.Context, user int64) (model.Cart, error) {
	items, err := u.repo.List(ctx, user)
	if err != nil {
		return model.Cart{}, errors.Wrap(err, "fetch cart items")
	}

	cart := model.Cart{}
	cart.Items = make([]*model.CartItem, 0, len(items))
	for _, item := range items {
		product, err := u.productResolver.Resolve(ctx, item.Sku)
		if err != nil {
			return model.Cart{}, errors.WithMessage(err, "sku resolve")
		}
		cart.Items = append(cart.Items, &model.CartItem{
			Sku:   item.Sku,
			Count: item.Count,
			Name:  product.Name,
			Price: product.Price,
		})
		cart.TotalPrice += product.Price * item.Count
	}
	return cart, nil
}
