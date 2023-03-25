package checkout

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"route256/checkout/internal/convert"
)

func TestPurchase(t *testing.T) {
	userID := gofakeit.Int64()
	orderIDExp := gofakeit.Int64()
	items := genItems(itemCount, userID)

	t.Run("should correctly create order", func(t *testing.T) {
		fx := tearUp(t)
		fx.repoMock.EXPECT().List(mock.Anything, userID).Return(items, nil).Once()
		fx.stocksCheckerMock.EXPECT().CreateOrder(mock.Anything, userID, convert.ToCreateOrderItems(items)).Return(orderIDExp, nil).Once()
		fx.repoMock.EXPECT().RemoveByUser(mock.Anything, userID).Return(nil).Once()
		orderIDAct, err := fx.useCase.Purchase(fx.ctx, userID)

		require.NoError(t, err)
		require.Equal(t, orderIDExp, orderIDAct)
	})

	t.Run("should fail if cart remove fail", func(t *testing.T) {
		fx := tearUp(t)
		fx.repoMock.EXPECT().List(mock.Anything, userID).Return(items, nil).Once()
		fx.stocksCheckerMock.EXPECT().CreateOrder(mock.Anything, userID, convert.ToCreateOrderItems(items)).Return(orderIDExp, nil).Once()
		fx.repoMock.EXPECT().RemoveByUser(mock.Anything, userID).Return(errors.New(gofakeit.Gamertag())).Once()
		orderIDAct, err := fx.useCase.Purchase(fx.ctx, userID)

		require.Error(t, err)
		require.Zero(t, orderIDAct)
	})

	t.Run("should fail if create order fail", func(t *testing.T) {
		fx := tearUp(t)
		fx.repoMock.EXPECT().List(mock.Anything, userID).Return(items, nil).Once()
		fx.stocksCheckerMock.EXPECT().CreateOrder(mock.Anything, userID, convert.ToCreateOrderItems(items)).Return(0, errors.New(gofakeit.MinecraftFood())).Once()
		orderIDAct, err := fx.useCase.Purchase(fx.ctx, userID)

		require.Error(t, err)
		require.Zero(t, orderIDAct)
	})

	t.Run("should fail if list order fail", func(t *testing.T) {
		fx := tearUp(t)
		fx.repoMock.EXPECT().List(mock.Anything, userID).Return(nil, errors.New(gofakeit.MinecraftFood())).Once()
		orderIDAct, err := fx.useCase.Purchase(fx.ctx, userID)

		require.Error(t, err)
		require.Zero(t, orderIDAct)
	})
}
