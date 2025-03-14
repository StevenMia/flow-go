package committees

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/onflow/flow-go/consensus/hotstuff/model"
	"github.com/onflow/flow-go/model/flow"
	"github.com/onflow/flow-go/model/flow/mapfunc"
	"github.com/onflow/flow-go/module/irrecoverable"
	"github.com/onflow/flow-go/state/protocol"
	protocolmock "github.com/onflow/flow-go/state/protocol/mock"
	"github.com/onflow/flow-go/utils/unittest"
	"github.com/onflow/flow-go/utils/unittest/mocks"
)

func TestConsensusCommittee(t *testing.T) {
	suite.Run(t, new(ConsensusSuite))
}

type ConsensusSuite struct {
	suite.Suite

	// mocks
	state    *protocolmock.State
	snapshot *protocolmock.Snapshot
	epochs   *mocks.EpochQuery

	// backend for mocked functions
	phase               flow.EpochPhase
	currentEpochCounter uint64
	myID                flow.Identifier

	committee *Consensus
	cancel    context.CancelFunc
}

// SetupTest instantiates mocks for a test case.
// By default, we start in the Staking phase with no epochs mocked; test cases must add their own epoch mocks.
func (suite *ConsensusSuite) SetupTest() {
	suite.phase = flow.EpochPhaseStaking
	suite.currentEpochCounter = 1
	suite.myID = unittest.IdentifierFixture()

	suite.state = new(protocolmock.State)
	suite.snapshot = new(protocolmock.Snapshot)
	suite.epochs = mocks.NewEpochQuery(suite.T(), suite.currentEpochCounter)

	suite.state.On("Final").Return(suite.snapshot)
	suite.snapshot.On("EpochPhase").Return(
		func() flow.EpochPhase { return suite.phase },
		func() error { return nil },
	)
	suite.snapshot.On("Epochs").Return(suite.epochs)
}

func (suite *ConsensusSuite) TearDownTest() {
	if suite.cancel != nil {
		suite.cancel()
	}
	unittest.AssertClosesBefore(suite.T(), suite.committee.Done(), time.Second)
}

// CreateAndStartCommittee instantiates and starts the committee.
// Should be called only once per test, after initial epoch mocks are created.
// It spawns a goroutine to detect fatal errors from the committee's error channel.
func (suite *ConsensusSuite) CreateAndStartCommittee() {
	committee, err := NewConsensusCommittee(suite.state, suite.myID)
	require.NoError(suite.T(), err)
	ctx, cancel, errCh := irrecoverable.WithSignallerAndCancel(context.Background())
	committee.Start(ctx)
	go unittest.FailOnIrrecoverableError(suite.T(), ctx.Done(), errCh)

	suite.committee = committee
	suite.cancel = cancel
}

// CommitEpoch adds the epoch to the protocol state and mimics the protocol state
// behaviour when committing an epoch, by sending the protocol event to the committee.
func (suite *ConsensusSuite) CommitEpoch(epoch protocol.CommittedEpoch) {
	firstBlockOfCommittedPhase := unittest.BlockHeaderFixture()
	suite.state.On("AtHeight", firstBlockOfCommittedPhase.Height).Return(suite.snapshot)
	suite.epochs.AddCommitted(epoch)
	suite.committee.EpochCommittedPhaseStarted(1, firstBlockOfCommittedPhase)

	// get the first view, to test when the epoch has been processed
	firstView := epoch.FirstView()

	// wait for the protocol event to be processed (async)
	assert.Eventually(suite.T(), func() bool {
		_, err := suite.committee.IdentitiesByEpoch(firstView)
		return err == nil
	}, time.Second, time.Millisecond)
}

// AssertKnownViews asserts that no errors is returned when querying identities by epoch for each of the input views.
func (suite *ConsensusSuite) AssertKnownViews(views ...uint64) {
	for _, view := range views {
		_, err := suite.committee.IdentitiesByEpoch(view)
		suite.Assert().NoError(err)
	}
}

// AssertUnknownViews asserts that a model.ErrViewForUnknownEpoch sentinel
// is returned when querying identities by epoch for each of the input views.
func (suite *ConsensusSuite) AssertUnknownViews(views ...uint64) {
	for _, view := range views {
		_, err := suite.committee.IdentitiesByEpoch(view)
		suite.Assert().Error(err)
		suite.Assert().ErrorIs(err, model.ErrViewForUnknownEpoch)
	}
}

