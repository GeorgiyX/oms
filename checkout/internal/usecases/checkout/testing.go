package checkout

import (
	"context"
	"testing"

	mockLoms "route256/checkout/internal/clients/loms/mocks"
	mocksProductService "route256/checkout/internal/clients/product_service/mocks"
	mocksRepository "route256/checkout/internal/repositories/cart/mocks"
)

type fixture struct {
	t                   *testing.T
	ctx                 context.Context
	stocksCheckerMock   *mockLoms.StocksChecker
	productResolverMock *mocksProductService.SkuResolver
	repoMock            *mocksRepository.Repository
	facade              *useCase
}

func tearUp(t *testing.T) *fixture {
	fx := &fixture{
		t:                   t,
		ctx:                 context.Background(),
		stocksCheckerMock:   mockLoms.NewStocksChecker(t),
		productResolverMock: mocksProductService.NewSkuResolver(t),
		repoMock:            mocksRepository.NewRepository(t),
		facade:              nil,
	}

	fx.facade = New(Config{
		StocksChecker: fx.stocksCheckerMock,
		SkuResolver:   fx.productResolverMock,
		Repository:    fx.repoMock,
	})

	return fx
}
