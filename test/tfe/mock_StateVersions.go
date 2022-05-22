// Code generated by mockery v2.12.2. DO NOT EDIT.

package test_tfe

import (
	context "context"
	testing "testing"

	mock "github.com/stretchr/testify/mock"

	tfe "github.com/hashicorp/go-tfe"
)

// MockStateVersions is an autogenerated mock type for the StateVersions type
type MockStateVersions struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, workspaceID, options
func (_m *MockStateVersions) Create(ctx context.Context, workspaceID string, options tfe.StateVersionCreateOptions) (*tfe.StateVersion, error) {
	ret := _m.Called(ctx, workspaceID, options)

	var r0 *tfe.StateVersion
	if rf, ok := ret.Get(0).(func(context.Context, string, tfe.StateVersionCreateOptions) *tfe.StateVersion); ok {
		r0 = rf(ctx, workspaceID, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*tfe.StateVersion)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, tfe.StateVersionCreateOptions) error); ok {
		r1 = rf(ctx, workspaceID, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Current provides a mock function with given fields: ctx, workspaceID
func (_m *MockStateVersions) Current(ctx context.Context, workspaceID string) (*tfe.StateVersion, error) {
	ret := _m.Called(ctx, workspaceID)

	var r0 *tfe.StateVersion
	if rf, ok := ret.Get(0).(func(context.Context, string) *tfe.StateVersion); ok {
		r0 = rf(ctx, workspaceID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*tfe.StateVersion)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, workspaceID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CurrentWithOptions provides a mock function with given fields: ctx, workspaceID, options
func (_m *MockStateVersions) CurrentWithOptions(ctx context.Context, workspaceID string, options *tfe.StateVersionCurrentOptions) (*tfe.StateVersion, error) {
	ret := _m.Called(ctx, workspaceID, options)

	var r0 *tfe.StateVersion
	if rf, ok := ret.Get(0).(func(context.Context, string, *tfe.StateVersionCurrentOptions) *tfe.StateVersion); ok {
		r0 = rf(ctx, workspaceID, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*tfe.StateVersion)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, *tfe.StateVersionCurrentOptions) error); ok {
		r1 = rf(ctx, workspaceID, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Download provides a mock function with given fields: ctx, url
func (_m *MockStateVersions) Download(ctx context.Context, url string) ([]byte, error) {
	ret := _m.Called(ctx, url)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(context.Context, string) []byte); ok {
		r0 = rf(ctx, url)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, url)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, options
func (_m *MockStateVersions) List(ctx context.Context, options tfe.StateVersionListOptions) (*tfe.StateVersionList, error) {
	ret := _m.Called(ctx, options)

	var r0 *tfe.StateVersionList
	if rf, ok := ret.Get(0).(func(context.Context, tfe.StateVersionListOptions) *tfe.StateVersionList); ok {
		r0 = rf(ctx, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*tfe.StateVersionList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, tfe.StateVersionListOptions) error); ok {
		r1 = rf(ctx, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Outputs provides a mock function with given fields: ctx, svID, options
func (_m *MockStateVersions) Outputs(ctx context.Context, svID string, options tfe.StateVersionOutputsListOptions) ([]*tfe.StateVersionOutput, error) {
	ret := _m.Called(ctx, svID, options)

	var r0 []*tfe.StateVersionOutput
	if rf, ok := ret.Get(0).(func(context.Context, string, tfe.StateVersionOutputsListOptions) []*tfe.StateVersionOutput); ok {
		r0 = rf(ctx, svID, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*tfe.StateVersionOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, tfe.StateVersionOutputsListOptions) error); ok {
		r1 = rf(ctx, svID, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Read provides a mock function with given fields: ctx, svID
func (_m *MockStateVersions) Read(ctx context.Context, svID string) (*tfe.StateVersion, error) {
	ret := _m.Called(ctx, svID)

	var r0 *tfe.StateVersion
	if rf, ok := ret.Get(0).(func(context.Context, string) *tfe.StateVersion); ok {
		r0 = rf(ctx, svID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*tfe.StateVersion)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, svID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadWithOptions provides a mock function with given fields: ctx, svID, options
func (_m *MockStateVersions) ReadWithOptions(ctx context.Context, svID string, options *tfe.StateVersionReadOptions) (*tfe.StateVersion, error) {
	ret := _m.Called(ctx, svID, options)

	var r0 *tfe.StateVersion
	if rf, ok := ret.Get(0).(func(context.Context, string, *tfe.StateVersionReadOptions) *tfe.StateVersion); ok {
		r0 = rf(ctx, svID, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*tfe.StateVersion)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, *tfe.StateVersionReadOptions) error); ok {
		r1 = rf(ctx, svID, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockStateVersions creates a new instance of MockStateVersions. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockStateVersions(t testing.TB) *MockStateVersions {
	mock := &MockStateVersions{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
