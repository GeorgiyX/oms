// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	model "route256/loms/internal/model"

	mock "github.com/stretchr/testify/mock"
)

// Notifier is an autogenerated mock type for the Notifier type
type Notifier struct {
	mock.Mock
}

type Notifier_Expecter struct {
	mock *mock.Mock
}

func (_m *Notifier) EXPECT() *Notifier_Expecter {
	return &Notifier_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields:
func (_m *Notifier) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Notifier_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type Notifier_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *Notifier_Expecter) Close() *Notifier_Close_Call {
	return &Notifier_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *Notifier_Close_Call) Run(run func()) *Notifier_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Notifier_Close_Call) Return(_a0 error) *Notifier_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

// SendNotification provides a mock function with given fields: orderID, status
func (_m *Notifier) SendNotification(orderID int64, status model.OrderStatus) error {
	ret := _m.Called(orderID, status)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64, model.OrderStatus) error); ok {
		r0 = rf(orderID, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Notifier_SendNotification_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendNotification'
type Notifier_SendNotification_Call struct {
	*mock.Call
}

// SendNotification is a helper method to define mock.On call
//  - orderID int64
//  - status model.OrderStatus
func (_e *Notifier_Expecter) SendNotification(orderID interface{}, status interface{}) *Notifier_SendNotification_Call {
	return &Notifier_SendNotification_Call{Call: _e.mock.On("SendNotification", orderID, status)}
}

func (_c *Notifier_SendNotification_Call) Run(run func(orderID int64, status model.OrderStatus)) *Notifier_SendNotification_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64), args[1].(model.OrderStatus))
	})
	return _c
}

func (_c *Notifier_SendNotification_Call) Return(_a0 error) *Notifier_SendNotification_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewNotifier interface {
	mock.TestingT
	Cleanup(func())
}

// NewNotifier creates a new instance of Notifier. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewNotifier(t mockConstructorTestingTNewNotifier) *Notifier {
	mock := &Notifier{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}