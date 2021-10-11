package repository

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/network/armnetwork"
	"github.com/cloudskiff/driftctl/pkg/remote/cache"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_ListAllVirtualNetwork_MultiplesResults(t *testing.T) {

	expected := []*armnetwork.VirtualNetwork{
		{
			Resource: armnetwork.Resource{
				ID: to.StringPtr("network1"),
			},
		},
		{
			Resource: armnetwork.Resource{
				ID: to.StringPtr("network2"),
			},
		},
		{
			Resource: armnetwork.Resource{
				ID: to.StringPtr("network3"),
			},
		},
		{
			Resource: armnetwork.Resource{
				ID: to.StringPtr("network4"),
			},
		},
	}

	fakeClient := &mockVirtualNetworkClient{}

	mockPager := &mockVirtualNetworksListAllPager{}
	mockPager.On("Err").Return(nil).Times(3)
	mockPager.On("NextPage", mock.Anything).Return(true).Times(2)
	mockPager.On("NextPage", mock.Anything).Return(false).Times(1)
	mockPager.On("PageResponse").Return(armnetwork.VirtualNetworksListAllResponse{
		VirtualNetworksListAllResult: armnetwork.VirtualNetworksListAllResult{
			VirtualNetworkListResult: armnetwork.VirtualNetworkListResult{
				Value: []*armnetwork.VirtualNetwork{
					{
						Resource: armnetwork.Resource{
							ID: to.StringPtr("network1"),
						},
					},
					{
						Resource: armnetwork.Resource{
							ID: to.StringPtr("network2"),
						},
					},
				},
			},
		},
	}).Times(1)
	mockPager.On("PageResponse").Return(armnetwork.VirtualNetworksListAllResponse{
		VirtualNetworksListAllResult: armnetwork.VirtualNetworksListAllResult{
			VirtualNetworkListResult: armnetwork.VirtualNetworkListResult{
				Value: []*armnetwork.VirtualNetwork{
					{
						Resource: armnetwork.Resource{
							ID: to.StringPtr("network3"),
						},
					},
					{
						Resource: armnetwork.Resource{
							ID: to.StringPtr("network4"),
						},
					},
				},
			},
		},
	}).Times(1)

	fakeClient.On("ListAll", mock.Anything).Return(mockPager)

	c := &cache.MockCache{}
	c.On("GetAndLock", "ListAllVirtualNetworks").Return(nil).Times(1)
	c.On("Unlock", "ListAllVirtualNetworks").Times(1)
	c.On("Put", "ListAllVirtualNetworks", expected).Return(true).Times(1)
	s := &networkRepository{
		virtualNetworksClient: fakeClient,
		cache:                 c,
	}
	got, err := s.ListAllVirtualNetworks()
	if err != nil {
		t.Errorf("ListAllVirtualNetworks() error = %v", err)
		return
	}

	mockPager.AssertExpectations(t)
	fakeClient.AssertExpectations(t)
	c.AssertExpectations(t)

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("ListAllVirtualNetworks() got = %v, want %v", got, expected)
	}
}

func Test_ListAllVirtualNetwork_MultiplesResults_WithCache(t *testing.T) {

	expected := []*armnetwork.VirtualNetwork{
		{
			Resource: armnetwork.Resource{
				ID: to.StringPtr("network3"),
			},
		},
	}

	fakeClient := &mockVirtualNetworkClient{}

	c := &cache.MockCache{}
	c.On("GetAndLock", "ListAllVirtualNetworks").Return(expected).Times(1)
	c.On("Unlock", "ListAllVirtualNetworks").Times(1)
	s := &networkRepository{
		virtualNetworksClient: fakeClient,
		cache:                 c,
	}
	got, err := s.ListAllVirtualNetworks()
	if err != nil {
		t.Errorf("ListAllVirtualNetworks() error = %v", err)
		return
	}

	fakeClient.AssertExpectations(t)
	c.AssertExpectations(t)

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("ListAllVirtualNetworks() got = %v, want %v", got, expected)
	}
}

func Test_ListAllVirtualNetwork_Error_OnPageResponse(t *testing.T) {

	fakeClient := &mockVirtualNetworkClient{}

	expectedErr := errors.New("unexpected error")

	mockPager := &mockVirtualNetworksListAllPager{}
	mockPager.On("Err").Return(expectedErr).Times(1)
	mockPager.On("NextPage", mock.Anything).Return(true).Times(1)
	mockPager.On("PageResponse").Return(armnetwork.VirtualNetworksListAllResponse{}).Times(1)

	fakeClient.On("ListAll", mock.Anything).Return(mockPager)

	s := &networkRepository{
		virtualNetworksClient: fakeClient,
		cache:                 cache.New(0),
	}
	got, err := s.ListAllVirtualNetworks()

	mockPager.AssertExpectations(t)
	fakeClient.AssertExpectations(t)

	assert.Equal(t, expectedErr, err)
	assert.Nil(t, got)
}

