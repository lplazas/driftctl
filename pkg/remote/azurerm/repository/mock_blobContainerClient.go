// Code generated by mockery v2.12.2. DO NOT EDIT.

package repository

import (
	armstorage "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// mockBlobContainerClient is an autogenerated mock type for the blobContainerClient type
type mockBlobContainerClient struct {
	mock.Mock
}

// List provides a mock function with given fields: resourceGroupName, accountName, options
func (_m *mockBlobContainerClient) List(resourceGroupName string, accountName string, options *armstorage.BlobContainersListOptions) blobContainerListPager {
	ret := _m.Called(resourceGroupName, accountName, options)

	var r0 blobContainerListPager
	if rf, ok := ret.Get(0).(func(string, string, *armstorage.BlobContainersListOptions) blobContainerListPager); ok {
		r0 = rf(resourceGroupName, accountName, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(blobContainerListPager)
		}
	}

	return r0
}

// newMockBlobContainerClient creates a new instance of mockBlobContainerClient. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func newMockBlobContainerClient(t testing.TB) *mockBlobContainerClient {
	mock := &mockBlobContainerClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