// AssertStoredEpochCounterRange asserts that the cached epochs are for exactly
// the given contiguous, inclusive counter range.
// Eg. for the input (2,4), the committee must have epochs cached with counters 2,3,4
func (suite *ConsensusSuite) AssertStoredEpochCounterRange(from, to uint64) {
	set := make(map[uint64]struct{})
	for i := from; i <= to; i++ {
		set[i] = struct{}{}
	}

	suite.committee.mu.RLock()
	defer suite.committee.mu.RUnlock()
	for epoch := range suite.committee.epochs {
		delete(set, epoch)
	}

	if !assert.Len(suite.T(), set, 0) {
		suite.T().Logf("%v should be empty, but isn't; expected epoch range [%d,%d]", set, from, to)
	}
}

// TestConstruction_CurrentEpoch tests construction with only a current epoch.
// Only the current epoch should be cached after construction.
func (suite *ConsensusSuite) TestConstruction_CurrentEpoch() {
	curEpoch := newMockCommittedEpoch(suite.currentEpochCounter, unittest.IdentityListFixture(10), 101, 200)
	suite.epochs.AddCommitted(curEpoch)

	suite.CreateAndStartCommittee()
	suite.Assert().Len(suite.committee.epochs, 1)
	suite.AssertStoredEpochCounterRange(suite.currentEpochCounter, suite.currentEpochCounter)
}

// TestConstruction_PreviousEpoch tests construction with a previous epoch.
// Both current and previous epoch should be cached after construction.
func (suite *ConsensusSuite) TestConstruction_PreviousEpoch() {
	prevEpoch := newMockCommittedEpoch(suite.currentEpochCounter-1, unittest.IdentityListFixture(10), 1, 100)
	curEpoch := newMockCommittedEpoch(suite.currentEpochCounter, unittest.IdentityListFixture(10), 101, 200)
	suite.epochs.AddCommitted(prevEpoch)
	suite.epochs.AddCommitted(curEpoch)

	suite.CreateAndStartCommittee()
	suite.Assert().Len(suite.committee.epochs, 2)
	suite.AssertStoredEpochCounterRange(suite.currentEpochCounter-1, suite.currentEpochCounter)
}

// TestConstruction_UncommittedNextEpoch tests construction with an uncommitted next epoch.
// Only the current epoch should be cached after construction.
func (suite *ConsensusSuite) TestConstruction_UncommittedNextEpoch() {
	suite.phase = flow.EpochPhaseSetup
	curEpoch := newMockCommittedEpoch(suite.currentEpochCounter, unittest.IdentityListFixture(10), 101, 200)
	nextEpoch := newMockTentativeEpoch(suite.currentEpochCounter+1, unittest.IdentityListFixture(10))
	suite.epochs.AddCommitted(curEpoch)
	suite.epochs.AddTentative(nextEpoch)

	suite.CreateAndStartCommittee()
	suite.Assert().Len(suite.committee.epochs, 1)
	suite.AssertStoredEpochCounterRange(suite.currentEpochCounter, suite.currentEpochCounter)
}

// TestConstruction_CommittedNextEpoch tests construction with a committed next epoch.
// Both current and next epochs should be cached after construction.
func (suite *ConsensusSuite) TestConstruction_CommittedNextEpoch() {
	curEpoch := newMockCommittedEpoch(suite.currentEpochCounter, unittest.IdentityListFixture(10), 101, 200)
	nextEpoch := newMockCommittedEpoch(suite.currentEpochCounter+1, unittest.IdentityListFixture(10), 201, 300)
	suite.epochs.AddCommitted(curEpoch)
	suite.epochs.AddCommitted(nextEpoch)
	suite.phase = flow.EpochPhaseCommitted

	suite.CreateAndStartCommittee()
	suite.Assert().Len(suite.committee.epochs, 2)
	suite.AssertStoredEpochCounterRange(suite.currentEpochCounter, suite.currentEpochCounter+1)
}

