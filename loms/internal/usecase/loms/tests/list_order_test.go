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

const orderItemsCount = 10

func genOrderItemsDB(count int) []model.OrderItemDB {
	items := make([]model.OrderItemDB, 0, count)
	for i := 0; i < count; i++ {
		items = append(items, model.OrderItemDB{
			Sku:         gofakeit.Uint32(),
			Count:       gofakeit.Uint32(),
			OrderInfoID: gofakeit.Int64(),
		})
	}
	return items
}

func TestListOrder(t *testing.T) {
	orderID := gofakeit.Int64()
	orderInfo := model.OrderInfo{
		ID:        orderID,
		UserID:    gofakeit.Int64(),
		CreatedAt: gofakeit.Date(),
		Status:    model.AwaitingPayment,
	}
	orderItems := genOrderItemsDB(orderItemsCount)
	orderExp := convert.ToOrder(orderInfo, orderItems)
	errGen := errors.New(gofakeit.MinecraftFood())

	t.Run("should list order info", func(t *testing.T) {
		fx := tearUp(t)

		fx.orderRepoMock.EXPECT().GetOrderInfo(mock.Anything, orderID).Return(orderInfo, nil).Once()
		fx.orderRepoMock.EXPECT().GetOrderItems(mock.Anything, orderID).Return(orderItems, nil).Once()

		orderAct, err := fx.useCase.ListOrder(fx.ctx, orderID)
		require.NoError(t, err)
		require.Equal(t, orderExp, orderAct)
	})

	t.Run("should fail if GetOrderItems fail", func(t *testing.T) {
		fx := tearUp(t)

		fx.orderRepoMock.EXPECT().GetOrderInfo(mock.Anything, orderID).Return(orderInfo, nil).Once()
		fx.orderRepoMock.EXPECT().GetOrderItems(mock.Anything, orderID).Return(nil, errGen).Once()

		orderAct, err := fx.useCase.ListOrder(fx.ctx, orderID)
		require.Error(t, err)
		require.Equal(t, model.Order{}, orderAct)
	})

	t.Run("should fail if GetOrderInfo fail", func(t *testing.T) {
		fx := tearUp(t)

		fx.orderRepoMock.EXPECT().GetOrderInfo(mock.Anything, orderID).Return(model.OrderInfo{}, errGen).Once()

		orderAct, err := fx.useCase.ListOrder(fx.ctx, orderID)
		require.Error(t, err)
		require.Equal(t, model.Order{}, orderAct)
	})
}
