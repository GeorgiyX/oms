package handlers

import (
	"context"

	"route256/checkout/internal/model"
)

func (h *Handler) DeleteFromCart(ctx context.Context, request model.DeleteFromCartRequest) (model.DeleteFromCartResponse, error) {
	return model.DeleteFromCartResponse{}, nil
}