// TestProtocolEvents_CommittedEpoch tests that protocol events notifying of a newly
// committed epoch are handled correctly. A committed epoch should be cached, and
// repeated events should be no-ops.
func (suite *ConsensusSuite) TestProtocolEvents_CommittedEpoch() {
	curEpoch := newMockCommittedEpoch(suite.currentEpochCounter, unittest.IdentityListFixture(10), 101, 200)
	suite.epochs.AddCommitted(curEpoch)

	suite.CreateAndStartCommittee()

	nextEpoch := newMockCommittedEpoch(suite.currentEpochCounter+1, unittest.IdentityListFixture(10), 201, 300)

	firstBlockOfCommittedPhase := unittest.BlockHeaderFixture()
	suite.state.On("AtHeight", firstBlockOfCommittedPhase.Height).Return(suite.snapshot)
	suite.epochs.AddCommitted(nextEpoch)
	suite.committee.EpochCommittedPhaseStarted(suite.currentEpochCounter, firstBlockOfCommittedPhase)
	// wait for the protocol event to be processed (async)
	assert.Eventually(suite.T(), func() bool {
		_, err := suite.committee.IdentitiesByEpoch(201)
		return err == nil
	}, 30*time.Second, 50*time.Millisecond)

	suite.Assert().Len(suite.committee.epochs, 2)
	suite.AssertStoredEpochCounterRange(suite.currentEpochCounter, suite.currentEpochCounter+1)

	// should handle multiple deliveries of the protocol event
	suite.committee.EpochCommittedPhaseStarted(suite.currentEpochCounter, firstBlockOfCommittedPhase)
	suite.committee.EpochCommittedPhaseStarted(suite.currentEpochCounter, firstBlockOfCommittedPhase)
	suite.committee.EpochCommittedPhaseStarted(suite.currentEpochCounter, firstBlockOfCommittedPhase)

	suite.Assert().Len(suite.committee.epochs, 2)
	suite.AssertStoredEpochCounterRange(suite.currentEpochCounter, suite.currentEpochCounter+1)

}

// TestProtocolEvents_EpochExtended tests that protocol events notifying of an epoch extension are handled correctly.
// An EpochExtension event should result in a re-computation of the leader selection (including the new extension).
// Repeated events should be no-ops.
func (suite *ConsensusSuite) TestProtocolEvents_EpochExtended() {
	curEpoch := newMockCommittedEpoch(suite.currentEpochCounter, unittest.IdentityListFixture(10), 101, 200)
	suite.epochs.AddCommitted(curEpoch)

	suite.CreateAndStartCommittee()

	suite.AssertUnknownViews(100, 201, 300, 301)

	extension := flow.EpochExtension{
		FirstView: 201,
		FinalView: 300,
	}
	refBlock := unittest.BlockHeaderFixture()
	addExtension(curEpoch, extension)
	suite.state.On("AtHeight", refBlock.Height).Return(suite.snapshot)

	suite.committee.EpochExtended(suite.currentEpochCounter, refBlock, extension)
	// wait for the protocol event to be processed (async)
	require.Eventually(suite.T(), func() bool {
		_, err := suite.committee.IdentitiesByEpoch(extension.FirstView)
		return err == nil
	}, time.Second, 50*time.Millisecond)

	// we should have the same number of cached epochs (an existing epoch has been extended
	suite.Assert().Len(suite.committee.epochs, 1)
	suite.AssertStoredEpochCounterRange(suite.currentEpochCounter, suite.currentEpochCounter)

	// should handle multiple deliveries of the protocol event
	suite.committee.EpochExtended(suite.currentEpochCounter, refBlock, extension)
	suite.committee.EpochExtended(suite.currentEpochCounter, refBlock, extension)
	suite.committee.EpochExtended(suite.currentEpochCounter, refBlock, extension)

	suite.Assert().Len(suite.committee.epochs, 1)
	suite.AssertStoredEpochCounterRange(suite.currentEpochCounter, suite.currentEpochCounter)
	// check the boundary values of the original epoch and the extension, plus a random view within the extension
	suite.AssertKnownViews(101, 200, 201, unittest.Uint64InRange(201, 300), 300)
	suite.AssertUnknownViews(100, 301)
}

// TestProtocolEvents_EpochExtendedMultiple tests that protocol events notifying of an epoch extension are handled correctly.
// An EpochExtension event should result in a re-computation of the leader selection (including the new extension).
// The Committee should handle multiple subsequent, contiguous epoch extensions.
// Repeated events should be no-ops.
func (suite *ConsensusSuite) TestProtocolEvents_EpochExtendedMultiple() {
	curEpoch := newMockCommittedEpoch(suite.currentEpochCounter, unittest.IdentityListFixture(10), 101, 200)
	suite.epochs.AddCommitted(curEpoch)

	suite.CreateAndStartCommittee()

	expectedKnownViews := []uint64{101, unittest.Uint64InRange(101, 200), 200}
	suite.AssertUnknownViews(100, 201, 300, 301)
	suite.AssertKnownViews(expectedKnownViews...)

	// Add several extensions in series
	for i := 0; i < 10; i++ {
		finalView := curEpoch.FinalView()
		extension := flow.EpochExtension{
			FirstView: finalView + 1,
			FinalView: finalView + 100,
		}
		refBlock := unittest.BlockHeaderFixture()
		addExtension(curEpoch, extension)
		suite.state.On("AtHeight", refBlock.Height).Return(suite.snapshot)

		suite.committee.EpochExtended(suite.currentEpochCounter, refBlock, extension)
		// wait for the protocol event to be processed (async)
		require.Eventually(suite.T(), func() bool {
			_, err := suite.committee.IdentitiesByEpoch(extension.FirstView)
			return err == nil
		}, time.Second, 50*time.Millisecond)

		// we should have the same number of cached epochs (an existing epoch has been extended
		suite.Assert().Len(suite.committee.epochs, 1)
		suite.AssertStoredEpochCounterRange(suite.currentEpochCounter, suite.currentEpochCounter)

		// should respond to queries for view range of new extension
		expectedKnownViews = append(expectedKnownViews, extension.FirstView, unittest.Uint64InRange(extension.FirstView, extension.FinalView), extension.FinalView)
		suite.AssertKnownViews(expectedKnownViews...)
		// should return sentinel for view outside extension
		suite.AssertUnknownViews(100, extension.FinalView+1)
	}
}

