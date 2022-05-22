// Code generated by mockery v2.12.2. DO NOT EDIT.

package repository

import (
	context "context"

	armresources "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"

	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// mockResourcesListPager is an autogenerated mock type for the resourcesListPager type
type mockResourcesListPager struct {
	mock.Mock
}

// Err provides a mock function with given fields:
func (_m *mockResourcesListPager) Err() error {
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
func (_m *mockResourcesListPager) NextPage(ctx context.Context) bool {
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
func (_m *mockResourcesListPager) PageResponse() armresources.ResourceGroupsListResponse {
	ret := _m.Called()

	var r0 armresources.ResourceGroupsListResponse
	if rf, ok := ret.Get(0).(func() armresources.ResourceGroupsListResponse); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(armresources.ResourceGroupsListResponse)
	}

	return r0
}

// newMockResourcesListPager creates a new instance of mockResourcesListPager. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func newMockResourcesListPager(t testing.TB) *mockResourcesListPager {
	mock := &mockResourcesListPager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
