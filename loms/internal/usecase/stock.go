package usecase

import (
	"context"

	"github.com/pkg/errors"
	"route256/loms/internal/convert"
	"route256/loms/internal/model"
)

func (u *useCase) Stock(ctx context.Context, sku uint32) ([]model.StocksItemInfo, error) {
	stocks, err := u.warehouseRepo.SkuStock(ctx, sku)
	if err != nil {
		return nil, errors.Wrap(err, "sku stock")
	}

	return convert.ToStocksItemInfo(stocks), nil
}
