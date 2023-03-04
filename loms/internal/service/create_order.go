package service

import (
	"context"
	"route256/loms/internal/convert"
	desc "route256/loms/pkg/loms"
)

func (h *Service) CreateOrder(ctx context.Context, req *desc.CreateOrderRequest) (*desc.CreateOrderResponse, error) {
	orderID, err := h.useCase.CreateOrder(ctx, req.GetUser(), convert.ToOrderItemsToCreate(req))
	if err != nil {
		return nil, err
	}
	return &desc.CreateOrderResponse{
		OrderId: orderID,
	}, nil
}
