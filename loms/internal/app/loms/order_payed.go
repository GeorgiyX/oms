package loms

import (
	"context"
	desc "route256/loms/pkg/loms"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *Service) OrderPayed(ctx context.Context, req *desc.OrderPayedRequest) (*emptypb.Empty, error) {
	err := h.useCase.OrderPayed(ctx, req.GetOrderId())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
