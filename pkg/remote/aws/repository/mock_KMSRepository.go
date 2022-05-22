// Code generated by mockery v2.12.2. DO NOT EDIT.

package repository

import (
	kms "github.com/aws/aws-sdk-go/service/kms"
	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// MockKMSRepository is an autogenerated mock type for the KMSRepository type
type MockKMSRepository struct {
	mock.Mock
}

// ListAllAliases provides a mock function with given fields:
func (_m *MockKMSRepository) ListAllAliases() ([]*kms.AliasListEntry, error) {
	ret := _m.Called()

	var r0 []*kms.AliasListEntry
	if rf, ok := ret.Get(0).(func() []*kms.AliasListEntry); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*kms.AliasListEntry)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAllKeys provides a mock function with given fields:
func (_m *MockKMSRepository) ListAllKeys() ([]*kms.KeyListEntry, error) {
	ret := _m.Called()

	var r0 []*kms.KeyListEntry
	if rf, ok := ret.Get(0).(func() []*kms.KeyListEntry); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*kms.KeyListEntry)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockKMSRepository creates a new instance of MockKMSRepository. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockKMSRepository(t testing.TB) *MockKMSRepository {
	mock := &MockKMSRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
