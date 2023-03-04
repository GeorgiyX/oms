package service

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
	desc "route256/checkout/pkg/checkout"
)

func (h *Service) AddToCart(ctx context.Context, req *desc.AddToCartRequest) (*emptypb.Empty, error) {
	err := h.useCase.AddToCart(ctx, req.GetUser(), req.GetSku(), req.GetCount())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
