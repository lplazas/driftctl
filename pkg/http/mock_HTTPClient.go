// Code generated by mockery v2.12.2. DO NOT EDIT.

package http

import (
	nethttp "net/http"

	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// MockHTTPClient is an autogenerated mock type for the HTTPClient type
type MockHTTPClient struct {
	mock.Mock
}

// Do provides a mock function with given fields: req
func (_m *MockHTTPClient) Do(req *nethttp.Request) (*nethttp.Response, error) {
	ret := _m.Called(req)

	var r0 *nethttp.Response
	if rf, ok := ret.Get(0).(func(*nethttp.Request) *nethttp.Response); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*nethttp.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*nethttp.Request) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockHTTPClient creates a new instance of MockHTTPClient. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockHTTPClient(t testing.TB) *MockHTTPClient {
	mock := &MockHTTPClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
