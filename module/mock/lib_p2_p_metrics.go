// Code generated by mockery v2.21.4. DO NOT EDIT.

package mock

import (
	channels "github.com/onflow/flow-go/network/channels"
	mock "github.com/stretchr/testify/mock"

	network "github.com/libp2p/go-libp2p/core/network"

	p2pmsg "github.com/onflow/flow-go/network/p2p/message"

	peer "github.com/libp2p/go-libp2p/core/peer"

	protocol "github.com/libp2p/go-libp2p/core/protocol"

	time "time"
)

// LibP2PMetrics is an autogenerated mock type for the LibP2PMetrics type
type LibP2PMetrics struct {
	mock.Mock
}

// AllowConn provides a mock function with given fields: dir, usefd
func (_m *LibP2PMetrics) AllowConn(dir network.Direction, usefd bool) {
	_m.Called(dir, usefd)
}

// AllowMemory provides a mock function with given fields: size
func (_m *LibP2PMetrics) AllowMemory(size int) {
	_m.Called(size)
}

// AllowPeer provides a mock function with given fields: p
func (_m *LibP2PMetrics) AllowPeer(p peer.ID) {
	_m.Called(p)
}

// AllowProtocol provides a mock function with given fields: proto
func (_m *LibP2PMetrics) AllowProtocol(proto protocol.ID) {
	_m.Called(proto)
}

// AllowService provides a mock function with given fields: svc
func (_m *LibP2PMetrics) AllowService(svc string) {
	_m.Called(svc)
}

// AllowStream provides a mock function with given fields: p, dir
func (_m *LibP2PMetrics) AllowStream(p peer.ID, dir network.Direction) {
	_m.Called(p, dir)
}

// AsyncProcessingFinished provides a mock function with given fields: duration
func (_m *LibP2PMetrics) AsyncProcessingFinished(duration time.Duration) {
	_m.Called(duration)
}

// AsyncProcessingStarted provides a mock function with given fields:
func (_m *LibP2PMetrics) AsyncProcessingStarted() {
	_m.Called()
}

// BlockConn provides a mock function with given fields: dir, usefd
func (_m *LibP2PMetrics) BlockConn(dir network.Direction, usefd bool) {
	_m.Called(dir, usefd)
}

// BlockMemory provides a mock function with given fields: size
func (_m *LibP2PMetrics) BlockMemory(size int) {
	_m.Called(size)
}

// BlockPeer provides a mock function with given fields: p
func (_m *LibP2PMetrics) BlockPeer(p peer.ID) {
	_m.Called(p)
}

// BlockProtocol provides a mock function with given fields: proto
func (_m *LibP2PMetrics) BlockProtocol(proto protocol.ID) {
	_m.Called(proto)
}

// BlockProtocolPeer provides a mock function with given fields: proto, p
func (_m *LibP2PMetrics) BlockProtocolPeer(proto protocol.ID, p peer.ID) {
	_m.Called(proto, p)
}

// BlockService provides a mock function with given fields: svc
func (_m *LibP2PMetrics) BlockService(svc string) {
	_m.Called(svc)
}

// BlockServicePeer provides a mock function with given fields: svc, p
func (_m *LibP2PMetrics) BlockServicePeer(svc string, p peer.ID) {
	_m.Called(svc, p)
}

// BlockStream provides a mock function with given fields: p, dir
func (_m *LibP2PMetrics) BlockStream(p peer.ID, dir network.Direction) {
	_m.Called(p, dir)
}

// DNSLookupDuration provides a mock function with given fields: duration
func (_m *LibP2PMetrics) DNSLookupDuration(duration time.Duration) {
	_m.Called(duration)
}

// InboundConnections provides a mock function with given fields: connectionCount
func (_m *LibP2PMetrics) InboundConnections(connectionCount uint) {
	_m.Called(connectionCount)
}

// InvalidControlMessageNotificationError provides a mock function with given fields: msgType, count
func (_m *LibP2PMetrics) InvalidControlMessageNotificationError(msgType p2pmsg.ControlMessageType, count float64) {
	_m.Called(msgType, count)
}

// OnAppSpecificScoreUpdated provides a mock function with given fields: _a0
func (_m *LibP2PMetrics) OnAppSpecificScoreUpdated(_a0 float64) {
	_m.Called(_a0)
}

// OnBehaviourPenaltyUpdated provides a mock function with given fields: _a0
func (_m *LibP2PMetrics) OnBehaviourPenaltyUpdated(_a0 float64) {
	_m.Called(_a0)
}

// OnControlMessagesTruncated provides a mock function with given fields: messageType, diff
func (_m *LibP2PMetrics) OnControlMessagesTruncated(messageType p2pmsg.ControlMessageType, diff int) {
	_m.Called(messageType, diff)
}

// OnDNSCacheHit provides a mock function with given fields:
func (_m *LibP2PMetrics) OnDNSCacheHit() {
	_m.Called()
}

