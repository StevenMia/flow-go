// Code generated by mockery v2.21.4. DO NOT EDIT.

package mockp2p

import (
	irrecoverable "github.com/onflow/flow-go/module/irrecoverable"
	mock "github.com/stretchr/testify/mock"

	p2p "github.com/onflow/flow-go/network/p2p"
)

// GossipSubInspectorNotificationDistributor is an autogenerated mock type for the GossipSubInspectorNotificationDistributor type
type GossipSubInspectorNotificationDistributor struct {
	mock.Mock
}

// AddConsumer provides a mock function with given fields: _a0
func (_m *GossipSubInspectorNotificationDistributor) AddConsumer(_a0 p2p.GossipSubInvalidControlMessageNotificationConsumer) {
	_m.Called(_a0)
}

// DistributeInvalidControlMessageNotification provides a mock function with given fields: notification
func (_m *GossipSubInspectorNotificationDistributor) DistributeInvalidControlMessageNotification(notification *p2p.InvalidControlMessageNotification) error {
	ret := _m.Called(notification)

	var r0 error
	if rf, ok := ret.Get(0).(func(*p2p.InvalidControlMessageNotification) error); ok {
		r0 = rf(notification)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Done provides a mock function with given fields:
func (_m *GossipSubInspectorNotificationDistributor) Done() <-chan struct{} {
	ret := _m.Called()

	var r0 <-chan struct{}
	if rf, ok := ret.Get(0).(func() <-chan struct{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan struct{})
		}
	}

	return r0
}

// Ready provides a mock function with given fields:
func (_m *GossipSubInspectorNotificationDistributor) Ready() <-chan struct{} {
	ret := _m.Called()

	var r0 <-chan struct{}
	if rf, ok := ret.Get(0).(func() <-chan struct{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan struct{})
		}
	}

	return r0
}

// Start provides a mock function with given fields: _a0
func (_m *GossipSubInspectorNotificationDistributor) Start(_a0 irrecoverable.SignalerContext) {
	_m.Called(_a0)
}

type mockConstructorTestingTNewGossipSubInspectorNotificationDistributor interface {
	mock.TestingT
	Cleanup(func())
}

// NewGossipSubInspectorNotificationDistributor creates a new instance of GossipSubInspectorNotificationDistributor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGossipSubInspectorNotificationDistributor(t mockConstructorTestingTNewGossipSubInspectorNotificationDistributor) *GossipSubInspectorNotificationDistributor {
	mock := &GossipSubInspectorNotificationDistributor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
