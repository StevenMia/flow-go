// Code generated by mockery v2.12.1. DO NOT EDIT.

package mock

import (
	flow "github.com/onflow/flow-go/model/flow"
	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// Collections is an autogenerated mock type for the Collections type
type Collections struct {
	mock.Mock
}

// ByID provides a mock function with given fields: collID
func (_m *Collections) ByID(collID flow.Identifier) (*flow.Collection, error) {
	ret := _m.Called(collID)

	var r0 *flow.Collection
	if rf, ok := ret.Get(0).(func(flow.Identifier) *flow.Collection); ok {
		r0 = rf(collID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.Collection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(flow.Identifier) error); ok {
		r1 = rf(collID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LightByID provides a mock function with given fields: collID
func (_m *Collections) LightByID(collID flow.Identifier) (*flow.LightCollection, error) {
	ret := _m.Called(collID)

	var r0 *flow.LightCollection
	if rf, ok := ret.Get(0).(func(flow.Identifier) *flow.LightCollection); ok {
		r0 = rf(collID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.LightCollection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(flow.Identifier) error); ok {
		r1 = rf(collID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LightByTransactionID provides a mock function with given fields: txID
func (_m *Collections) LightByTransactionID(txID flow.Identifier) (*flow.LightCollection, error) {
	ret := _m.Called(txID)

	var r0 *flow.LightCollection
	if rf, ok := ret.Get(0).(func(flow.Identifier) *flow.LightCollection); ok {
		r0 = rf(txID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.LightCollection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(flow.Identifier) error); ok {
		r1 = rf(txID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Remove provides a mock function with given fields: collID
func (_m *Collections) Remove(collID flow.Identifier) error {
	ret := _m.Called(collID)

	var r0 error
	if rf, ok := ret.Get(0).(func(flow.Identifier) error); ok {
		r0 = rf(collID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Store provides a mock function with given fields: collection
func (_m *Collections) Store(collection *flow.Collection) error {
	ret := _m.Called(collection)

	var r0 error
	if rf, ok := ret.Get(0).(func(*flow.Collection) error); ok {
		r0 = rf(collection)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StoreLight provides a mock function with given fields: collection
func (_m *Collections) StoreLight(collection *flow.LightCollection) error {
	ret := _m.Called(collection)

	var r0 error
	if rf, ok := ret.Get(0).(func(*flow.LightCollection) error); ok {
		r0 = rf(collection)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StoreLightAndIndexByTransaction provides a mock function with given fields: collection
func (_m *Collections) StoreLightAndIndexByTransaction(collection *flow.LightCollection) error {
	ret := _m.Called(collection)

	var r0 error
	if rf, ok := ret.Get(0).(func(*flow.LightCollection) error); ok {
		r0 = rf(collection)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewCollections creates a new instance of Collections. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewCollections(t testing.TB) *Collections {
	mock := &Collections{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