// TestIdentitiesByBlock tests retrieving committee members by block.
// * should use up-to-block committee information
// * should exclude non-committee members
func (suite *ConsensusSuite) TestIdentitiesByBlock() {
	t := suite.T()

	realIdentity := unittest.IdentityFixture(unittest.WithRole(flow.RoleConsensus))
	joiningConsensusIdentity := unittest.IdentityFixture(unittest.WithRole(flow.RoleConsensus), unittest.WithParticipationStatus(flow.EpochParticipationStatusJoining))
	leavingConsensusIdentity := unittest.IdentityFixture(unittest.WithRole(flow.RoleConsensus), unittest.WithParticipationStatus(flow.EpochParticipationStatusLeaving))
	ejectedConsensusIdentity := unittest.IdentityFixture(unittest.WithRole(flow.RoleConsensus), unittest.WithParticipationStatus(flow.EpochParticipationStatusEjected))
	validNonConsensusIdentity := unittest.IdentityFixture(unittest.WithRole(flow.RoleVerification))
	validConsensusIdentities := []*flow.Identity{
		realIdentity,
		joiningConsensusIdentity,
		leavingConsensusIdentity,
		validNonConsensusIdentity,
		ejectedConsensusIdentity,
	}
	fakeID := unittest.IdentifierFixture()
	blockID := unittest.IdentifierFixture()

	// create a mock epoch for leader selection setup in constructor
	currEpoch := newMockCommittedEpoch(1, unittest.IdentityListFixture(10), 1, 100)
	suite.epochs.AddCommitted(currEpoch)

	suite.state.On("AtBlockID", blockID).Return(suite.snapshot)
	for _, identity := range validConsensusIdentities {
		i := identity // copy
		suite.snapshot.On("Identity", i.NodeID).Return(i, nil)
	}
	suite.snapshot.On("Identity", fakeID).Return(nil, protocol.IdentityNotFoundError{})

	suite.CreateAndStartCommittee()

	t.Run("non-existent identity should return InvalidSignerError", func(t *testing.T) {
		_, err := suite.committee.IdentityByBlock(blockID, fakeID)
		require.True(t, model.IsInvalidSignerError(err))
	})

	t.Run("existent but non-committee-member identity should return InvalidSignerError", func(t *testing.T) {
		t.Run("joining consensus node", func(t *testing.T) {
			_, err := suite.committee.IdentityByBlock(blockID, joiningConsensusIdentity.NodeID)
			require.True(t, model.IsInvalidSignerError(err))
		})

		t.Run("leaving consensus node", func(t *testing.T) {
			_, err := suite.committee.IdentityByBlock(blockID, leavingConsensusIdentity.NodeID)
			require.True(t, model.IsInvalidSignerError(err))
		})

		t.Run("ejected consensus node", func(t *testing.T) {
			_, err := suite.committee.IdentityByBlock(blockID, ejectedConsensusIdentity.NodeID)
			require.True(t, model.IsInvalidSignerError(err))
		})

		t.Run("otherwise valid non-consensus node", func(t *testing.T) {
			_, err := suite.committee.IdentityByBlock(blockID, validNonConsensusIdentity.NodeID)
			require.True(t, model.IsInvalidSignerError(err))
		})
	})

	t.Run("should be able to retrieve real identity", func(t *testing.T) {
		actual, err := suite.committee.IdentityByBlock(blockID, realIdentity.NodeID)
		require.NoError(t, err)
		require.Equal(t, realIdentity, actual)
	})
	t.Run("should propagate unexpected errors", func(t *testing.T) {
		mockErr := errors.New("unexpected")
		suite.snapshot.On("Identity", mock.Anything).Return(nil, mockErr)
		_, err := suite.committee.IdentityByBlock(blockID, unittest.IdentifierFixture())
		assert.ErrorIs(t, err, mockErr)
	})
}

