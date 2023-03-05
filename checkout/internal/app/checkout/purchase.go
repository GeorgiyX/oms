package checkout

import (
	"context"

	"route256/checkout/internal/convert"
	desc "route256/checkout/pkg/checkout"
)

func (h *Service) Purchase(ctx context.Context, req *desc.PurchaseRequest) (*desc.PurchaseResponse, error) {
	orderID, err := h.useCase.Purchase(ctx, req.GetUser())
	if err != nil {
		return nil, err
	}

	return convert.ToPurchaseResponse(orderID), nil
}
