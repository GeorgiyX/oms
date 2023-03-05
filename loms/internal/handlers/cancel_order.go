package handlers

import (
	"context"
	"route256/loms/internal/model"
)

func (h *Handler) CancelOrder(ctx context.Context, request model.CancelOrderRequest) (response model.CancelOrderResponse, err error) {
	err = h.useCase.CancelOrder(ctx, request.OrderID)
	return
}
