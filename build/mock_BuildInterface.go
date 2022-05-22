// Code generated by mockery v2.12.2. DO NOT EDIT.

package build

import (
	testing "testing"

	mock "github.com/stretchr/testify/mock"
)

// MockBuildInterface is an autogenerated mock type for the BuildInterface type
type MockBuildInterface struct {
	mock.Mock
}

// IsRelease provides a mock function with given fields:
func (_m *MockBuildInterface) IsRelease() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsUsageReportingEnabled provides a mock function with given fields:
func (_m *MockBuildInterface) IsUsageReportingEnabled() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// NewMockBuildInterface creates a new instance of MockBuildInterface. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockBuildInterface(t testing.TB) *MockBuildInterface {
	mock := &MockBuildInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
