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
	CreateOrder(ctx context.Context, user int64, item []model.CreateOrderItem) (int64, error)
}

type clientLOMS struct {
	url            string
	urlStocks      string
	urlCreateOrder string
}

func New(url string) *clientLOMS {
	return &clientLOMS{
		url:            url,
		urlStocks:      url + "/stocks",
		urlCreateOrder: url + "/createOrder",
	}
}

func (c *clientLOMS) Stocks(ctx context.Context, sku uint32) ([]model.Stock, error) {
	request := model.StocksRequest{SKU: sku}
	response, err := httpaux.Request[model.StocksRequest, model.StocksResponse](ctx, http.MethodPost, c.urlStocks, request)
	if err != nil {
		return nil, errors.Wrap(err, "stock client Stocks")
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

func (c *clientLOMS) CreateOrder(ctx context.Context, user int64, items []model.CreateOrderItem) (int64, error) {
	request := model.CreateOrderRequest{
		User:  user,
		Items: make([]model.CreateOrderRequestItem, 0, len(items)),
	}
	for _, item := range items {
		request.Items = append(request.Items, model.CreateOrderRequestItem{
			SKU:   item.SKU,
			Count: item.Count,
		})
	}

	response, err := httpaux.Request[model.CreateOrderRequest, model.CreateOrderResponse](ctx, http.MethodPost, c.urlCreateOrder, request)
	if err != nil {
		return 0, errors.Wrap(err, "stock client CreateOrder")
	}

	return response.OrderID, nil
}
