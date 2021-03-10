// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package repository

import (
	ec2 "github.com/aws/aws-sdk-go/service/ec2"
	mock "github.com/stretchr/testify/mock"
)

// MockEC2Repository is an autogenerated mock type for the EC2Repository type
type MockEC2Repository struct {
	mock.Mock
}

// ListAllAddresses provides a mock function with given fields:
func (_m *MockEC2Repository) ListAllAddresses() ([]*ec2.Address, error) {
	ret := _m.Called()

	var r0 []*ec2.Address
	if rf, ok := ret.Get(0).(func() []*ec2.Address); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ec2.Address)
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

// ListAllAddressesAssociation provides a mock function with given fields:
func (_m *MockEC2Repository) ListAllAddressesAssociation() ([]string, error) {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
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

// ListAllImages provides a mock function with given fields:
func (_m *MockEC2Repository) ListAllImages() ([]*ec2.Image, error) {
	ret := _m.Called()

	var r0 []*ec2.Image
	if rf, ok := ret.Get(0).(func() []*ec2.Image); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ec2.Image)
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

// ListAllInstances provides a mock function with given fields:
func (_m *MockEC2Repository) ListAllInstances() ([]*ec2.Instance, error) {
	ret := _m.Called()

	var r0 []*ec2.Instance
	if rf, ok := ret.Get(0).(func() []*ec2.Instance); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ec2.Instance)
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

// ListAllKeyPairs provides a mock function with given fields:
func (_m *MockEC2Repository) ListAllKeyPairs() ([]*ec2.KeyPairInfo, error) {
	ret := _m.Called()

	var r0 []*ec2.KeyPairInfo
	if rf, ok := ret.Get(0).(func() []*ec2.KeyPairInfo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ec2.KeyPairInfo)
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

// ListAllSnapshots provides a mock function with given fields:
func (_m *MockEC2Repository) ListAllSnapshots() ([]*ec2.Snapshot, error) {
	ret := _m.Called()

	var r0 []*ec2.Snapshot
	if rf, ok := ret.Get(0).(func() []*ec2.Snapshot); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ec2.Snapshot)
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

// ListAllVolumes provides a mock function with given fields:
func (_m *MockEC2Repository) ListAllVolumes() ([]*ec2.Volume, error) {
	ret := _m.Called()

	var r0 []*ec2.Volume
	if rf, ok := ret.Get(0).(func() []*ec2.Volume); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ec2.Volume)
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