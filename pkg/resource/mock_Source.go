// Code generated by mockery v2.12.2. DO NOT EDIT.

package resource

import (
	testing "testing"

	mock "github.com/stretchr/testify/mock"
)

// MockSource is an autogenerated mock type for the Source type
type MockSource struct {
	mock.Mock
}

// InternalName provides a mock function with given fields:
func (_m *MockSource) InternalName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Namespace provides a mock function with given fields:
func (_m *MockSource) Namespace() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Source provides a mock function with given fields:
func (_m *MockSource) Source() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// NewMockSource creates a new instance of MockSource. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockSource(t testing.TB) *MockSource {
	mock := &MockSource{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
