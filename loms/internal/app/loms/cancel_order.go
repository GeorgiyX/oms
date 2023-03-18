package loms

import (
	"context"
	desc "route256/loms/pkg/loms"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *Service) CancelOrder(ctx context.Context, req *desc.CancelOrderRequest) (*emptypb.Empty, error) {
	err := h.useCase.CancelOrder(ctx, req.GetOrderId())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
