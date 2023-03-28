package checkout

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"route256/checkout/internal/model"
	"route256/checkout/internal/usecases/checkout"
)

const (
	warehousesCount = 3
	skuCount        = 100
)

func genStocks(warehousesCount int, skuCountTotal int) []model.Stock {
	skuPerWarehouse := skuCountTotal / warehousesCount
	stocks := make([]model.Stock, 0, warehousesCount)
	for i := 0; i < warehousesCount; i++ {
		stocks = append(stocks, model.Stock{
			WarehouseID: gofakeit.Int64(),
			Count:       uint64(skuPerWarehouse),
		})
	}
	extra := skuCountTotal % warehousesCount
	stocks[len(stocks)-1].Count += uint64(extra)
	return stocks
}

func TestAddToCart(t *testing.T) {
	user := gofakeit.Int64()
	sku := gofakeit.Uint32()

	t.Run("should correctly add to cart", func(t *testing.T) {
		fx := tearUp(t)

		fx.stocksCheckerMock.EXPECT().Stocks(mock.Anything, sku).Return(genStocks(warehousesCount, skuCount+10), nil).Once()
		fx.repoMock.EXPECT().Add(mock.Anything, user, sku, uint32(skuCount)).Return(nil).Once()

		err := fx.useCase.AddToCart(fx.ctx, user, sku, skuCount)
		require.NoError(t, err)
	})

	t.Run("should fail if not enough items at warehouses", func(t *testing.T) {
		fx := tearUp(t)

		fx.stocksCheckerMock.EXPECT().Stocks(mock.Anything, sku).Return(genStocks(warehousesCount, skuCount-10), nil).Once()

		err := fx.useCase.AddToCart(fx.ctx, user, sku, skuCount)
		require.EqualError(t, err, checkout.ErrInsufficientStocks.Error())
	})

	t.Run("should fail if not stocks check fail", func(t *testing.T) {
		fx := tearUp(t)

		fx.stocksCheckerMock.EXPECT().Stocks(mock.Anything, sku).Return(genStocks(warehousesCount, skuCount+10), errors.New(gofakeit.MinecraftBiome())).Once()

		err := fx.useCase.AddToCart(fx.ctx, user, sku, skuCount)
		require.Error(t, err)
	})

	t.Run("should fail if add fail", func(t *testing.T) {
		fx := tearUp(t)

		fx.stocksCheckerMock.EXPECT().Stocks(mock.Anything, sku).Return(genStocks(warehousesCount, skuCount+10), errors.New(gofakeit.MinecraftBiome())).Once()

		err := fx.useCase.AddToCart(fx.ctx, user, sku, skuCount)
		require.Error(t, err)
	})
}
