package tests

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"route256/loms/internal/model"
)

func TestOrderPayed(t *testing.T) {
	orderID := gofakeit.Int64()

	t.Run("should mark order payed", func(t *testing.T) {
		fx := tearUp(t)

		fx.orderRepoMock.EXPECT().SetOrderStatuses(mock.Anything, []int64{orderID}, model.Payed).Return(nil).Once()
		fx.notifierMock.EXPECT().SendNotification(orderID, model.Payed).Return(nil).Once()

		err := fx.useCase.OrderPayed(fx.ctx, orderID)
		require.NoError(t, err)
	})

	t.Run("should fail if SetOrderStatuses fail", func(t *testing.T) {
		fx := tearUp(t)

		fx.orderRepoMock.EXPECT().SetOrderStatuses(mock.Anything, []int64{orderID}, model.Payed).Return(errors.New(gofakeit.MinecraftFood())).Once()

		err := fx.useCase.OrderPayed(fx.ctx, orderID)
		require.Error(t, err)
	})
}
