package handlers

import (
	"context"
	"route256/checkout/internal/model"
)

func (h *Handler) AddToCart(ctx context.Context, req model.AddToCartRequest) (response model.AddToCartResponse, err error) {
	err = h.useCase.AddToCart(ctx, req.User, req.Sku, req.Count)
	return
}
