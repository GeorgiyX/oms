package checkout

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
	desc "route256/checkout/pkg/checkout"
)

func (h *Service) DeleteFromCart(ctx context.Context, req *desc.DeleteFromCartRequest) (*emptypb.Empty, error) {
	err := h.useCase.DeleteFromCart(ctx, req.GetUser(), req.GetSku(), req.GetCount())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil

}
