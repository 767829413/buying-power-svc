// Code generated by mockery v1.0.0. DO NOT EDIT.

package copart

import context "context"
import mock "github.com/stretchr/testify/mock"

// MockLiveBidderer is an autogenerated mock type for the LiveBidderer type
type MockLiveBidderer struct {
	mock.Mock
}

// Connector provides a mock function with given fields: ctx
func (_m *MockLiveBidderer) Connector(ctx context.Context) (Connector, error) {
	ret := _m.Called(ctx)

	var r0 Connector
	if rf, ok := ret.Get(0).(func(context.Context) Connector); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Connector)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetConfig provides a mock function with given fields:
func (_m *MockLiveBidderer) GetConfig() Config {
	ret := _m.Called()

	var r0 Config
	if rf, ok := ret.Get(0).(func() Config); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(Config)
	}

	return r0
}

// LiveBidder provides a mock function with given fields: ctx
func (_m *MockLiveBidderer) LiveBidder(ctx context.Context) (LiveBidder, error) {
	ret := _m.Called(ctx)

	var r0 LiveBidder
	if rf, ok := ret.Get(0).(func(context.Context) LiveBidder); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(LiveBidder)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
