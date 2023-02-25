package loms

import (
	"context"
	"net/http"

	"github.com/pkg/errors"
	"route256/checkout/internal/model"
	"route256/libs/httpaux"
)

var _ StocksChecker = (*clientLOMS)(nil)

type StocksChecker interface {
	Stocks(ctx context.Context, sku uint32) ([]model.Stock, error)
}

type clientLOMS struct {
	url       string
	urlStocks string
}

func New(url string) *clientLOMS {
	return &clientLOMS{
		url:       url,
		urlStocks: url + "/stocks",
	}
}

func (c *clientLOMS) Stocks(ctx context.Context, sku uint32) ([]model.Stock, error) {
	request := model.StocksRequest{SKU: sku}
	response, err := httpaux.Request[model.StocksRequest, model.StocksResponse](ctx, http.MethodPost, c.urlStocks, request)
	if err != nil {
		return nil, errors.Wrap(err, "stock client")
	}

	stocks := make([]model.Stock, 0, len(response.Stocks))
	for _, stock := range response.Stocks {
		stocks = append(stocks, model.Stock{
			WarehouseID: stock.WarehouseID,
			Count:       stock.Count,
		})
	}

	return stocks, nil
}
