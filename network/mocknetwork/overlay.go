// Code generated by mockery v2.21.4. DO NOT EDIT.

package mocknetwork

import (
	flow "github.com/onflow/flow-go/model/flow"
	mock "github.com/stretchr/testify/mock"

	network "github.com/onflow/flow-go/network"

	peer "github.com/libp2p/go-libp2p/core/peer"
)

// Overlay is an autogenerated mock type for the Overlay type
type Overlay struct {
	mock.Mock
}

// Identities provides a mock function with given fields:
func (_m *Overlay) Identities() flow.IdentityList {
	ret := _m.Called()

	var r0 flow.IdentityList
	if rf, ok := ret.Get(0).(func() flow.IdentityList); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(flow.IdentityList)
		}
	}

	return r0
}

// Identity provides a mock function with given fields: _a0
func (_m *Overlay) Identity(_a0 peer.ID) (*flow.Identity, bool) {
	ret := _m.Called(_a0)

	var r0 *flow.Identity
	var r1 bool
	if rf, ok := ret.Get(0).(func(peer.ID) (*flow.Identity, bool)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(peer.ID) *flow.Identity); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.Identity)
		}
	}

	if rf, ok := ret.Get(1).(func(peer.ID) bool); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// Receive provides a mock function with given fields: _a0
func (_m *Overlay) Receive(_a0 *network.IncomingMessageScope) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*network.IncomingMessageScope) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Topology provides a mock function with given fields:
func (_m *Overlay) Topology() flow.IdentityList {
	ret := _m.Called()

	var r0 flow.IdentityList
	if rf, ok := ret.Get(0).(func() flow.IdentityList); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(flow.IdentityList)
		}
	}

	return r0
}

type mockConstructorTestingTNewOverlay interface {
	mock.TestingT
	Cleanup(func())
}

// NewOverlay creates a new instance of Overlay. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewOverlay(t mockConstructorTestingTNewOverlay) *Overlay {
	mock := &Overlay{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
