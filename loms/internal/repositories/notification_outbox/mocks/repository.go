// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"
	model "route256/loms/internal/model"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

type Repository_Expecter struct {
	mock *mock.Mock
}

func (_m *Repository) EXPECT() *Repository_Expecter {
	return &Repository_Expecter{mock: &_m.Mock}
}

// GetPendingNotifications provides a mock function with given fields: ctx, offset, limit
func (_m *Repository) GetPendingNotifications(ctx context.Context, offset uint64, limit uint64) ([]model.StatusChangeDatabase, error) {
	ret := _m.Called(ctx, offset, limit)

	var r0 []model.StatusChangeDatabase
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64) []model.StatusChangeDatabase); ok {
		r0 = rf(ctx, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.StatusChangeDatabase)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64, uint64) error); ok {
		r1 = rf(ctx, offset, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Repository_GetPendingNotifications_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPendingNotifications'
type Repository_GetPendingNotifications_Call struct {
	*mock.Call
}

// GetPendingNotifications is a helper method to define mock.On call
//  - ctx context.Context
//  - offset uint64
//  - limit uint64
func (_e *Repository_Expecter) GetPendingNotifications(ctx interface{}, offset interface{}, limit interface{}) *Repository_GetPendingNotifications_Call {
	return &Repository_GetPendingNotifications_Call{Call: _e.mock.On("GetPendingNotifications", ctx, offset, limit)}
}

func (_c *Repository_GetPendingNotifications_Call) Run(run func(ctx context.Context, offset uint64, limit uint64)) *Repository_GetPendingNotifications_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64), args[2].(uint64))
	})
	return _c
}

func (_c *Repository_GetPendingNotifications_Call) Return(_a0 []model.StatusChangeDatabase, _a1 error) *Repository_GetPendingNotifications_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// ScheduleNotification provides a mock function with given fields: ctx, orderID, status
func (_m *Repository) ScheduleNotification(ctx context.Context, orderID int64, status model.NotificationStatus) error {
	ret := _m.Called(ctx, orderID, status)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, model.NotificationStatus) error); ok {
		r0 = rf(ctx, orderID, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Repository_ScheduleNotification_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ScheduleNotification'
type Repository_ScheduleNotification_Call struct {
	*mock.Call
}

// ScheduleNotification is a helper method to define mock.On call
//  - ctx context.Context
//  - orderID int64
//  - status model.NotificationStatus
func (_e *Repository_Expecter) ScheduleNotification(ctx interface{}, orderID interface{}, status interface{}) *Repository_ScheduleNotification_Call {
	return &Repository_ScheduleNotification_Call{Call: _e.mock.On("ScheduleNotification", ctx, orderID, status)}
}

func (_c *Repository_ScheduleNotification_Call) Run(run func(ctx context.Context, orderID int64, status model.NotificationStatus)) *Repository_ScheduleNotification_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(model.NotificationStatus))
	})
	return _c
}

func (_c *Repository_ScheduleNotification_Call) Return(_a0 error) *Repository_ScheduleNotification_Call {
	_c.Call.Return(_a0)
	return _c
}

// SetStatus provides a mock function with given fields: ctx, orderID, notificationStatus
func (_m *Repository) SetStatus(ctx context.Context, orderID int64, notificationStatus model.NotificationStatus) error {
	ret := _m.Called(ctx, orderID, notificationStatus)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, model.NotificationStatus) error); ok {
		r0 = rf(ctx, orderID, notificationStatus)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Repository_SetStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetStatus'
type Repository_SetStatus_Call struct {
	*mock.Call
}

// SetStatus is a helper method to define mock.On call
//  - ctx context.Context
//  - orderID int64
//  - notificationStatus model.NotificationStatus
func (_e *Repository_Expecter) SetStatus(ctx interface{}, orderID interface{}, notificationStatus interface{}) *Repository_SetStatus_Call {
	return &Repository_SetStatus_Call{Call: _e.mock.On("SetStatus", ctx, orderID, notificationStatus)}
}

func (_c *Repository_SetStatus_Call) Run(run func(ctx context.Context, orderID int64, notificationStatus model.NotificationStatus)) *Repository_SetStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(model.NotificationStatus))
	})
	return _c
}

func (_c *Repository_SetStatus_Call) Return(_a0 error) *Repository_SetStatus_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t mockConstructorTestingTNewRepository) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}