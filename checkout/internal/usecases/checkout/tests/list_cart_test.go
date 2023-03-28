package tests

import (
	"sort"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"route256/checkout/internal/model"
)

const itemCount = 10

func genItems(count int, userID int64) []model.CartItemDB {
	items := make([]model.CartItemDB, 0, count)
	for i := 0; i < count; i++ {
		items = append(items, model.CartItemDB{
			UserID: userID,
			Sku:    gofakeit.Uint32(),
			Count:  gofakeit.Uint32(),
		})
	}
	return items
}

func toCart(items []model.CartItemDB) (model.Cart, map[uint32]*model.Product) {
	cart := model.Cart{
		Items:      make([]*model.CartItem, 0, len(items)),
		TotalPrice: 0,
	}
	mapping := make(map[uint32]*model.Product, len(items))
	for _, item := range items {
		price := gofakeit.Uint32() + 1
		name := gofakeit.BeerName()
		cart.Items = append(cart.Items, &model.CartItem{
			Sku:   item.Sku,
			Count: item.Count,
			Name:  name,
			Price: price,
		})
		mapping[item.Sku] = &model.Product{
			Name:  name,
			Price: price,
		}
		cart.TotalPrice += price * item.Count
	}

	sort.Slice(cart.Items, func(i, j int) bool {
		return cart.Items[i].Sku < cart.Items[j].Sku
	})

	return cart, mapping
}

func TestListCart(t *testing.T) {
	userID := gofakeit.Int64()
	items := genItems(itemCount, userID)
	cartExp, productsMapping := toCart(items)
	emptyCartExp := model.Cart{
		Items:      nil,
		TotalPrice: 0,
	}

	t.Run("should correctly return cart items", func(t *testing.T) {
		fx := tearUp(t)
		fx.repoMock.EXPECT().List(mock.Anything, userID).Return(items, nil).Once()
		for _, item := range items {
			fx.productResolverMock.EXPECT().Resolve(mock.Anything, item.Sku).Return(productsMapping[item.Sku], nil).Once()
		}
		cartAct, err := fx.useCase.ListCart(fx.ctx, userID)

		require.NoError(t, err)
		require.NotNil(t, cartAct)
		require.Equal(t, cartExp.TotalPrice, cartAct.TotalPrice)
		require.NotNil(t, cartExp.Items)
		require.Len(t, cartAct.Items, len(cartExp.Items))
		sort.Slice(cartAct.Items, func(i, j int) bool {
			return cartAct.Items[i].Sku < cartAct.Items[j].Sku
		})
		for i := 0; i < len(cartExp.Items); i++ {
			require.Equal(t, cartExp.Items[i].Sku, cartAct.Items[i].Sku)
			require.Equal(t, cartExp.Items[i].Price, cartAct.Items[i].Price)
			require.Equal(t, cartExp.Items[i].Name, cartAct.Items[i].Name)
			require.Equal(t, cartExp.Items[i].Count, cartAct.Items[i].Count)
		}
	})

	t.Run("should fail if repo list fail", func(t *testing.T) {
		fx := tearUp(t)
		fx.repoMock.EXPECT().List(mock.Anything, userID).Return(nil, errors.New(gofakeit.Gamertag())).Once()
		cartAct, err := fx.useCase.ListCart(fx.ctx, userID)

		require.Error(t, err)
		require.Equal(t, emptyCartExp, cartAct)
	})

	t.Run("should fail if item resolve fail", func(t *testing.T) {
		fx := tearUp(t)
		fx.repoMock.EXPECT().List(mock.Anything, userID).Return(items, nil).Once()
		fx.productResolverMock.EXPECT().Resolve(mock.Anything, mock.Anything).Return(nil, errors.New(gofakeit.Gamertag()))
		cartAct, err := fx.useCase.ListCart(fx.ctx, userID)

		require.Error(t, err)
		require.Equal(t, emptyCartExp, cartAct)
	})
}
