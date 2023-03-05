package handlers

import (
	"context"
	"route256/loms/internal/model"
)

func (h *Handler) Stock(ctx context.Context, request model.StocksRequest) (response model.StocksResponse, err error) {
	itemInfo, err := h.useCase.Stock(ctx, request.SKU)
	if err != nil {
		return model.StocksResponse{}, err
	}

	response.Stocks = make([]model.StocksItem, 0, len(itemInfo))
	for _, info := range itemInfo {
		response.Stocks = append(response.Stocks, model.StocksItem(info))
	}
	return
}
