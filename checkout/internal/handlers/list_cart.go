package handlers

import (
	"context"

	"route256/checkout/internal/model"
)

func (h *Handler) ListCart(ctx context.Context, request model.ListCartRequest) (response model.CartResponse, err error) {
	cart, err := h.useCase.ListCart(ctx, request.User)
	if err != nil {
		return
	}

	response.TotalPrice = cart.TotalPrice
	response.Items = make([]*model.CartItemResponse, 0, len(cart.Items))
	for _, item := range cart.Items {
		response.Items = append(response.Items, &model.CartItemResponse{
			Sku:   item.Sku,
			Count: item.Count,
			Name:  item.Name,
			Price: item.Price,
		})
	}
	return
}