// TestIdentitiesByEpoch tests that identities can be queried by epoch.
// * should use static epoch info (initial identities)
// * should exclude non-committee members
// * should correctly map views to epochs
// * should return ErrViewForUnknownEpoch sentinel for unknown epochs
func (suite *ConsensusSuite) TestIdentitiesByEpoch() {
	t := suite.T()

	// epoch 1 identities with varying conditions which would disqualify them
	// from committee participation
	realIdentity := unittest.IdentityFixture(unittest.WithRole(flow.RoleConsensus))
	zeroWeightConsensusIdentity := unittest.IdentityFixture(unittest.WithRole(flow.RoleConsensus),
		unittest.WithInitialWeight(0))
	ejectedConsensusIdentity := unittest.IdentityFixture(unittest.WithRole(flow.RoleConsensus),
		unittest.WithParticipationStatus(flow.EpochParticipationStatusEjected))
	validNonConsensusIdentity := unittest.IdentityFixture(unittest.WithRole(flow.RoleVerification))
	epoch1Identities := flow.IdentityList{realIdentity, zeroWeightConsensusIdentity, ejectedConsensusIdentity, validNonConsensusIdentity}

	// a single consensus node for epoch 2:
	epoch2Identity := unittest.IdentityFixture(unittest.WithRole(flow.RoleConsensus))
	epoch2Identities := flow.IdentityList{epoch2Identity}

	// create a mock epoch for leader selection setup in constructor
	epoch1 := newMockCommittedEpoch(suite.currentEpochCounter, epoch1Identities, 1, 100)
	// initially epoch 2 is not committed
	epoch2 := newMockCommittedEpoch(suite.currentEpochCounter+1, epoch2Identities, 101, 200)
	suite.epochs.AddCommitted(epoch1)

	suite.CreateAndStartCommittee()

	t.Run("only epoch 1 committed", func(t *testing.T) {
		t.Run("non-existent identity should return InvalidSignerError", func(t *testing.T) {
			_, err := suite.committee.IdentityByEpoch(unittest.Uint64InRange(1, 100), unittest.IdentifierFixture())
			require.True(t, model.IsInvalidSignerError(err))
		})

		t.Run("existent but non-committee-member identity should return InvalidSignerError", func(t *testing.T) {
			t.Run("zero-weight consensus node", func(t *testing.T) {
				_, err := suite.committee.IdentityByEpoch(unittest.Uint64InRange(1, 100), zeroWeightConsensusIdentity.NodeID)
				require.True(t, model.IsInvalidSignerError(err))
			})

			t.Run("otherwise valid non-consensus node", func(t *testing.T) {
				_, err := suite.committee.IdentityByEpoch(unittest.Uint64InRange(1, 100), validNonConsensusIdentity.NodeID)
				require.True(t, model.IsInvalidSignerError(err))
			})
		})

		t.Run("should be able to retrieve real identity", func(t *testing.T) {
			actual, err := suite.committee.IdentityByEpoch(unittest.Uint64InRange(1, 100), realIdentity.NodeID)
			require.NoError(t, err)
			require.Equal(t, realIdentity.IdentitySkeleton, *actual)
		})

		t.Run("should return ErrViewForUnknownEpoch for view outside existing epoch", func(t *testing.T) {
			_, err := suite.committee.IdentityByEpoch(unittest.Uint64InRange(101, 1_000_000), epoch2Identity.NodeID)
			require.ErrorIs(t, err, model.ErrViewForUnknownEpoch)
		})
	})

	// commit epoch 2
	suite.CommitEpoch(epoch2)

	t.Run("epoch 1 and 2 committed", func(t *testing.T) {
		t.Run("should be able to retrieve epoch 1 identity in epoch 1", func(t *testing.T) {
			actual, err := suite.committee.IdentityByEpoch(unittest.Uint64InRange(1, 100), realIdentity.NodeID)
			require.NoError(t, err)
			require.Equal(t, realIdentity.IdentitySkeleton, *actual)
		})

		t.Run("should be unable to retrieve epoch 1 identity in epoch 2", func(t *testing.T) {
			_, err := suite.committee.IdentityByEpoch(unittest.Uint64InRange(101, 200), realIdentity.NodeID)
			require.Error(t, err)
			require.True(t, model.IsInvalidSignerError(err))
		})

		t.Run("should be unable to retrieve epoch 2 identity in epoch 1", func(t *testing.T) {
			_, err := suite.committee.IdentityByEpoch(unittest.Uint64InRange(1, 100), epoch2Identity.NodeID)
			require.Error(t, err)
			require.True(t, model.IsInvalidSignerError(err))
		})

		t.Run("should be able to retrieve epoch 2 identity in epoch 2", func(t *testing.T) {
			actual, err := suite.committee.IdentityByEpoch(unittest.Uint64InRange(101, 200), epoch2Identity.NodeID)
			require.NoError(t, err)
			require.Equal(t, epoch2Identity.IdentitySkeleton, *actual)
		})

		t.Run("should return ErrViewForUnknownEpoch for view outside existing epochs", func(t *testing.T) {
			_, err := suite.committee.IdentityByEpoch(unittest.Uint64InRange(201, 1_000_000), epoch2Identity.NodeID)
			require.ErrorIs(t, err, model.ErrViewForUnknownEpoch)
		})
	})

}

