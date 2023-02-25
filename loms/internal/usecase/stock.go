package usecase

import (
	"context"
	"route256/loms/internal/model"

	"github.com/brianvoe/gofakeit/v6"
)

const (
	warehousesCount = 3
)

func (u *useCase) Stock(ctx context.Context, sku uint32) ([]model.StocksItemInfo, error) {
	items := make([]model.StocksItemInfo, 0, warehousesCount)
	for i := 0; i < warehousesCount; i++ {
		items = append(items, model.StocksItemInfo{
			WarehouseID: gofakeit.Int64(),
			Count:       gofakeit.Uint64(),
		})
	}
	return items, nil
}
