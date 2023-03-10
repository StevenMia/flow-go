// Code generated by mockery v2.21.4. DO NOT EDIT.

package mock

import (
	cadence "github.com/onflow/cadence"

	flow "github.com/onflow/flow-go/model/flow"

	mock "github.com/stretchr/testify/mock"
)

// ContractUpdaterStubs is an autogenerated mock type for the ContractUpdaterStubs type
type ContractUpdaterStubs struct {
	mock.Mock
}

// GetAuthorizedAccounts provides a mock function with given fields: path
func (_m *ContractUpdaterStubs) GetAuthorizedAccounts(path cadence.Path) []flow.Address {
	ret := _m.Called(path)

	var r0 []flow.Address
	if rf, ok := ret.Get(0).(func(cadence.Path) []flow.Address); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]flow.Address)
		}
	}

	return r0
}

// RestrictedDeploymentEnabled provides a mock function with given fields:
func (_m *ContractUpdaterStubs) RestrictedDeploymentEnabled() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// RestrictedRemovalEnabled provides a mock function with given fields:
func (_m *ContractUpdaterStubs) RestrictedRemovalEnabled() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// UseContractAuditVoucher provides a mock function with given fields: address, code
func (_m *ContractUpdaterStubs) UseContractAuditVoucher(address flow.Address, code []byte) (bool, error) {
	ret := _m.Called(address, code)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(flow.Address, []byte) (bool, error)); ok {
		return rf(address, code)
	}
	if rf, ok := ret.Get(0).(func(flow.Address, []byte) bool); ok {
		r0 = rf(address, code)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(flow.Address, []byte) error); ok {
		r1 = rf(address, code)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewContractUpdaterStubs interface {
	mock.TestingT
	Cleanup(func())
}

// NewContractUpdaterStubs creates a new instance of ContractUpdaterStubs. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewContractUpdaterStubs(t mockConstructorTestingTNewContractUpdaterStubs) *ContractUpdaterStubs {
	mock := &ContractUpdaterStubs{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
