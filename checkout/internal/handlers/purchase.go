package handlers

import (
	"context"

	"route256/checkout/internal/model"
)

func (h *Handler) Purchase(ctx context.Context, request model.PurchaseRequest) (response model.PurchaseResponse, err error) {
	orderID, err := h.useCase.Purchase(ctx, request.User)
	if err != nil {
		return
	}
	response.OrderID = orderID
	return response, nil
}
