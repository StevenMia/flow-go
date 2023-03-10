// Code generated by mockery v2.21.4. DO NOT EDIT.

package mock

import (
	flow "github.com/onflow/flow-go/model/flow"
	mock "github.com/stretchr/testify/mock"
)

// PendingBlockBuffer is an autogenerated mock type for the PendingBlockBuffer type
type PendingBlockBuffer struct {
	mock.Mock
}

// Add provides a mock function with given fields: originID, block
func (_m *PendingBlockBuffer) Add(originID flow.Identifier, block *flow.Block) bool {
	ret := _m.Called(originID, block)

	var r0 bool
	if rf, ok := ret.Get(0).(func(flow.Identifier, *flow.Block) bool); ok {
		r0 = rf(originID, block)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ByID provides a mock function with given fields: blockID
func (_m *PendingBlockBuffer) ByID(blockID flow.Identifier) (flow.Slashable[flow.Block], bool) {
	ret := _m.Called(blockID)

	var r0 flow.Slashable[flow.Block]
	var r1 bool
	if rf, ok := ret.Get(0).(func(flow.Identifier) (flow.Slashable[flow.Block], bool)); ok {
		return rf(blockID)
	}
	if rf, ok := ret.Get(0).(func(flow.Identifier) flow.Slashable[flow.Block]); ok {
		r0 = rf(blockID)
	} else {
		r0 = ret.Get(0).(flow.Slashable[flow.Block])
	}

	if rf, ok := ret.Get(1).(func(flow.Identifier) bool); ok {
		r1 = rf(blockID)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// ByParentID provides a mock function with given fields: parentID
func (_m *PendingBlockBuffer) ByParentID(parentID flow.Identifier) ([]flow.Slashable[flow.Block], bool) {
	ret := _m.Called(parentID)

	var r0 []flow.Slashable[flow.Block]
	var r1 bool
	if rf, ok := ret.Get(0).(func(flow.Identifier) ([]flow.Slashable[flow.Block], bool)); ok {
		return rf(parentID)
	}
	if rf, ok := ret.Get(0).(func(flow.Identifier) []flow.Slashable[flow.Block]); ok {
		r0 = rf(parentID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]flow.Slashable[flow.Block])
		}
	}

	if rf, ok := ret.Get(1).(func(flow.Identifier) bool); ok {
		r1 = rf(parentID)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// DropForParent provides a mock function with given fields: parentID
func (_m *PendingBlockBuffer) DropForParent(parentID flow.Identifier) {
	_m.Called(parentID)
}

// PruneByView provides a mock function with given fields: view
func (_m *PendingBlockBuffer) PruneByView(view uint64) {
	_m.Called(view)
}

// Size provides a mock function with given fields:
func (_m *PendingBlockBuffer) Size() uint {
	ret := _m.Called()

	var r0 uint
	if rf, ok := ret.Get(0).(func() uint); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint)
	}

	return r0
}

type mockConstructorTestingTNewPendingBlockBuffer interface {
	mock.TestingT
	Cleanup(func())
}

// NewPendingBlockBuffer creates a new instance of PendingBlockBuffer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPendingBlockBuffer(t mockConstructorTestingTNewPendingBlockBuffer) *PendingBlockBuffer {
	mock := &PendingBlockBuffer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