func Test_ListAllVirtualNetwork_Error(t *testing.T) {

	fakeClient := &mockVirtualNetworkClient{}

	expectedErr := errors.New("unexpected error")

	mockPager := &mockVirtualNetworksListAllPager{}
	mockPager.On("Err").Return(expectedErr).Times(1)
	mockPager.On("NextPage", mock.Anything).Return(false).Times(1)

	fakeClient.On("ListAll", mock.Anything).Return(mockPager)

	s := &networkRepository{
		virtualNetworksClient: fakeClient,
		cache:                 cache.New(0),
	}
	got, err := s.ListAllVirtualNetworks()

	mockPager.AssertExpectations(t)
	fakeClient.AssertExpectations(t)

	assert.Equal(t, expectedErr, err)
	assert.Nil(t, got)
}

func Test_ListAllRouteTables_MultiplesResults(t *testing.T) {

	expected := []*armnetwork.RouteTable{
		{
			Resource: armnetwork.Resource{
				ID: to.StringPtr("table1"),
			},
		},
		{
			Resource: armnetwork.Resource{
				ID: to.StringPtr("table2"),
			},
		},
		{
			Resource: armnetwork.Resource{
				ID: to.StringPtr("table3"),
			},
		},
		{
			Resource: armnetwork.Resource{
				ID: to.StringPtr("table4"),
			},
		},
	}

	fakeClient := &mockRouteTablesClient{}

	mockPager := &mockRouteTablesListAllPager{}
	mockPager.On("Err").Return(nil).Times(3)
	mockPager.On("NextPage", mock.Anything).Return(true).Times(2)
	mockPager.On("NextPage", mock.Anything).Return(false).Times(1)
	mockPager.On("PageResponse").Return(armnetwork.RouteTablesListAllResponse{
		RouteTablesListAllResult: armnetwork.RouteTablesListAllResult{
			RouteTableListResult: armnetwork.RouteTableListResult{
				Value: expected[:2],
			},
		},
	}).Times(1)
	mockPager.On("PageResponse").Return(armnetwork.RouteTablesListAllResponse{
		RouteTablesListAllResult: armnetwork.RouteTablesListAllResult{
			RouteTableListResult: armnetwork.RouteTableListResult{
				Value: expected[2:],
			},
		},
	}).Times(1)

	fakeClient.On("ListAll", mock.Anything).Return(mockPager)

	c := &cache.MockCache{}
	c.On("GetAndLock", "ListAllRouteTables").Return(nil).Times(1)
	c.On("Unlock", "ListAllRouteTables").Times(1)
	c.On("Put", "ListAllRouteTables", expected).Return(true).Times(1)
	s := &networkRepository{
		routeTableClient: fakeClient,
		cache:            c,
	}
	got, err := s.ListAllRouteTables()
	if err != nil {
		t.Errorf("ListAllRouteTables() error = %v", err)
		return
	}

	mockPager.AssertExpectations(t)
	fakeClient.AssertExpectations(t)
	c.AssertExpectations(t)

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("ListAllRouteTables() got = %v, want %v", got, expected)
	}
}

func Test_ListAllRouteTables_MultiplesResults_WithCache(t *testing.T) {

	expected := []*armnetwork.RouteTable{
		{
			Resource: armnetwork.Resource{
				ID: to.StringPtr("table1"),
			},
		},
	}

	fakeClient := &mockRouteTablesClient{}

	c := &cache.MockCache{}
	c.On("GetAndLock", "ListAllRouteTables").Return(expected).Times(1)
	c.On("Unlock", "ListAllRouteTables").Times(1)
	s := &networkRepository{
		routeTableClient: fakeClient,
		cache:            c,
	}
	got, err := s.ListAllRouteTables()
	if err != nil {
		t.Errorf("ListAllRouteTables() error = %v", err)
		return
	}

	fakeClient.AssertExpectations(t)
	c.AssertExpectations(t)

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("ListAllRouteTables() got = %v, want %v", got, expected)
	}
}

