package checkout

import (
	"context"
	"route256/checkout/internal/clients/loms"
	productService "route256/checkout/internal/clients/product_service"
	"route256/checkout/internal/model"
	"route256/checkout/internal/repositories/cart"
)

var _ UseCase = (*useCase)(nil)

type UseCase interface {
	AddToCart(ctx context.Context, user int64, sku uint32, count uint32) error
	DeleteFromCart(ctx context.Context, user int64, sku uint32, count uint32) error
	ListCart(ctx context.Context, user int64) (model.Cart, error)
	Purchase(ctx context.Context, user int64) (int64, error)
}

type useCase struct {
	stocksChecker   loms.StocksChecker
	productResolver productService.SkuResolver
	repo            cart.Repository
}

type Config struct {
	loms.StocksChecker
	productService.SkuResolver
	cart.Repository
}

func New(config Config) *useCase {
	return &useCase{
		stocksChecker:   config.StocksChecker,
		productResolver: config.SkuResolver,
		repo:            config.Repository,
	}
}
