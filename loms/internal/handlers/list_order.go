package handlers

import (
	"context"
	"route256/loms/internal/model"
)

func (h *Handler) ListOrder(ctx context.Context, request model.ListOrderRequest) (response model.ListOrderResponse, err error) {
	order, err := h.useCase.ListOrder(ctx, request.OrderID)
	if err != nil {
		return model.ListOrderResponse{}, err
	}
	response.Items = make([]model.ListOrderItemResponse, 0, len(order.Items))
	response.OrderID = order.OrderID
	response.User = order.User
	response.Status = order.Status
	for _, item := range order.Items {
		response.Items = append(response.Items, model.ListOrderItemResponse(item))
	}
	return
}
