package product_service

import (
	"context"
	"net/http"
	"route256/checkout/internal/config"
	"route256/checkout/internal/model"
	"route256/libs/httpaux"

	"github.com/pkg/errors"
)

var _ SkuResolver = (*clientSkuResolver)(nil)

type SkuResolver interface {
	Resolve(ctx context.Context, sku uint32) (*model.Product, error)
}

type clientSkuResolver struct {
	url           string
	urlGetProduct string
}

func New(url string) *clientSkuResolver {
	return &clientSkuResolver{
		url:           url,
		urlGetProduct: url + "/get_product",
	}
}

func (c *clientSkuResolver) Resolve(ctx context.Context, sku uint32) (*model.Product, error) {
	request := model.ProductRequest{
		Token: config.Instance.Token,
		SKU:   sku,
	}
	response, err := httpaux.Request[model.ProductRequest, model.ProductResponse](ctx, http.MethodPost, c.urlGetProduct, request)
	if err != nil {
		return nil, errors.Wrap(err, "product client Resolve")
	}

	return &model.Product{
		Name:  response.Name,
		Price: response.Price,
	}, nil
}
