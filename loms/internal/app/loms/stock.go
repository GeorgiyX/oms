package loms

import (
	"context"

	"route256/loms/internal/convert"
	desc "route256/loms/pkg/loms"
)

func (h *Service) Stock(ctx context.Context, req *desc.StocksRequest) (*desc.StocksResponse, error) {
	itemInfo, err := h.useCase.Stock(ctx, req.GetSku())
	if err != nil {
		return nil, err
	}

	return convert.ToStocksResponse(itemInfo), nil
}
