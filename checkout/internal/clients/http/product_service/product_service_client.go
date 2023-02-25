package product_service

import (
	"context"
	"net/http"

	"github.com/pkg/errors"
	"route256/checkout/internal/model"
	"route256/libs/httpaux"
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
	request := model.ProductRequest{SKU: sku}
	response, err := httpaux.Request[model.ProductRequest, model.ProductResponse](ctx, http.MethodPost, c.urlGetProduct, request)
	if err != nil {
		return nil, errors.Wrap(err, "product client")
	}

	return &model.Product{
		Name:  response.Name,
		Price: response.Price,
	}, nil
}
