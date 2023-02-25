package handlers

import (
	"context"
	"route256/loms/internal/model"
)

func (h *Handler) OrderPayed(ctx context.Context, request model.OrderPayedRequest) (response model.OrderPayedResponse, err error) {
	err = h.useCase.OrderPayed(ctx, request.OrderID)
	return
}