func Test_ListAllRouteTables_Error_OnPageResponse(t *testing.T) {

	fakeClient := &mockRouteTablesClient{}

	expectedErr := errors.New("unexpected error")

	mockPager := &mockRouteTablesListAllPager{}
	mockPager.On("Err").Return(expectedErr).Times(1)
	mockPager.On("NextPage", mock.Anything).Return(true).Times(1)
	mockPager.On("PageResponse").Return(armnetwork.RouteTablesListAllResponse{}).Times(1)

	fakeClient.On("ListAll", mock.Anything).Return(mockPager)

	s := &networkRepository{
		routeTableClient: fakeClient,
		cache:            cache.New(0),
	}
	got, err := s.ListAllRouteTables()

	mockPager.AssertExpectations(t)
	fakeClient.AssertExpectations(t)

	assert.Equal(t, expectedErr, err)
	assert.Nil(t, got)
}

func Test_ListAllRouteTables_Error(t *testing.T) {

	fakeClient := &mockRouteTablesClient{}

	expectedErr := errors.New("unexpected error")

	mockPager := &mockRouteTablesListAllPager{}
	mockPager.On("Err").Return(expectedErr).Times(1)
	mockPager.On("NextPage", mock.Anything).Return(false).Times(1)

	fakeClient.On("ListAll", mock.Anything).Return(mockPager)

	s := &networkRepository{
		routeTableClient: fakeClient,
		cache:            cache.New(0),
	}
	got, err := s.ListAllRouteTables()

	mockPager.AssertExpectations(t)
	fakeClient.AssertExpectations(t)

	assert.Equal(t, expectedErr, err)
	assert.Nil(t, got)
}

func Test_ListAllSubnets_MultiplesResults(t *testing.T) {

	network := &armnetwork.VirtualNetwork{
		Resource: armnetwork.Resource{
			Name: to.StringPtr("network1"),
			ID:   to.StringPtr("/subscriptions/7bfb2c5c-0000-0000-0000-fffa356eb406/resourceGroups/test-dev/providers/Microsoft.Network/virtualNetworks/network1"),
		},
	}

	expected := []*armnetwork.Subnet{
		{
			SubResource: armnetwork.SubResource{
				ID: to.StringPtr("subnet1"),
			},
		},
		{
			SubResource: armnetwork.SubResource{
				ID: to.StringPtr("subnet2"),
			},
		},
		{
			SubResource: armnetwork.SubResource{
				ID: to.StringPtr("subnet3"),
			},
		},
		{
			SubResource: armnetwork.SubResource{
				ID: to.StringPtr("subnet4"),
			},
		},
	}

	fakeClient := &mockSubnetsClient{}

	mockPager := &mockSubnetsListPager{}
	mockPager.On("Err").Return(nil).Times(3)
	mockPager.On("NextPage", mock.Anything).Return(true).Times(2)
	mockPager.On("NextPage", mock.Anything).Return(false).Times(1)
	mockPager.On("PageResponse").Return(armnetwork.SubnetsListResponse{
		SubnetsListResult: armnetwork.SubnetsListResult{
			SubnetListResult: armnetwork.SubnetListResult{
				Value: []*armnetwork.Subnet{
					{
						SubResource: armnetwork.SubResource{
							ID: to.StringPtr("subnet1"),
						},
					},
					{
						SubResource: armnetwork.SubResource{
							ID: to.StringPtr("subnet2"),
						},
					},
				},
			},
		},
	}).Times(1)
	mockPager.On("PageResponse").Return(armnetwork.SubnetsListResponse{
		SubnetsListResult: armnetwork.SubnetsListResult{
			SubnetListResult: armnetwork.SubnetListResult{
				Value: []*armnetwork.Subnet{
					{
						SubResource: armnetwork.SubResource{
							ID: to.StringPtr("subnet3"),
						},
					},
					{
						SubResource: armnetwork.SubResource{
							ID: to.StringPtr("subnet4"),
						},
					},
				},
			},
		},
	}).Times(1)

	fakeClient.On("List", "test-dev", "network1", mock.Anything).Return(mockPager)

	c := &cache.MockCache{}
	cacheKey := fmt.Sprintf("ListAllSubnets_%s", *network.ID)
	c.On("Get", cacheKey).Return(nil).Times(1)
	c.On("Put", cacheKey, expected).Return(true).Times(1)
	s := &networkRepository{
		subnetsClient: fakeClient,
		cache:         c,
	}
	got, err := s.ListAllSubnets(network)
	if err != nil {
		t.Errorf("ListAllSubnets() error = %v", err)
		return
	}

	mockPager.AssertExpectations(t)
	fakeClient.AssertExpectations(t)
	c.AssertExpectations(t)

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("ListAllSubnets() got = %v, want %v", got, expected)
	}
}

