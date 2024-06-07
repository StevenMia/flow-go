// Code generated by mockery v2.21.4. DO NOT EDIT.

package mock

import (
	flow "github.com/onflow/flow-go/model/flow"
	mock "github.com/stretchr/testify/mock"
)

// InstanceParams is an autogenerated mock type for the InstanceParams type
type InstanceParams struct {
	mock.Mock
}

// FinalizedRoot provides a mock function with given fields:
func (_m *InstanceParams) FinalizedRoot() *flow.Header {
	ret := _m.Called()

	var r0 *flow.Header
	if rf, ok := ret.Get(0).(func() *flow.Header); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.Header)
		}
	}

	return r0
}

// Seal provides a mock function with given fields:
func (_m *InstanceParams) Seal() *flow.Seal {
	ret := _m.Called()

	var r0 *flow.Seal
	if rf, ok := ret.Get(0).(func() *flow.Seal); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.Seal)
		}
	}

	return r0
}

// SealedRoot provides a mock function with given fields:
func (_m *InstanceParams) SealedRoot() *flow.Header {
	ret := _m.Called()

	var r0 *flow.Header
	if rf, ok := ret.Get(0).(func() *flow.Header); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.Header)
		}
	}

	return r0
}

type mockConstructorTestingTNewInstanceParams interface {
	mock.TestingT
	Cleanup(func())
}

// NewInstanceParams creates a new instance of InstanceParams. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewInstanceParams(t mockConstructorTestingTNewInstanceParams) *InstanceParams {
	mock := &InstanceParams{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
