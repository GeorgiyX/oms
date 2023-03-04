package service

import (
	"context"

	"route256/checkout/internal/convert"
	desc "route256/checkout/pkg/checkout"
)

func (h *Service) ListCart(ctx context.Context, req *desc.ListCartRequest) (*desc.ListCartResponse, error) {
	cart, err := h.useCase.ListCart(ctx, req.GetUser())
	if err != nil {
		return nil, err
	}

	return convert.ToCartResponse(cart), nil
}