func Test_ListAllSubnets_MultiplesResults_WithCache(t *testing.T) {

	network := &armnetwork.VirtualNetwork{
		Resource: armnetwork.Resource{
			ID: to.StringPtr("networkID"),
		},
	}

	expected := []*armnetwork.Subnet{
		{
			Name: to.StringPtr("network1"),
		},
	}
	fakeClient := &mockSubnetsClient{}

	c := &cache.MockCache{}
	c.On("Get", "ListAllSubnets_networkID").Return(expected).Times(1)
	s := &networkRepository{
		subnetsClient: fakeClient,
		cache:         c,
	}
	got, err := s.ListAllSubnets(network)
	if err != nil {
		t.Errorf("ListAllSubnets() error = %v", err)
		return
	}

	fakeClient.AssertExpectations(t)
	c.AssertExpectations(t)

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("ListAllSubnets() got = %v, want %v", got, expected)
	}
}

func Test_ListAllSubnets_Error_OnPageResponse(t *testing.T) {

	network := &armnetwork.VirtualNetwork{
		Resource: armnetwork.Resource{
			Name: to.StringPtr("network1"),
			ID:   to.StringPtr("/subscriptions/7bfb2c5c-0000-0000-0000-fffa356eb406/resourceGroups/test-dev/providers/Microsoft.Network/virtualNetworks/network1"),
		},
	}

	fakeClient := &mockSubnetsClient{}

	expectedErr := errors.New("unexpected error")

	mockPager := &mockSubnetsListPager{}
	mockPager.On("Err").Return(expectedErr).Times(1)
	mockPager.On("NextPage", mock.Anything).Return(true).Times(1)
	mockPager.On("PageResponse").Return(armnetwork.SubnetsListResponse{}).Times(1)

	fakeClient.On("List", "test-dev", "network1", mock.Anything).Return(mockPager)

	s := &networkRepository{
		subnetsClient: fakeClient,
		cache:         cache.New(0),
	}
	got, err := s.ListAllSubnets(network)

	mockPager.AssertExpectations(t)
	fakeClient.AssertExpectations(t)

	assert.Equal(t, expectedErr, err)
	assert.Nil(t, got)
}

func Test_ListAllSubnets_Error(t *testing.T) {

	network := &armnetwork.VirtualNetwork{
		Resource: armnetwork.Resource{
			Name: to.StringPtr("network1"),
			ID:   to.StringPtr("/subscriptions/7bfb2c5c-0000-0000-0000-fffa356eb406/resourceGroups/test-dev/providers/Microsoft.Network/virtualNetworks/network1"),
		},
	}

	fakeClient := &mockSubnetsClient{}

	expectedErr := errors.New("unexpected error")

	mockPager := &mockSubnetsListPager{}
	mockPager.On("Err").Return(expectedErr).Times(1)
	mockPager.On("NextPage", mock.Anything).Return(false).Times(1)

	fakeClient.On("List", "test-dev", "network1", mock.Anything).Return(mockPager)

	s := &networkRepository{
		subnetsClient: fakeClient,
		cache:         cache.New(0),
	}
	got, err := s.ListAllSubnets(network)

	mockPager.AssertExpectations(t)
	fakeClient.AssertExpectations(t)

	assert.Equal(t, expectedErr, err)
	assert.Nil(t, got)
}

func Test_ListAllSubnets_ErrorOnInvalidNetworkID(t *testing.T) {

	network := &armnetwork.VirtualNetwork{
		Resource: armnetwork.Resource{
			Name: to.StringPtr("network1"),
			ID:   to.StringPtr("foobar"),
		},
	}

	fakeClient := &mockSubnetsClient{}

	expectedErr := errors.New("parsing failed for foobar. Invalid resource Id format")

	s := &networkRepository{
		subnetsClient: fakeClient,
		cache:         cache.New(0),
	}
	got, err := s.ListAllSubnets(network)

	fakeClient.AssertExpectations(t)

	assert.Equal(t, expectedErr.Error(), err.Error())
	assert.Nil(t, got)
}

