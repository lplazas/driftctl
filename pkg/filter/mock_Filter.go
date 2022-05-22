// Code generated by mockery v2.12.2. DO NOT EDIT.

package filter

import (
	resource "github.com/snyk/driftctl/pkg/resource"
	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// MockFilter is an autogenerated mock type for the Filter type
type MockFilter struct {
	mock.Mock
}

// IsFieldIgnored provides a mock function with given fields: res, path
func (_m *MockFilter) IsFieldIgnored(res *resource.Resource, path []string) bool {
	ret := _m.Called(res, path)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*resource.Resource, []string) bool); ok {
		r0 = rf(res, path)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsResourceIgnored provides a mock function with given fields: res
func (_m *MockFilter) IsResourceIgnored(res *resource.Resource) bool {
	ret := _m.Called(res)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*resource.Resource) bool); ok {
		r0 = rf(res)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsTypeIgnored provides a mock function with given fields: ty
func (_m *MockFilter) IsTypeIgnored(ty resource.ResourceType) bool {
	ret := _m.Called(ty)

	var r0 bool
	if rf, ok := ret.Get(0).(func(resource.ResourceType) bool); ok {
		r0 = rf(ty)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// NewMockFilter creates a new instance of MockFilter. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockFilter(t testing.TB) *MockFilter {
	mock := &MockFilter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
