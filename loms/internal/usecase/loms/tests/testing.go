package tests

import (
	"context"
	mocksDB "route256/libs/db/mocks"
	"route256/loms/internal/model"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"route256/libs/db"
	mocksNotifier "route256/loms/internal/notifier/mocks"
	mocksOrderRepo "route256/loms/internal/repositories/order/mocks"
	mocksWarehouseRepo "route256/loms/internal/repositories/warehouse/mocks"
	"route256/loms/internal/usecase/loms"
)

type fixture struct {
	t                 *testing.T
	ctx               context.Context
	warehouseRepoMock *mocksWarehouseRepo.Repository
	orderRepoMock     *mocksOrderRepo.Repository
	notifierMock      *mocksNotifier.Notifier
	dbMock            *mocksDB.TxDB
	useCase           loms.UseCase
}

func tearUp(t *testing.T) *fixture {
	fx := &fixture{
		t:                 t,
		ctx:               context.Background(),
		warehouseRepoMock: mocksWarehouseRepo.NewRepository(t),
		orderRepoMock:     mocksOrderRepo.NewRepository(t),
		notifierMock:      mocksNotifier.NewNotifier(t),
		dbMock:            mocksDB.NewTxDB(t),
		useCase:           nil,
	}

	fx.useCase = loms.New(loms.Config{
		WarehouseRepository: fx.warehouseRepoMock,
		OrderRepository:     fx.orderRepoMock,
		Notifier:            fx.notifierMock,
		TxDB:                fx.dbMock,
	})

	return fx
}

func (fx *fixture) mockSendNotification(err error, orderIDs []int64, status model.OrderStatus) {
	for _, orderID := range orderIDs {
		fx.notifierMock.EXPECT().SendNotification(orderID, status).Return(err).Once()
	}
}

func (fx *fixture) mockDB(times int, isErr bool) {
	fx.dbMock.EXPECT().InTx(
		mock.Anything,
		mock.AnythingOfType("db.TxLevel"),
		mock.AnythingOfType("func(context.Context) error")).Run(
		func(ctx context.Context, lvl db.TxLevel, fn func(context.Context) error) {
			err := fn(fx.ctx)
			if isErr {
				require.Error(fx.t, err)
				return
			}
			require.NoError(fx.t, err)
		}).Return(errIf(isErr)).Times(times)
}

func errIf(isErr bool) error {
	if isErr {
		return errors.New(gofakeit.MinecraftFood())
	}
	return nil
}
