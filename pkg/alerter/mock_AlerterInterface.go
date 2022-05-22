// Code generated by mockery v2.12.2. DO NOT EDIT.

package alerter

import (
	testing "testing"

	mock "github.com/stretchr/testify/mock"
)

// MockAlerterInterface is an autogenerated mock type for the AlerterInterface type
type MockAlerterInterface struct {
	mock.Mock
}

// SendAlert provides a mock function with given fields: key, alert
func (_m *MockAlerterInterface) SendAlert(key string, alert Alert) {
	_m.Called(key, alert)
}

// NewMockAlerterInterface creates a new instance of MockAlerterInterface. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockAlerterInterface(t testing.TB) *MockAlerterInterface {
	mock := &MockAlerterInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
