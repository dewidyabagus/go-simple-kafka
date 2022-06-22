// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	context "context"
	order "learn/kafka/business/order"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// CheckExistingTransNo provides a mock function with given fields: transNo
func (_m *Repository) CheckExistingTransNo(transNo string) (bool, error) {
	ret := _m.Called(transNo)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(transNo)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(transNo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateNewOrder provides a mock function with given fields: ctx, orders
func (_m *Repository) CreateNewOrder(ctx context.Context, orders []order.Order) error {
	ret := _m.Called(ctx, orders)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []order.Order) error); ok {
		r0 = rf(ctx, orders)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetOrderByTransNo provides a mock function with given fields: ctx, transNo
func (_m *Repository) GetOrderByTransNo(ctx context.Context, transNo string) ([]order.Order, error) {
	ret := _m.Called(ctx, transNo)

	var r0 []order.Order
	if rf, ok := ret.Get(0).(func(context.Context, string) []order.Order); ok {
		r0 = rf(ctx, transNo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]order.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, transNo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
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
