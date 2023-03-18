package loms

import (
	"context"
	"route256/loms/internal/convert"
	desc "route256/loms/pkg/loms"
)

func (h *Service) ListOrder(ctx context.Context, req *desc.ListOrderRequest) (*desc.ListOrderResponse, error) {
	order, err := h.useCase.ListOrder(ctx, req.GetOrderId())
	if err != nil {
		return nil, err
	}

	return convert.ToListOrderResponse(&order), nil
}
