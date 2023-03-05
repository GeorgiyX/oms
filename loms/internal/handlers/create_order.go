package handlers

import (
	"context"
	"route256/loms/internal/model"
)

func (h *Handler) CreateOrder(ctx context.Context, request model.CreateOrderRequest) (response model.CreateOrderResponse, err error) {
	items := make([]model.OrderItemToCreate, 0, len(request.Items))
	for _, item := range request.Items {
		items = append(items, model.OrderItemToCreate(item))
	}
	response.OrderID, err = h.useCase.CreateOrder(ctx, request.User, items)
	if err != nil {
		return model.CreateOrderResponse{}, nil
	}
	return
}
