package loms

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
	desc "route256/loms/pkg/loms"
)

func (h *Service) CancelOrder(ctx context.Context, req *desc.CancelOrderRequest) (*emptypb.Empty, error) {
	err := h.useCase.CancelOrder(ctx, req.GetOrderId())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
