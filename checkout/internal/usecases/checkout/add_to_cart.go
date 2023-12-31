package checkout

import (
	"context"

	"github.com/pkg/errors"
)

var (
	ErrInsufficientStocks = errors.New("insufficient stocks")
)

func (u *useCase) AddToCart(ctx context.Context, user int64, sku uint32, count uint32) error {
	stocks, err := u.stocksChecker.Stocks(ctx, sku)
	if err != nil {
		return errors.WithMessage(err, "checking stocks")
	}

	counter := int64(count)
	for _, stock := range stocks {
		counter -= int64(stock.Count)
		if counter <= 0 {
			break
		}
	}

	if counter > 0 {
		return ErrInsufficientStocks
	}

	err = u.repo.Add(ctx, user, sku, count)
	if err != nil {
		return errors.Wrap(err, "add sku")
	}

	return nil
}
