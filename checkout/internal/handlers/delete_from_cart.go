package handlers

import (
	"context"

	"route256/checkout/internal/model"
)

func (h *Handler) DeleteFromCart(ctx context.Context, request model.DeleteFromCartRequest) (response model.DeleteFromCartResponse, err error) {
	err = h.useCase.DeleteFromCart(ctx, request.User, request.Sku, request.Count)
	return
}
