// Code generated by mockery v2.12.2. DO NOT EDIT.

package aws

import (
	testing "testing"

	mock "github.com/stretchr/testify/mock"
)

// MockFakeRequestFailure is an autogenerated mock type for the FakeRequestFailure type
type MockFakeRequestFailure struct {
	mock.Mock
}

// Code provides a mock function with given fields:
func (_m *MockFakeRequestFailure) Code() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Error provides a mock function with given fields:
func (_m *MockFakeRequestFailure) Error() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// HostID provides a mock function with given fields:
func (_m *MockFakeRequestFailure) HostID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Message provides a mock function with given fields:
func (_m *MockFakeRequestFailure) Message() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// OrigErr provides a mock function with given fields:
func (_m *MockFakeRequestFailure) OrigErr() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RequestID provides a mock function with given fields:
func (_m *MockFakeRequestFailure) RequestID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// StatusCode provides a mock function with given fields:
func (_m *MockFakeRequestFailure) StatusCode() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// NewMockFakeRequestFailure creates a new instance of MockFakeRequestFailure. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockFakeRequestFailure(t testing.TB) *MockFakeRequestFailure {
	mock := &MockFakeRequestFailure{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
