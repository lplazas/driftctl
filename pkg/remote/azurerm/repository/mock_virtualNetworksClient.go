// Code generated by mockery v2.12.2. DO NOT EDIT.

package repository

import (
	armnetwork "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// mockVirtualNetworksClient is an autogenerated mock type for the virtualNetworksClient type
type mockVirtualNetworksClient struct {
	mock.Mock
}

// ListAll provides a mock function with given fields: options
func (_m *mockVirtualNetworksClient) ListAll(options *armnetwork.VirtualNetworksListAllOptions) virtualNetworksListAllPager {
	ret := _m.Called(options)

	var r0 virtualNetworksListAllPager
	if rf, ok := ret.Get(0).(func(*armnetwork.VirtualNetworksListAllOptions) virtualNetworksListAllPager); ok {
		r0 = rf(options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(virtualNetworksListAllPager)
		}
	}

	return r0
}

// newMockVirtualNetworksClient creates a new instance of mockVirtualNetworksClient. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func newMockVirtualNetworksClient(t testing.TB) *mockVirtualNetworksClient {
	mock := &mockVirtualNetworksClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}