// OnDNSCacheInvalidated provides a mock function with given fields:
func (_m *LibP2PMetrics) OnDNSCacheInvalidated() {
	_m.Called()
}

// OnDNSCacheMiss provides a mock function with given fields:
func (_m *LibP2PMetrics) OnDNSCacheMiss() {
	_m.Called()
}

// OnDNSLookupRequestDropped provides a mock function with given fields:
func (_m *LibP2PMetrics) OnDNSLookupRequestDropped() {
	_m.Called()
}

// OnDialRetryBudgetResetToDefault provides a mock function with given fields:
func (_m *LibP2PMetrics) OnDialRetryBudgetResetToDefault() {
	_m.Called()
}

// OnDialRetryBudgetUpdated provides a mock function with given fields: budget
func (_m *LibP2PMetrics) OnDialRetryBudgetUpdated(budget uint64) {
	_m.Called(budget)
}

// OnEstablishStreamFailure provides a mock function with given fields: duration, attempts
func (_m *LibP2PMetrics) OnEstablishStreamFailure(duration time.Duration, attempts int) {
	_m.Called(duration, attempts)
}

// OnFirstMessageDeliveredUpdated provides a mock function with given fields: _a0, _a1
func (_m *LibP2PMetrics) OnFirstMessageDeliveredUpdated(_a0 channels.Topic, _a1 float64) {
	_m.Called(_a0, _a1)
}

// OnIHaveControlMessageIdsTruncated provides a mock function with given fields: diff
func (_m *LibP2PMetrics) OnIHaveControlMessageIdsTruncated(diff int) {
	_m.Called(diff)
}

// OnIHaveMessageIDsReceived provides a mock function with given fields: channel, msgIdCount
func (_m *LibP2PMetrics) OnIHaveMessageIDsReceived(channel string, msgIdCount int) {
	_m.Called(channel, msgIdCount)
}

// OnIPColocationFactorUpdated provides a mock function with given fields: _a0
func (_m *LibP2PMetrics) OnIPColocationFactorUpdated(_a0 float64) {
	_m.Called(_a0)
}

// OnIWantControlMessageIdsTruncated provides a mock function with given fields: diff
func (_m *LibP2PMetrics) OnIWantControlMessageIdsTruncated(diff int) {
	_m.Called(diff)
}

// OnIWantMessageIDsReceived provides a mock function with given fields: msgIdCount
func (_m *LibP2PMetrics) OnIWantMessageIDsReceived(msgIdCount int) {
	_m.Called(msgIdCount)
}

// OnIncomingRpcReceived provides a mock function with given fields: iHaveCount, iWantCount, graftCount, pruneCount, msgCount
func (_m *LibP2PMetrics) OnIncomingRpcReceived(iHaveCount int, iWantCount int, graftCount int, pruneCount int, msgCount int) {
	_m.Called(iHaveCount, iWantCount, graftCount, pruneCount, msgCount)
}

// OnInvalidMessageDeliveredUpdated provides a mock function with given fields: _a0, _a1
func (_m *LibP2PMetrics) OnInvalidMessageDeliveredUpdated(_a0 channels.Topic, _a1 float64) {
	_m.Called(_a0, _a1)
}

// OnLocalMeshSizeUpdated provides a mock function with given fields: topic, size
func (_m *LibP2PMetrics) OnLocalMeshSizeUpdated(topic string, size int) {
	_m.Called(topic, size)
}

// OnLocalPeerJoinedTopic provides a mock function with given fields:
func (_m *LibP2PMetrics) OnLocalPeerJoinedTopic() {
	_m.Called()
}

// OnLocalPeerLeftTopic provides a mock function with given fields:
func (_m *LibP2PMetrics) OnLocalPeerLeftTopic() {
	_m.Called()
}

// OnMeshMessageDeliveredUpdated provides a mock function with given fields: _a0, _a1
func (_m *LibP2PMetrics) OnMeshMessageDeliveredUpdated(_a0 channels.Topic, _a1 float64) {
	_m.Called(_a0, _a1)
}

// OnMessageDeliveredToAllSubscribers provides a mock function with given fields: size
func (_m *LibP2PMetrics) OnMessageDeliveredToAllSubscribers(size int) {
	_m.Called(size)
}

// OnMessageDuplicate provides a mock function with given fields: size
func (_m *LibP2PMetrics) OnMessageDuplicate(size int) {
	_m.Called(size)
}

// OnMessageEnteredValidation provides a mock function with given fields: size
func (_m *LibP2PMetrics) OnMessageEnteredValidation(size int) {
	_m.Called(size)
}

// OnMessageRejected provides a mock function with given fields: size, reason
func (_m *LibP2PMetrics) OnMessageRejected(size int, reason string) {
	_m.Called(size, reason)
}

// OnOutboundRpcDropped provides a mock function with given fields:
func (_m *LibP2PMetrics) OnOutboundRpcDropped() {
	_m.Called()
}

