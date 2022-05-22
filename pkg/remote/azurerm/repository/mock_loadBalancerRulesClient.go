// Code generated by mockery v2.12.2. DO NOT EDIT.

package repository

import (
	armnetwork "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// mockLoadBalancerRulesClient is an autogenerated mock type for the loadBalancerRulesClient type
type mockLoadBalancerRulesClient struct {
	mock.Mock
}

// List provides a mock function with given fields: _a0, _a1, _a2
func (_m *mockLoadBalancerRulesClient) List(_a0 string, _a1 string, _a2 *armnetwork.LoadBalancerLoadBalancingRulesListOptions) loadBalancerRulesListAllPager {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 loadBalancerRulesListAllPager
	if rf, ok := ret.Get(0).(func(string, string, *armnetwork.LoadBalancerLoadBalancingRulesListOptions) loadBalancerRulesListAllPager); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(loadBalancerRulesListAllPager)
		}
	}

	return r0
}

// newMockLoadBalancerRulesClient creates a new instance of mockLoadBalancerRulesClient. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func newMockLoadBalancerRulesClient(t testing.TB) *mockLoadBalancerRulesClient {
	mock := &mockLoadBalancerRulesClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
