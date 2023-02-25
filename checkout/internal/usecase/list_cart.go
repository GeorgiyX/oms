package usecase

import (
	"context"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/pkg/errors"
	"route256/checkout/internal/model"
)

var cartSKUs = []uint32{
	1076963,
	1148162,
	1625903,
	2618151,
	2956315,
	2958025,
	3596599,
	3618852,
	4288068,
	4465995,
}

func (u *useCase) ListCart(ctx context.Context, user int64) (model.Cart, error) {
	cart := model.Cart{}
	cart.Items = make([]*model.CartItem, 0, len(cartSKUs))
	for _, sku := range cartSKUs {
		product, err := u.productResolver.Resolve(ctx, sku)
		if err != nil {
			return model.Cart{}, errors.WithMessage(err, "sku resolve")
		}
		cart.Items = append(cart.Items, &model.CartItem{
			Sku:   sku,
			Count: gofakeit.Uint16(),
			Name:  product.Name,
			Price: product.Price,
		})
		cart.TotalPrice += product.Price
	}
	return cart, nil
}
