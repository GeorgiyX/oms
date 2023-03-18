package checkout

import (
	"context"
	desc "route256/checkout/pkg/checkout"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *Service) AddToCart(ctx context.Context, req *desc.AddToCartRequest) (*emptypb.Empty, error) {
	err := h.useCase.AddToCart(ctx, req.GetUser(), req.GetSku(), req.GetCount())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
