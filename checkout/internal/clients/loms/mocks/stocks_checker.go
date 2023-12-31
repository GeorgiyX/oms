// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	model "route256/checkout/internal/model"
)

// StocksChecker is an autogenerated mock type for the StocksChecker type
type StocksChecker struct {
	mock.Mock
}

type StocksChecker_Expecter struct {
	mock *mock.Mock
}

func (_m *StocksChecker) EXPECT() *StocksChecker_Expecter {
	return &StocksChecker_Expecter{mock: &_m.Mock}
}

// CreateOrder provides a mock function with given fields: ctx, user, item
func (_m *StocksChecker) CreateOrder(ctx context.Context, user int64, item []model.CreateOrderItem) (int64, error) {
	ret := _m.Called(ctx, user, item)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, int64, []model.CreateOrderItem) int64); ok {
		r0 = rf(ctx, user, item)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64, []model.CreateOrderItem) error); ok {
		r1 = rf(ctx, user, item)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StocksChecker_CreateOrder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateOrder'
type StocksChecker_CreateOrder_Call struct {
	*mock.Call
}

// CreateOrder is a helper method to define mock.On call
//  - ctx context.Context
//  - user int64
//  - item []model.CreateOrderItem
func (_e *StocksChecker_Expecter) CreateOrder(ctx interface{}, user interface{}, item interface{}) *StocksChecker_CreateOrder_Call {
	return &StocksChecker_CreateOrder_Call{Call: _e.mock.On("CreateOrder", ctx, user, item)}
}

func (_c *StocksChecker_CreateOrder_Call) Run(run func(ctx context.Context, user int64, item []model.CreateOrderItem)) *StocksChecker_CreateOrder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].([]model.CreateOrderItem))
	})
	return _c
}

func (_c *StocksChecker_CreateOrder_Call) Return(_a0 int64, _a1 error) *StocksChecker_CreateOrder_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Stocks provides a mock function with given fields: ctx, sku
func (_m *StocksChecker) Stocks(ctx context.Context, sku uint32) ([]model.Stock, error) {
	ret := _m.Called(ctx, sku)

	var r0 []model.Stock
	if rf, ok := ret.Get(0).(func(context.Context, uint32) []model.Stock); ok {
		r0 = rf(ctx, sku)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Stock)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint32) error); ok {
		r1 = rf(ctx, sku)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StocksChecker_Stocks_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Stocks'
type StocksChecker_Stocks_Call struct {
	*mock.Call
}

// Stocks is a helper method to define mock.On call
//  - ctx context.Context
//  - sku uint32
func (_e *StocksChecker_Expecter) Stocks(ctx interface{}, sku interface{}) *StocksChecker_Stocks_Call {
	return &StocksChecker_Stocks_Call{Call: _e.mock.On("Stocks", ctx, sku)}
}

func (_c *StocksChecker_Stocks_Call) Run(run func(ctx context.Context, sku uint32)) *StocksChecker_Stocks_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32))
	})
	return _c
}

func (_c *StocksChecker_Stocks_Call) Return(_a0 []model.Stock, _a1 error) *StocksChecker_Stocks_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewStocksChecker interface {
	mock.TestingT
	Cleanup(func())
}

// NewStocksChecker creates a new instance of StocksChecker. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStocksChecker(t mockConstructorTestingTNewStocksChecker) *StocksChecker {
	mock := &StocksChecker{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
