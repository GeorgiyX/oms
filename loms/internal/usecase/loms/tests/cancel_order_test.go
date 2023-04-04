package tests

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"route256/loms/internal/model"
)

func TestCancelOrder(t *testing.T) {
	orderID := gofakeit.Int64()

	t.Run("should correctly cancel order", func(t *testing.T) {
		fx := tearUp(t)

		fx.mockDB(1, false)
		fx.warehouseRepoMock.EXPECT().CancelReserves(mock.Anything, []int64{orderID}).Return(nil).Once()
		fx.orderRepoMock.EXPECT().SetOrderStatuses(mock.Anything, []int64{orderID}, model.Cancelled).Return(nil).Once()
		fx.notifierMock.EXPECT().SendNotification(orderID, model.Cancelled).Return(nil).Once()

		err := fx.useCase.CancelOrder(fx.ctx, orderID)
		require.NoError(t, err)
	})

	t.Run("should fail if CancelReserves fail", func(t *testing.T) {
		fx := tearUp(t)

		fx.mockDB(1, true)
		fx.warehouseRepoMock.EXPECT().CancelReserves(mock.Anything, []int64{orderID}).Return(errors.New(gofakeit.MinecraftFood())).Once()

		err := fx.useCase.CancelOrder(fx.ctx, orderID)
		require.Error(t, err)
	})

	t.Run("should fail if SetOrderStatuses fail", func(t *testing.T) {
		fx := tearUp(t)

		fx.mockDB(1, true)
		fx.warehouseRepoMock.EXPECT().CancelReserves(mock.Anything, []int64{orderID}).Return(nil).Once()
		fx.orderRepoMock.EXPECT().SetOrderStatuses(mock.Anything, []int64{orderID}, model.Cancelled).Return(errors.New(gofakeit.MinecraftFood())).Once()

		err := fx.useCase.CancelOrder(fx.ctx, orderID)
		require.Error(t, err)
	})
}
