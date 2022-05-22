// Code generated by mockery v2.12.2. DO NOT EDIT.

package repository

import (
	armstorage "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// mockStorageAccountClient is an autogenerated mock type for the storageAccountClient type
type mockStorageAccountClient struct {
	mock.Mock
}

// List provides a mock function with given fields: options
func (_m *mockStorageAccountClient) List(options *armstorage.StorageAccountsListOptions) storageAccountListPager {
	ret := _m.Called(options)

	var r0 storageAccountListPager
	if rf, ok := ret.Get(0).(func(*armstorage.StorageAccountsListOptions) storageAccountListPager); ok {
		r0 = rf(options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(storageAccountListPager)
		}
	}

	return r0
}

// newMockStorageAccountClient creates a new instance of mockStorageAccountClient. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func newMockStorageAccountClient(t testing.TB) *mockStorageAccountClient {
	mock := &mockStorageAccountClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