func Test_ListAllFirewalls_MultiplesResults(t *testing.T) {

	expected := []*armnetwork.AzureFirewall{
		{
			Resource: armnetwork.Resource{
				ID: to.StringPtr("firewall1"),
			},
		},
		{
			Resource: armnetwork.Resource{
				ID: to.StringPtr("firewall2"),
			},
		},
		{
			Resource: armnetwork.Resource{
				ID: to.StringPtr("firewall3"),
			},
		},
		{
			Resource: armnetwork.Resource{
				ID: to.StringPtr("firewall4"),
			},
		},
	}

	fakeClient := &mockFirewallsClient{}

	mockPager := &mockFirewallsListAllPager{}
	mockPager.On("Err").Return(nil).Times(3)
	mockPager.On("NextPage", mock.Anything).Return(true).Times(2)
	mockPager.On("NextPage", mock.Anything).Return(false).Times(1)
	mockPager.On("PageResponse").Return(armnetwork.AzureFirewallsListAllResponse{
		AzureFirewallsListAllResult: armnetwork.AzureFirewallsListAllResult{
			AzureFirewallListResult: armnetwork.AzureFirewallListResult{
				Value: expected[:2],
			},
		},
	}).Times(1)
	mockPager.On("PageResponse").Return(armnetwork.AzureFirewallsListAllResponse{
		AzureFirewallsListAllResult: armnetwork.AzureFirewallsListAllResult{
			AzureFirewallListResult: armnetwork.AzureFirewallListResult{
				Value: expected[2:],
			},
		},
	}).Times(1)

	fakeClient.On("ListAll", mock.Anything).Return(mockPager)

	c := &cache.MockCache{}
	c.On("Get", "ListAllFirewalls").Return(nil).Times(1)
	c.On("Put", "ListAllFirewalls", expected).Return(true).Times(1)
	s := &networkRepository{
		firewallsClient: fakeClient,
		cache:           c,
	}
	got, err := s.ListAllFirewalls()
	if err != nil {
		t.Errorf("ListAllFirewalls() error = %v", err)
		return
	}

	mockPager.AssertExpectations(t)
	fakeClient.AssertExpectations(t)
	c.AssertExpectations(t)

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("ListAllFirewalls() got = %v, want %v", got, expected)
	}
}

func Test_ListAllFirewalls_MultiplesResults_WithCache(t *testing.T) {

	expected := []*armnetwork.AzureFirewall{
		{
			Resource: armnetwork.Resource{
				ID: to.StringPtr("firewall1"),
			},
		},
	}

	fakeClient := &mockFirewallsClient{}

	c := &cache.MockCache{}
	c.On("Get", "ListAllFirewalls").Return(expected).Times(1)
	s := &networkRepository{
		firewallsClient: fakeClient,
		cache:           c,
	}
	got, err := s.ListAllFirewalls()
	if err != nil {
		t.Errorf("ListAllFirewalls() error = %v", err)
		return
	}

	fakeClient.AssertExpectations(t)
	c.AssertExpectations(t)

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("ListAllFirewalls() got = %v, want %v", got, expected)
	}
}

func Test_ListAllFirewalls_Error_OnPageResponse(t *testing.T) {

	fakeClient := &mockFirewallsClient{}

	expectedErr := errors.New("unexpected error")

	mockPager := &mockFirewallsListAllPager{}
	mockPager.On("Err").Return(expectedErr).Times(1)
	mockPager.On("NextPage", mock.Anything).Return(true).Times(1)
	mockPager.On("PageResponse").Return(armnetwork.AzureFirewallsListAllResponse{}).Times(1)

	fakeClient.On("ListAll", mock.Anything).Return(mockPager)

	s := &networkRepository{
		firewallsClient: fakeClient,
		cache:           cache.New(0),
	}
	got, err := s.ListAllFirewalls()

	mockPager.AssertExpectations(t)
	fakeClient.AssertExpectations(t)

	assert.Equal(t, expectedErr, err)
	assert.Nil(t, got)
}

func Test_ListAllFirewalls_Error(t *testing.T) {

	fakeClient := &mockFirewallsClient{}

	expectedErr := errors.New("unexpected error")

	mockPager := &mockFirewallsListAllPager{}
	mockPager.On("Err").Return(expectedErr).Times(1)
	mockPager.On("NextPage", mock.Anything).Return(false).Times(1)

	fakeClient.On("ListAll", mock.Anything).Return(mockPager)

	s := &networkRepository{
		firewallsClient: fakeClient,
		cache:           cache.New(0),
	}
	got, err := s.ListAllFirewalls()

	mockPager.AssertExpectations(t)
	fakeClient.AssertExpectations(t)

	assert.Equal(t, expectedErr, err)
	assert.Nil(t, got)
}