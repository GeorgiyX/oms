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

const itemsCount = 10

func genOrderItemsToCreate(count int) []model.OrderItemToCreate {
	items := make([]model.OrderItemToCreate, 0, count)
	for i := 0; i < count; i++ {
		items = append(items, model.OrderItemToCreate{
			SKU:   gofakeit.Uint32(),
			Count: gofakeit.Uint32(),
		})
	}
	return items
}

func TestCreateOrder(t *testing.T) {
	orderID := gofakeit.Int64()
	user := gofakeit.Int64()
	items := genOrderItemsToCreate(itemsCount)
	errGen := errors.New(gofakeit.MinecraftFood())

	t.Run("should correctly create order", func(t *testing.T) {
		fx := tearUp(t)

		fx.mockDB(1, false)
		fx.orderRepoMock.EXPECT().CreateOrder(mock.Anything, user).Return(orderID, nil).Once()
		fx.orderRepoMock.EXPECT().AddToOrder(mock.Anything, convert.ToOrderItemsDB(orderID, items), orderID).Return(nil).Once()
		for _, item := range items {
			fx.warehouseRepoMock.EXPECT().IsEnough(mock.Anything, item.SKU, item.Count).Return(true, nil).Once()
			fx.warehouseRepoMock.EXPECT().ReserveNext(mock.Anything, item.SKU, item.Count, orderID).Return(0, nil).Once()
		}
		fx.orderRepoMock.EXPECT().SetOrderStatuses(mock.Anything, []int64{orderID}, model.AwaitingPayment).Return(nil).Once()

		actOrderId, err := fx.useCase.CreateOrder(fx.ctx, user, items)
		require.NoError(t, err)
		require.Equal(t, orderID, actOrderId)
	})

	t.Run("should fail if SetOrderStatuses fail", func(t *testing.T) {
		fx := tearUp(t)

		fx.mockDB(1, true)
		fx.orderRepoMock.EXPECT().CreateOrder(mock.Anything, user).Return(orderID, nil).Once()
		fx.orderRepoMock.EXPECT().AddToOrder(mock.Anything, convert.ToOrderItemsDB(orderID, items), orderID).Return(nil).Once()
		for _, item := range items {
			fx.warehouseRepoMock.EXPECT().IsEnough(mock.Anything, item.SKU, item.Count).Return(true, nil).Once()
			fx.warehouseRepoMock.EXPECT().ReserveNext(mock.Anything, item.SKU, item.Count, orderID).Return(0, nil).Once()
		}
		fx.orderRepoMock.EXPECT().SetOrderStatuses(mock.Anything, []int64{orderID}, model.AwaitingPayment).Return(errGen).Once()

		actOrderId, err := fx.useCase.CreateOrder(fx.ctx, user, items)
		require.Error(t, err)
		require.Zero(t, actOrderId)
	})

	t.Run("should set Failed status if ReserveNext fail", func(t *testing.T) {
		fx := tearUp(t)

		fx.mockDB(1, false)
		fx.orderRepoMock.EXPECT().CreateOrder(mock.Anything, user).Return(orderID, nil).Once()
		fx.orderRepoMock.EXPECT().AddToOrder(mock.Anything, convert.ToOrderItemsDB(orderID, items), orderID).Return(nil).Once()
		fx.warehouseRepoMock.EXPECT().IsEnough(mock.Anything, items[0].SKU, items[0].Count).Return(true, nil).Once()
		fx.warehouseRepoMock.EXPECT().ReserveNext(mock.Anything, items[0].SKU, items[0].Count, orderID).Return(0, errGen).Once()
		fx.orderRepoMock.EXPECT().SetOrderStatuses(mock.Anything, []int64{orderID}, model.Failed).Return(nil).Once()

		actOrderId, err := fx.useCase.CreateOrder(fx.ctx, user, items)
		require.NoError(t, err)
		require.Equal(t, orderID, actOrderId)
	})
}