// OnOverallPeerScoreUpdated provides a mock function with given fields: _a0
func (_m *LibP2PMetrics) OnOverallPeerScoreUpdated(_a0 float64) {
	_m.Called(_a0)
}

// OnPeerAddedToProtocol provides a mock function with given fields: _a0
func (_m *LibP2PMetrics) OnPeerAddedToProtocol(_a0 string) {
	_m.Called(_a0)
}

// OnPeerDialFailure provides a mock function with given fields: duration, attempts
func (_m *LibP2PMetrics) OnPeerDialFailure(duration time.Duration, attempts int) {
	_m.Called(duration, attempts)
}

// OnPeerDialed provides a mock function with given fields: duration, attempts
func (_m *LibP2PMetrics) OnPeerDialed(duration time.Duration, attempts int) {
	_m.Called(duration, attempts)
}

// OnPeerGraftTopic provides a mock function with given fields: topic
func (_m *LibP2PMetrics) OnPeerGraftTopic(topic string) {
	_m.Called(topic)
}

// OnPeerPruneTopic provides a mock function with given fields: topic
func (_m *LibP2PMetrics) OnPeerPruneTopic(topic string) {
	_m.Called(topic)
}

// OnPeerRemovedFromProtocol provides a mock function with given fields:
func (_m *LibP2PMetrics) OnPeerRemovedFromProtocol() {
	_m.Called()
}

// OnPeerThrottled provides a mock function with given fields:
func (_m *LibP2PMetrics) OnPeerThrottled() {
	_m.Called()
}

// OnRpcReceived provides a mock function with given fields: msgCount, iHaveCount, iWantCount, graftCount, pruneCount
func (_m *LibP2PMetrics) OnRpcReceived(msgCount int, iHaveCount int, iWantCount int, graftCount int, pruneCount int) {
	_m.Called(msgCount, iHaveCount, iWantCount, graftCount, pruneCount)
}

// OnRpcSent provides a mock function with given fields: msgCount, iHaveCount, iWantCount, graftCount, pruneCount
func (_m *LibP2PMetrics) OnRpcSent(msgCount int, iHaveCount int, iWantCount int, graftCount int, pruneCount int) {
	_m.Called(msgCount, iHaveCount, iWantCount, graftCount, pruneCount)
}

// OnStreamCreated provides a mock function with given fields: duration, attempts
func (_m *LibP2PMetrics) OnStreamCreated(duration time.Duration, attempts int) {
	_m.Called(duration, attempts)
}

// OnStreamCreationFailure provides a mock function with given fields: duration, attempts
func (_m *LibP2PMetrics) OnStreamCreationFailure(duration time.Duration, attempts int) {
	_m.Called(duration, attempts)
}

// OnStreamCreationRetryBudgetResetToDefault provides a mock function with given fields:
func (_m *LibP2PMetrics) OnStreamCreationRetryBudgetResetToDefault() {
	_m.Called()
}

// OnStreamCreationRetryBudgetUpdated provides a mock function with given fields: budget
func (_m *LibP2PMetrics) OnStreamCreationRetryBudgetUpdated(budget uint64) {
	_m.Called(budget)
}

// OnStreamEstablished provides a mock function with given fields: duration, attempts
func (_m *LibP2PMetrics) OnStreamEstablished(duration time.Duration, attempts int) {
	_m.Called(duration, attempts)
}

// OnTimeInMeshUpdated provides a mock function with given fields: _a0, _a1
func (_m *LibP2PMetrics) OnTimeInMeshUpdated(_a0 channels.Topic, _a1 time.Duration) {
	_m.Called(_a0, _a1)
}

// OnUndeliveredMessage provides a mock function with given fields:
func (_m *LibP2PMetrics) OnUndeliveredMessage() {
	_m.Called()
}

// OutboundConnections provides a mock function with given fields: connectionCount
func (_m *LibP2PMetrics) OutboundConnections(connectionCount uint) {
	_m.Called(connectionCount)
}

// RoutingTablePeerAdded provides a mock function with given fields:
func (_m *LibP2PMetrics) RoutingTablePeerAdded() {
	_m.Called()
}

// RoutingTablePeerRemoved provides a mock function with given fields:
func (_m *LibP2PMetrics) RoutingTablePeerRemoved() {
	_m.Called()
}

// SetWarningStateCount provides a mock function with given fields: _a0
func (_m *LibP2PMetrics) SetWarningStateCount(_a0 uint) {
	_m.Called(_a0)
}

type mockConstructorTestingTNewLibP2PMetrics interface {
	mock.TestingT
	Cleanup(func())
}

// NewLibP2PMetrics creates a new instance of LibP2PMetrics. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewLibP2PMetrics(t mockConstructorTestingTNewLibP2PMetrics) *LibP2PMetrics {
	mock := &LibP2PMetrics{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
