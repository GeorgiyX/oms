package loms

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
	desc "route256/loms/pkg/loms"
)

func (h *Service) OrderPayed(ctx context.Context, req *desc.OrderPayedRequest) (*emptypb.Empty, error) {
	err := h.useCase.OrderPayed(ctx, req.GetOrderId())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
