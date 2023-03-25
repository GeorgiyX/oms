package product_service

//go:generate mockery --case underscore --name SkuResolver --with-expecter

import (
	"context"
	"route256/checkout/internal/config"
	"route256/checkout/internal/convert"
	"route256/checkout/internal/model"
	desc "route256/checkout/pkg/product-service"

	"github.com/pkg/errors"
)

var _ SkuResolver = (*clientSkuResolver)(nil)

type SkuResolver interface {
	Resolve(ctx context.Context, sku uint32) (*model.Product, error)
}

type clientSkuResolver struct {
	client desc.ProductServiceClient
}

func New(client desc.ProductServiceClient) *clientSkuResolver {
	return &clientSkuResolver{
		client: client,
	}
}

func (c *clientSkuResolver) Resolve(ctx context.Context, sku uint32) (*model.Product, error) {
	resp, err := c.client.GetProduct(ctx, &desc.GetProductRequest{
		Token: config.Instance.Token,
		Sku:   sku,
	})
	if err != nil {
		return nil, errors.Wrap(err, "product client Resolve")
	}

	return convert.ToProduct(resp), nil
}
