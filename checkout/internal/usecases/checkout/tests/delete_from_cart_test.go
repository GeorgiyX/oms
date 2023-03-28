package checkout

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestDeleteFromCart(t *testing.T) {
	user := gofakeit.Int64()
	sku := gofakeit.Uint32()

	t.Run("should correctly delete from cart", func(t *testing.T) {
		fx := tearUp(t)

		fx.repoMock.EXPECT().Delete(mock.Anything, user, sku, uint32(skuCount)).Return(nil).Once()

		err := fx.useCase.DeleteFromCart(fx.ctx, user, sku, skuCount)
		require.NoError(t, err)
	})

	t.Run("should fail if delete from cart fail ", func(t *testing.T) {
		fx := tearUp(t)

		fx.repoMock.EXPECT().Delete(mock.Anything, user, sku, uint32(skuCount)).Return(errors.New(gofakeit.MinecraftBiome())).Once()

		err := fx.useCase.DeleteFromCart(fx.ctx, user, sku, skuCount)
		require.Error(t, err)
	})
}
