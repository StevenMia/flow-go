// Code generated by mockery v2.21.4. DO NOT EDIT.

package mock

import (
	flow "github.com/onflow/flow-go/model/flow"
	mock "github.com/stretchr/testify/mock"

	protocol "github.com/onflow/flow-go/state/protocol"
)

// Snapshot is an autogenerated mock type for the Snapshot type
type Snapshot struct {
	mock.Mock
}

// Commit provides a mock function with given fields:
func (_m *Snapshot) Commit() (flow.StateCommitment, error) {
	ret := _m.Called()

	var r0 flow.StateCommitment
	var r1 error
	if rf, ok := ret.Get(0).(func() (flow.StateCommitment, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() flow.StateCommitment); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(flow.StateCommitment)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Descendants provides a mock function with given fields:
func (_m *Snapshot) Descendants() ([]flow.Identifier, error) {
	ret := _m.Called()

	var r0 []flow.Identifier
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]flow.Identifier, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []flow.Identifier); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]flow.Identifier)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Epochs provides a mock function with given fields:
func (_m *Snapshot) Epochs() protocol.EpochQuery {
	ret := _m.Called()

	var r0 protocol.EpochQuery
	if rf, ok := ret.Get(0).(func() protocol.EpochQuery); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(protocol.EpochQuery)
		}
	}

	return r0
}

// Head provides a mock function with given fields:
func (_m *Snapshot) Head() (*flow.Header, error) {
	ret := _m.Called()

	var r0 *flow.Header
	var r1 error
	if rf, ok := ret.Get(0).(func() (*flow.Header, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *flow.Header); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.Header)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Identities provides a mock function with given fields: selector
func (_m *Snapshot) Identities(selector flow.IdentityFilter) (flow.IdentityList, error) {
	ret := _m.Called(selector)

	var r0 flow.IdentityList
	var r1 error
	if rf, ok := ret.Get(0).(func(flow.IdentityFilter) (flow.IdentityList, error)); ok {
		return rf(selector)
	}
	if rf, ok := ret.Get(0).(func(flow.IdentityFilter) flow.IdentityList); ok {
		r0 = rf(selector)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(flow.IdentityList)
		}
	}

	if rf, ok := ret.Get(1).(func(flow.IdentityFilter) error); ok {
		r1 = rf(selector)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Identity provides a mock function with given fields: nodeID
func (_m *Snapshot) Identity(nodeID flow.Identifier) (*flow.Identity, error) {
	ret := _m.Called(nodeID)

	var r0 *flow.Identity
	var r1 error
	if rf, ok := ret.Get(0).(func(flow.Identifier) (*flow.Identity, error)); ok {
		return rf(nodeID)
	}
	if rf, ok := ret.Get(0).(func(flow.Identifier) *flow.Identity); ok {
		r0 = rf(nodeID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.Identity)
		}
	}

	if rf, ok := ret.Get(1).(func(flow.Identifier) error); ok {
		r1 = rf(nodeID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Params provides a mock function with given fields:
func (_m *Snapshot) Params() protocol.GlobalParams {
	ret := _m.Called()

	var r0 protocol.GlobalParams
	if rf, ok := ret.Get(0).(func() protocol.GlobalParams); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(protocol.GlobalParams)
		}
	}

	return r0
}

// Phase provides a mock function with given fields:
func (_m *Snapshot) Phase() (flow.EpochPhase, error) {
	ret := _m.Called()

	var r0 flow.EpochPhase
	var r1 error
	if rf, ok := ret.Get(0).(func() (flow.EpochPhase, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() flow.EpochPhase); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(flow.EpochPhase)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QuorumCertificate provides a mock function with given fields:
func (_m *Snapshot) QuorumCertificate() (*flow.QuorumCertificate, error) {
	ret := _m.Called()

	var r0 *flow.QuorumCertificate
	var r1 error
	if rf, ok := ret.Get(0).(func() (*flow.QuorumCertificate, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *flow.QuorumCertificate); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.QuorumCertificate)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RandomSource provides a mock function with given fields:
func (_m *Snapshot) RandomSource() ([]byte, error) {
	ret := _m.Called()

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]byte, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SealedResult provides a mock function with given fields:
func (_m *Snapshot) SealedResult() (*flow.ExecutionResult, *flow.Seal, error) {
	ret := _m.Called()

	var r0 *flow.ExecutionResult
	var r1 *flow.Seal
	var r2 error
	if rf, ok := ret.Get(0).(func() (*flow.ExecutionResult, *flow.Seal, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *flow.ExecutionResult); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.ExecutionResult)
		}
	}

	if rf, ok := ret.Get(1).(func() *flow.Seal); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*flow.Seal)
		}
	}

	if rf, ok := ret.Get(2).(func() error); ok {
		r2 = rf()
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// SealingSegment provides a mock function with given fields:
func (_m *Snapshot) SealingSegment() (*flow.SealingSegment, error) {
	ret := _m.Called()

	var r0 *flow.SealingSegment
	var r1 error
	if rf, ok := ret.Get(0).(func() (*flow.SealingSegment, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *flow.SealingSegment); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.SealingSegment)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VersionBeacon provides a mock function with given fields:
func (_m *Snapshot) VersionBeacon() (*flow.SealedVersionBeacon, error) {
	ret := _m.Called()

	var r0 *flow.SealedVersionBeacon
	var r1 error
	if rf, ok := ret.Get(0).(func() (*flow.SealedVersionBeacon, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *flow.SealedVersionBeacon); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.SealedVersionBeacon)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewSnapshot interface {
	mock.TestingT
	Cleanup(func())
}

// NewSnapshot creates a new instance of Snapshot. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSnapshot(t mockConstructorTestingTNewSnapshot) *Snapshot {
	mock := &Snapshot{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
