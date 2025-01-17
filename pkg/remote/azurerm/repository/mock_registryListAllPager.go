// Code generated by mockery v2.12.2. DO NOT EDIT.

package repository

import (
	context "context"

	armcontainerregistry "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"

	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// mockRegistryListAllPager is an autogenerated mock type for the registryListAllPager type
type mockRegistryListAllPager struct {
	mock.Mock
}

// Err provides a mock function with given fields:
func (_m *mockRegistryListAllPager) Err() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NextPage provides a mock function with given fields: ctx
func (_m *mockRegistryListAllPager) NextPage(ctx context.Context) bool {
	ret := _m.Called(ctx)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context) bool); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// PageResponse provides a mock function with given fields:
func (_m *mockRegistryListAllPager) PageResponse() armcontainerregistry.RegistriesListResponse {
	ret := _m.Called()

	var r0 armcontainerregistry.RegistriesListResponse
	if rf, ok := ret.Get(0).(func() armcontainerregistry.RegistriesListResponse); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(armcontainerregistry.RegistriesListResponse)
	}

	return r0
}

// newMockRegistryListAllPager creates a new instance of mockRegistryListAllPager. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func newMockRegistryListAllPager(t testing.TB) *mockRegistryListAllPager {
	mock := &mockRegistryListAllPager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