// TestThresholds tests that the weight threshold methods return the
// correct thresholds for the previous and current epoch and that it returns the
// appropriate sentinel for the next epoch if it is not yet ready.
//
// There are 3 epochs in this test case, each with the same identities but different
// weights.
func (suite *ConsensusSuite) TestThresholds() {
	t := suite.T()

	identities := unittest.IdentityListFixture(10)

	prevEpoch := newMockCommittedEpoch(suite.currentEpochCounter-1, identities.Map(mapfunc.WithInitialWeight(100)), 1, 100)
	currEpoch := newMockCommittedEpoch(suite.currentEpochCounter, identities.Map(mapfunc.WithInitialWeight(200)), 101, 200)
	suite.epochs.AddCommitted(prevEpoch)
	suite.epochs.AddCommitted(currEpoch)

	suite.CreateAndStartCommittee()

	t.Run("next epoch not ready", func(t *testing.T) {
		t.Run("previous epoch", func(t *testing.T) {
			threshold, err := suite.committee.QuorumThresholdForView(unittest.Uint64InRange(1, 100))
			require.NoError(t, err)
			assert.Equal(t, WeightThresholdToBuildQC(1000), threshold)
			threshold, err = suite.committee.TimeoutThresholdForView(unittest.Uint64InRange(1, 100))
			require.NoError(t, err)
			assert.Equal(t, WeightThresholdToTimeout(1000), threshold)
		})

		t.Run("current epoch", func(t *testing.T) {
			threshold, err := suite.committee.QuorumThresholdForView(unittest.Uint64InRange(101, 200))
			require.NoError(t, err)
			assert.Equal(t, WeightThresholdToBuildQC(2000), threshold)
			threshold, err = suite.committee.TimeoutThresholdForView(unittest.Uint64InRange(101, 200))
			require.NoError(t, err)
			assert.Equal(t, WeightThresholdToTimeout(2000), threshold)
		})

		t.Run("after current epoch - should return ErrViewForUnknownEpoch", func(t *testing.T) {
			// get threshold for view in next epoch when it is not set up yet
			_, err := suite.committee.QuorumThresholdForView(unittest.Uint64InRange(201, 300))
			assert.ErrorIs(t, err, model.ErrViewForUnknownEpoch)
			_, err = suite.committee.TimeoutThresholdForView(unittest.Uint64InRange(201, 300))
			assert.ErrorIs(t, err, model.ErrViewForUnknownEpoch)
		})
	})

	// now, add a valid next epoch
	nextEpoch := newMockCommittedEpoch(suite.currentEpochCounter+1, identities.Map(mapfunc.WithInitialWeight(300)), 201, 300)
	suite.CommitEpoch(nextEpoch)

	t.Run("next epoch ready", func(t *testing.T) {
		t.Run("previous epoch", func(t *testing.T) {
			threshold, err := suite.committee.QuorumThresholdForView(unittest.Uint64InRange(1, 100))
			require.NoError(t, err)
			assert.Equal(t, WeightThresholdToBuildQC(1000), threshold)
			threshold, err = suite.committee.TimeoutThresholdForView(unittest.Uint64InRange(1, 100))
			require.NoError(t, err)
			assert.Equal(t, WeightThresholdToTimeout(1000), threshold)
		})

		t.Run("current epoch", func(t *testing.T) {
			threshold, err := suite.committee.QuorumThresholdForView(unittest.Uint64InRange(101, 200))
			require.NoError(t, err)
			assert.Equal(t, WeightThresholdToBuildQC(2000), threshold)
			threshold, err = suite.committee.TimeoutThresholdForView(unittest.Uint64InRange(101, 200))
			require.NoError(t, err)
			assert.Equal(t, WeightThresholdToTimeout(2000), threshold)
		})

		t.Run("next epoch", func(t *testing.T) {
			threshold, err := suite.committee.QuorumThresholdForView(unittest.Uint64InRange(201, 300))
			require.NoError(t, err)
			assert.Equal(t, WeightThresholdToBuildQC(3000), threshold)
			threshold, err = suite.committee.TimeoutThresholdForView(unittest.Uint64InRange(201, 300))
			require.NoError(t, err)
			assert.Equal(t, WeightThresholdToTimeout(3000), threshold)
		})

		t.Run("beyond known epochs", func(t *testing.T) {
			// get threshold for view in next epoch when it is not set up yet
			_, err := suite.committee.QuorumThresholdForView(unittest.Uint64InRange(301, 10_000))
			assert.ErrorIs(t, err, model.ErrViewForUnknownEpoch)
			_, err = suite.committee.TimeoutThresholdForView(unittest.Uint64InRange(301, 10_000))
			assert.ErrorIs(t, err, model.ErrViewForUnknownEpoch)
		})
	})
}

