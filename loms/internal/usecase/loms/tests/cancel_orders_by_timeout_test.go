package tests

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"route256/loms/internal/model"
)

func TestCancelOrdersByTimeout(t *testing.T) {
	ordersID := []int64{
		gofakeit.Int64(),
		gofakeit.Int64(),
	}

	t.Run("should correctly cancel order", func(t *testing.T) {
		fx := tearUp(t)

		fx.mockDB(1, false)
		fx.orderRepoMock.EXPECT().GetExpiredPaymentOrders(mock.Anything).Return(ordersID, nil).Once()
		fx.warehouseRepoMock.EXPECT().CancelReserves(mock.Anything, ordersID).Return(nil).Once()
		fx.orderRepoMock.EXPECT().SetOrderStatuses(mock.Anything, ordersID, model.Cancelled).Return(nil).Once()
		fx.mockSendNotification(nil, ordersID, model.Cancelled)

		err := fx.useCase.CancelOrdersByTimeout(fx.ctx)
		require.NoError(t, err)
	})

	t.Run("should fail if SetOrderStatuses fail order", func(t *testing.T) {
		fx := tearUp(t)

		fx.mockDB(1, true)
		fx.orderRepoMock.EXPECT().GetExpiredPaymentOrders(mock.Anything).Return(ordersID, nil).Once()
		fx.warehouseRepoMock.EXPECT().CancelReserves(mock.Anything, ordersID).Return(nil).Once()
		fx.orderRepoMock.EXPECT().SetOrderStatuses(mock.Anything, ordersID, model.Cancelled).Return(errors.New(gofakeit.MinecraftFood())).Once()

		err := fx.useCase.CancelOrdersByTimeout(fx.ctx)
		require.Error(t, err)
	})

	t.Run("should fail if CancelReserves fail order", func(t *testing.T) {
		fx := tearUp(t)

		fx.mockDB(1, true)
		fx.orderRepoMock.EXPECT().GetExpiredPaymentOrders(mock.Anything).Return(ordersID, nil).Once()
		fx.warehouseRepoMock.EXPECT().CancelReserves(mock.Anything, ordersID).Return(errors.New(gofakeit.MinecraftFood())).Once()

		err := fx.useCase.CancelOrdersByTimeout(fx.ctx)
		require.Error(t, err)
	})

	t.Run("should fail if GetExpiredPaymentOrders fail order", func(t *testing.T) {
		fx := tearUp(t)

		fx.mockDB(1, true)
		fx.orderRepoMock.EXPECT().GetExpiredPaymentOrders(mock.Anything).Return(nil, errors.New(gofakeit.MinecraftFood())).Once()

		err := fx.useCase.CancelOrdersByTimeout(fx.ctx)
		require.Error(t, err)
	})
}

func TestGetCancelOrdersByTimeoutCron(t *testing.T) {
	t.Run("should return TaskDescriptor", func(t *testing.T) {
		fx := tearUp(t)
		desc := fx.useCase.GetCancelOrdersByTimeoutCron()
		require.NotNil(t, desc)
	})
}
