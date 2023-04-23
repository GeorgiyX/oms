package product_service

//go:generate mockery --case underscore --name SkuResolver --with-expecter

import (
	"context"
	"route256/checkout/internal/config"
	"route256/checkout/internal/convert"
	"route256/checkout/internal/model"
	desc "route256/checkout/pkg/product-service"
	"route256/libs/cache"
	"strconv"

	"github.com/pkg/errors"
)

var _ SkuResolver = (*clientSkuResolver)(nil)

type SkuResolverCache cache.Cache[model.Product]

type SkuResolver interface {
	Resolve(ctx context.Context, sku uint32) (*model.Product, error)
}

type clientSkuResolver struct {
	client desc.ProductServiceClient
	cache  SkuResolverCache
}

func New(client desc.ProductServiceClient, cache SkuResolverCache) *clientSkuResolver {
	return &clientSkuResolver{
		client: client,
		cache:  cache,
	}
}

func (c *clientSkuResolver) Resolve(ctx context.Context, sku uint32) (*model.Product, error) {
	product, err := c.cache.Get(ctx, strconv.Itoa(int(sku)), c.getResolveFunc(sku))
	if err != nil {
		return nil, errors.Wrap(err, "error from cache")
	}

	return product, nil
}

func (c *clientSkuResolver) getResolveFunc(sku uint32) cache.GetFunc[model.Product] {
	return func(ctx context.Context) (*model.Product, error) {
		resp, err := c.client.GetProduct(ctx, &desc.GetProductRequest{
			Token: config.Instance.Token,
			Sku:   sku,
		})
		if err != nil {
			return nil, errors.Wrap(err, "product client Resolve")
		}

		return convert.ToProduct(resp), nil
	}
}
