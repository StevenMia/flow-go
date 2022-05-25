// Code generated by mockery v2.12.1. DO NOT EDIT.

package mock

import (
	mock "github.com/stretchr/testify/mock"

	testing "testing"

	time "time"
)

// CleanerMetrics is an autogenerated mock type for the CleanerMetrics type
type CleanerMetrics struct {
	mock.Mock
}

// RanGC provides a mock function with given fields: took
func (_m *CleanerMetrics) RanGC(took time.Duration) {
	_m.Called(took)
}

// NewCleanerMetrics creates a new instance of CleanerMetrics. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewCleanerMetrics(t testing.TB) *CleanerMetrics {
	mock := &CleanerMetrics{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
