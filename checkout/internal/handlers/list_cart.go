package handlers

import (
	"context"

	"route256/checkout/internal/model"
)

func (h *Handler) ListCart(ctx context.Context, request model.ListCartRequest) (model.CartResponse, error) {
	return model.CartResponse{}, nil
}
