// Code generated by mockery v2.12.2. DO NOT EDIT.

package resource

import (
	testing "testing"

	mock "github.com/stretchr/testify/mock"
)

// MockSchemaRepositoryInterface is an autogenerated mock type for the SchemaRepositoryInterface type
type MockSchemaRepositoryInterface struct {
	mock.Mock
}

// GetSchema provides a mock function with given fields: resourceType
func (_m *MockSchemaRepositoryInterface) GetSchema(resourceType string) (*Schema, bool) {
	ret := _m.Called(resourceType)

	var r0 *Schema
	if rf, ok := ret.Get(0).(func(string) *Schema); ok {
		r0 = rf(resourceType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Schema)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(resourceType)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// SetDiscriminantFunc provides a mock function with given fields: _a0, _a1
func (_m *MockSchemaRepositoryInterface) SetDiscriminantFunc(_a0 string, _a1 func(*Resource, *Resource) bool) {
	_m.Called(_a0, _a1)
}

// SetFlags provides a mock function with given fields: typ, flags
func (_m *MockSchemaRepositoryInterface) SetFlags(typ string, flags ...Flags) {
	_va := make([]interface{}, len(flags))
	for _i := range flags {
		_va[_i] = flags[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, typ)
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// SetHumanReadableAttributesFunc provides a mock function with given fields: typ, humanReadableAttributesFunc
func (_m *MockSchemaRepositoryInterface) SetHumanReadableAttributesFunc(typ string, humanReadableAttributesFunc func(*Resource) map[string]string) {
	_m.Called(typ, humanReadableAttributesFunc)
}

// SetNormalizeFunc provides a mock function with given fields: typ, normalizeFunc
func (_m *MockSchemaRepositoryInterface) SetNormalizeFunc(typ string, normalizeFunc func(*Resource)) {
	_m.Called(typ, normalizeFunc)
}

// SetResolveReadAttributesFunc provides a mock function with given fields: typ, resolveReadAttributesFunc
func (_m *MockSchemaRepositoryInterface) SetResolveReadAttributesFunc(typ string, resolveReadAttributesFunc func(*Resource) map[string]string) {
	_m.Called(typ, resolveReadAttributesFunc)
}

// UpdateSchema provides a mock function with given fields: typ, schemasMutators
func (_m *MockSchemaRepositoryInterface) UpdateSchema(typ string, schemasMutators map[string]func(*AttributeSchema)) {
	_m.Called(typ, schemasMutators)
}

// NewMockSchemaRepositoryInterface creates a new instance of MockSchemaRepositoryInterface. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockSchemaRepositoryInterface(t testing.TB) *MockSchemaRepositoryInterface {
	mock := &MockSchemaRepositoryInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
