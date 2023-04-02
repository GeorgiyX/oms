package tests

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"route256/loms/internal/convert"
	"route256/loms/internal/model"
)

func TestStock(t *testing.T) {
	sku := gofakeit.Uint32()
	warehouses := []model.Warehouse{
		{
			WarehouseID:      gofakeit.Int64(),
			Sku:              gofakeit.Uint32(),
			AvailableToOrder: gofakeit.Uint64(),
		},
	}
	stockExp := convert.ToStocksItemInfo(warehouses)

	t.Run("should get stock", func(t *testing.T) {
		fx := tearUp(t)

		fx.warehouseRepoMock.EXPECT().SkuStock(mock.Anything, sku).Return(warehouses, nil).Once()

		stockAct, err := fx.useCase.Stock(fx.ctx, sku)
		require.NoError(t, err)
		require.Equal(t, stockExp, stockAct)
	})

	t.Run("should fail if stock fail", func(t *testing.T) {
		fx := tearUp(t)

		fx.warehouseRepoMock.EXPECT().SkuStock(mock.Anything, sku).Return(nil, errors.New(gofakeit.MinecraftFood())).Once()

		stockAct, err := fx.useCase.Stock(fx.ctx, sku)
		require.Error(t, err)
		require.Nil(t, stockAct)
	})
}