// TestLeaderForView tests that LeaderForView returns a valid leader
// for the previous and current epoch and that it returns the appropriate
// sentinel for the next epoch if it is not yet ready
func (suite *ConsensusSuite) TestLeaderForView() {
	t := suite.T()

	identities := unittest.IdentityListFixture(10)

	prevEpoch := newMockCommittedEpoch(suite.currentEpochCounter-1, identities, 1, 100)
	currEpoch := newMockCommittedEpoch(suite.currentEpochCounter, identities, 101, 200)
	suite.epochs.AddCommitted(currEpoch)
	suite.epochs.AddCommitted(prevEpoch)

	suite.CreateAndStartCommittee()

	t.Run("next epoch not ready", func(t *testing.T) {
		t.Run("previous epoch", func(t *testing.T) {
			// get leader for view in previous epoch
			leaderID, err := suite.committee.LeaderForView(unittest.Uint64InRange(1, 100))
			assert.NoError(t, err)
			_, exists := identities.ByNodeID(leaderID)
			assert.True(t, exists)
		})

		t.Run("current epoch", func(t *testing.T) {
			// get leader for view in current epoch
			leaderID, err := suite.committee.LeaderForView(unittest.Uint64InRange(101, 200))
			assert.NoError(t, err)
			_, exists := identities.ByNodeID(leaderID)
			assert.True(t, exists)
		})

		t.Run("after current epoch - should return ErrViewForUnknownEpoch", func(t *testing.T) {
			// get leader for view in next epoch when it is not set up yet
			_, err := suite.committee.LeaderForView(unittest.Uint64InRange(201, 300))
			assert.ErrorIs(t, err, model.ErrViewForUnknownEpoch)
		})
	})

	// now, add a valid next epoch
	nextEpoch := newMockCommittedEpoch(suite.currentEpochCounter+1, identities, 201, 300)
	suite.CommitEpoch(nextEpoch)

	t.Run("next epoch ready", func(t *testing.T) {
		t.Run("previous epoch", func(t *testing.T) {
			// get leader for view in previous epoch
			leaderID, err := suite.committee.LeaderForView(unittest.Uint64InRange(1, 100))
			require.NoError(t, err)
			_, exists := identities.ByNodeID(leaderID)
			assert.True(t, exists)
		})

		t.Run("current epoch", func(t *testing.T) {
			// get leader for view in current epoch
			leaderID, err := suite.committee.LeaderForView(unittest.Uint64InRange(101, 200))
			require.NoError(t, err)
			_, exists := identities.ByNodeID(leaderID)
			assert.True(t, exists)
		})

		t.Run("next epoch", func(t *testing.T) {
			// get leader for view in next epoch after it has been set up
			leaderID, err := suite.committee.LeaderForView(unittest.Uint64InRange(201, 300))
			require.NoError(t, err)
			_, exists := identities.ByNodeID(leaderID)
			assert.True(t, exists)
		})

		t.Run("beyond known epochs", func(t *testing.T) {
			_, err := suite.committee.LeaderForView(unittest.Uint64InRange(301, 1_000_000))
			assert.ErrorIs(t, err, model.ErrViewForUnknownEpoch)
		})
	})
}

