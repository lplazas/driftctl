// Code generated by mockery v2.12.2. DO NOT EDIT.

package repository

import (
	cloudfront "github.com/aws/aws-sdk-go/service/cloudfront"
	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// MockCloudfrontRepository is an autogenerated mock type for the CloudfrontRepository type
type MockCloudfrontRepository struct {
	mock.Mock
}

// ListAllDistributions provides a mock function with given fields:
func (_m *MockCloudfrontRepository) ListAllDistributions() ([]*cloudfront.DistributionSummary, error) {
	ret := _m.Called()

	var r0 []*cloudfront.DistributionSummary
	if rf, ok := ret.Get(0).(func() []*cloudfront.DistributionSummary); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*cloudfront.DistributionSummary)
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

// NewMockCloudfrontRepository creates a new instance of MockCloudfrontRepository. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockCloudfrontRepository(t testing.TB) *MockCloudfrontRepository {
	mock := &MockCloudfrontRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
