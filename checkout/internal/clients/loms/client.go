package loms

//go:generate mockery --case underscore --name StocksChecker --with-expecter

import (
	"context"
	"route256/checkout/internal/convert"
	"route256/checkout/internal/model"
	desc "route256/loms/pkg/loms"

	"github.com/pkg/errors"
)

var _ StocksChecker = (*clientLOMS)(nil)

type StocksChecker interface {
	Stocks(ctx context.Context, sku uint32) ([]model.Stock, error)
	CreateOrder(ctx context.Context, user int64, item []model.CreateOrderItem) (int64, error)
}

type clientLOMS struct {
	client desc.LomsClient
}

func New(client desc.LomsClient) *clientLOMS {
	return &clientLOMS{
		client: client,
	}
}

func (c *clientLOMS) Stocks(ctx context.Context, sku uint32) ([]model.Stock, error) {
	resp, err := c.client.Stock(ctx, &desc.StocksRequest{
		Sku: sku,
	})
	if err != nil {
		return nil, errors.Wrap(err, "stock client Stocks")
	}

	return convert.ToStocks(resp), nil
}

func (c *clientLOMS) CreateOrder(ctx context.Context, user int64, items []model.CreateOrderItem) (int64, error) {
	resp, err := c.client.CreateOrder(ctx, convert.ToCreateOrderRequest(user, items))
	if err != nil {
		return 0, errors.Wrap(err, "stock client CreateOrder")
	}

	return resp.GetOrderId(), nil
}