// TestRemoveOldEpochs tests that old epochs are pruned
func TestRemoveOldEpochs(t *testing.T) {

	identities := unittest.IdentityListFixture(10)
	me := identities[0].NodeID

	// keep track of epoch counter and views
	firstEpochCounter := uint64(1)
	currentEpochCounter := firstEpochCounter
	epochFinalView := uint64(100)

	epoch1 := newMockCommittedEpoch(currentEpochCounter, identities, 1, epochFinalView)

	// create mocks
	state := protocolmock.NewState(t)
	snapshot := protocolmock.NewSnapshot(t)
	state.On("Final").Return(snapshot)

	epochQuery := mocks.NewEpochQuery(t, currentEpochCounter, epoch1)
	snapshot.On("Epochs").Return(epochQuery)
	currentEpochPhase := flow.EpochPhaseStaking
	snapshot.On("EpochPhase").Return(
		func() flow.EpochPhase { return currentEpochPhase },
		func() error { return nil },
	).Maybe()

	com, err := NewConsensusCommittee(state, me)
	require.NoError(t, err)

	ctx, cancel, errCh := irrecoverable.WithSignallerAndCancel(context.Background())
	com.Start(ctx)
	go unittest.FailOnIrrecoverableError(t, ctx.Done(), errCh)
	defer cancel()

	// we should start with only current epoch (epoch 1) pre-computed
	// since there is no previous epoch
	assert.Equal(t, 1, len(com.epochs))

	// test for 10 epochs
	for currentEpochCounter < 10 {

		// add another epoch
		firstView := epochFinalView + 1
		epochFinalView = epochFinalView + 100
		currentEpochCounter++
		nextEpoch := newMockCommittedEpoch(currentEpochCounter, identities, firstView, epochFinalView)
		epochQuery.AddCommitted(nextEpoch)

		currentEpochPhase = flow.EpochPhaseCommitted
		firstBlockOfCommittedPhase := unittest.BlockHeaderFixture()
		state.On("AtHeight", firstBlockOfCommittedPhase.Height).Return(snapshot)
		com.EpochCommittedPhaseStarted(currentEpochCounter, firstBlockOfCommittedPhase)
		// wait for the protocol event to be processed (async)
		require.Eventually(t, func() bool {
			_, err := com.IdentityByEpoch(unittest.Uint64InRange(firstView, epochFinalView), unittest.IdentifierFixture())
			return !errors.Is(err, model.ErrViewForUnknownEpoch)
		}, time.Second, time.Millisecond)

		// query a view from the new epoch
		_, err = com.LeaderForView(firstView)
		require.NoError(t, err)
		// transition to the next epoch
		epochQuery.Transition()

		t.Run(fmt.Sprintf("epoch %d", currentEpochCounter), func(t *testing.T) {
			// check we have the right number of epochs stored
			if currentEpochCounter <= 3 {
				assert.Equal(t, int(currentEpochCounter), len(com.epochs))
			} else {
				assert.Equal(t, 3, len(com.epochs))
			}

			// check we have the correct epochs stored
			for i := uint64(0); i < 3; i++ {
				counter := currentEpochCounter - i
				if counter < firstEpochCounter {
					break
				}
				_, exists := com.epochs[counter]
				assert.True(t, exists, "missing epoch with counter %d max counter is %d", counter, currentEpochCounter)
			}
		})
	}
}

// addExtension adds the extension to the mocked epoch, by updating its final view.
func addExtension(epoch *protocolmock.CommittedEpoch, ext flow.EpochExtension) {
	epoch.On("FinalView").Unset()
	epoch.On("FinalView").Return(ext.FinalView)
}

// newMockCommittedEpoch returns a new mocked committed epoch with the given fields
func newMockCommittedEpoch(counter uint64, identities flow.IdentityList, firstView uint64, finalView uint64) *protocolmock.CommittedEpoch {
	epoch := new(protocolmock.CommittedEpoch)
	epoch.On("Counter").Return(counter)
	epoch.On("RandomSource").Return(unittest.RandomBytes(32))
	epoch.On("InitialIdentities").Return(identities.ToSkeleton())
	epoch.On("FirstView").Return(firstView)
	epoch.On("FinalView").Return(finalView)
	epoch.On("DKG").Return(nil, nil)

	return epoch
}

// newMockTentativeEpoch returns a new mocked tentative epoch with the given fields
func newMockTentativeEpoch(counter uint64, identities flow.IdentityList) *protocolmock.TentativeEpoch {
	epoch := new(protocolmock.TentativeEpoch)
	epoch.On("Counter").Return(counter)
	epoch.On("InitialIdentities").Return(identities.ToSkeleton())
	return epoch
}
