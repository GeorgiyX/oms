// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"
	db "route256/libs/db"

	mock "github.com/stretchr/testify/mock"
)

// TxDB is an autogenerated mock type for the TxDB type
type TxDB struct {
	mock.Mock
}

type TxDB_Expecter struct {
	mock *mock.Mock
}

func (_m *TxDB) EXPECT() *TxDB_Expecter {
	return &TxDB_Expecter{mock: &_m.Mock}
}

// Exec provides a mock function with given fields: ctx, query, args
func (_m *TxDB) Exec(ctx context.Context, query string, args ...interface{}) (db.RowsAffecter, error) {
	var _ca []interface{}
	_ca = append(_ca, ctx, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 db.RowsAffecter
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) db.RowsAffecter); ok {
		r0 = rf(ctx, query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(db.RowsAffecter)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, ...interface{}) error); ok {
		r1 = rf(ctx, query, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TxDB_Exec_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Exec'
type TxDB_Exec_Call struct {
	*mock.Call
}

// Exec is a helper method to define mock.On call
//  - ctx context.Context
//  - query string
//  - args ...interface{}
func (_e *TxDB_Expecter) Exec(ctx interface{}, query interface{}, args ...interface{}) *TxDB_Exec_Call {
	return &TxDB_Exec_Call{Call: _e.mock.On("Exec",
		append([]interface{}{ctx, query}, args...)...)}
}

func (_c *TxDB_Exec_Call) Run(run func(ctx context.Context, query string, args ...interface{})) *TxDB_Exec_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *TxDB_Exec_Call) Return(_a0 db.RowsAffecter, _a1 error) *TxDB_Exec_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Get provides a mock function with given fields: ctx, dst, query, args
func (_m *TxDB) Get(ctx context.Context, dst interface{}, query string, args ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, ctx, dst, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, string, ...interface{}) error); ok {
		r0 = rf(ctx, dst, query, args...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TxDB_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type TxDB_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//  - ctx context.Context
//  - dst interface{}
//  - query string
//  - args ...interface{}
func (_e *TxDB_Expecter) Get(ctx interface{}, dst interface{}, query interface{}, args ...interface{}) *TxDB_Get_Call {
	return &TxDB_Get_Call{Call: _e.mock.On("Get",
		append([]interface{}{ctx, dst, query}, args...)...)}
}

func (_c *TxDB_Get_Call) Run(run func(ctx context.Context, dst interface{}, query string, args ...interface{})) *TxDB_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-3)
		for i, a := range args[3:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(interface{}), args[2].(string), variadicArgs...)
	})
	return _c
}

func (_c *TxDB_Get_Call) Return(_a0 error) *TxDB_Get_Call {
	_c.Call.Return(_a0)
	return _c
}

// InTx provides a mock function with given fields: ctx, lvl, fx
func (_m *TxDB) InTx(ctx context.Context, lvl db.TxLevel, fx func(context.Context) error) error {
	ret := _m.Called(ctx, lvl, fx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, db.TxLevel, func(context.Context) error) error); ok {
		r0 = rf(ctx, lvl, fx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TxDB_InTx_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'InTx'
type TxDB_InTx_Call struct {
	*mock.Call
}

// InTx is a helper method to define mock.On call
//  - ctx context.Context
//  - lvl db.TxLevel
//  - fx func(context.Context) error
func (_e *TxDB_Expecter) InTx(ctx interface{}, lvl interface{}, fx interface{}) *TxDB_InTx_Call {
	return &TxDB_InTx_Call{Call: _e.mock.On("InTx", ctx, lvl, fx)}
}

func (_c *TxDB_InTx_Call) Run(run func(ctx context.Context, lvl db.TxLevel, fx func(context.Context) error)) *TxDB_InTx_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(db.TxLevel), args[2].(func(context.Context) error))
	})
	return _c
}

func (_c *TxDB_InTx_Call) Return(_a0 error) *TxDB_InTx_Call {
	_c.Call.Return(_a0)
	return _c
}

// Select provides a mock function with given fields: ctx, dst, query, args
func (_m *TxDB) Select(ctx context.Context, dst interface{}, query string, args ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, ctx, dst, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, string, ...interface{}) error); ok {
		r0 = rf(ctx, dst, query, args...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TxDB_Select_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Select'
type TxDB_Select_Call struct {
	*mock.Call
}

// Select is a helper method to define mock.On call
//  - ctx context.Context
//  - dst interface{}
//  - query string
//  - args ...interface{}
func (_e *TxDB_Expecter) Select(ctx interface{}, dst interface{}, query interface{}, args ...interface{}) *TxDB_Select_Call {
	return &TxDB_Select_Call{Call: _e.mock.On("Select",
		append([]interface{}{ctx, dst, query}, args...)...)}
}

func (_c *TxDB_Select_Call) Run(run func(ctx context.Context, dst interface{}, query string, args ...interface{})) *TxDB_Select_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-3)
		for i, a := range args[3:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(interface{}), args[2].(string), variadicArgs...)
	})
	return _c
}

func (_c *TxDB_Select_Call) Return(_a0 error) *TxDB_Select_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewTxDB interface {
	mock.TestingT
	Cleanup(func())
}

// NewTxDB creates a new instance of TxDB. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTxDB(t mockConstructorTestingTNewTxDB) *TxDB {
	mock := &TxDB{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}