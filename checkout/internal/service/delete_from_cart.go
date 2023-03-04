package service

import (
	"context"
	desc "route256/checkout/pkg/checkout"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *Service) DeleteFromCart(ctx context.Context, req *desc.DeleteFromCartRequest) (*emptypb.Empty, error) {
	err := h.useCase.DeleteFromCart(ctx, req.GetUser(), req.GetSku(), req.GetCount())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil

}
