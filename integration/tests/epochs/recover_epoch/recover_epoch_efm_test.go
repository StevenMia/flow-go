package recover_epoch

import (
	"strings"
	"testing"
	"time"

	"github.com/onflow/cadence"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	sdk "github.com/onflow/flow-go-sdk"

	"github.com/onflow/flow-go/integration/utils"
	"github.com/onflow/flow-go/model/flow"
)

func TestRecoverEpoch(t *testing.T) {
	suite.Run(t, new(RecoverEpochSuite))
}

type RecoverEpochSuite struct {
	Suite
}

// TestRecoverEpoch ensures that the recover epoch governance transaction flow works as expected and a network that
// enters Epoch Fallback Mode can successfully recover. This test will do the following:
// 1. Triggers EFM by turning off the sole collection node before the end of the DKG forcing the DKG to fail.
// 2. Generates epoch recover transaction args using the epoch efm-recover-tx-args.
// 3. Submit recover epoch transaction.
// 4. Ensure expected EpochRecover event is emitted.
// TODO(EFM, #6164): Currently, this test does not test the processing of the EpochRecover event
func (s *RecoverEpochSuite) TestRecoverEpoch() {
	// 1. Manually trigger EFM
	// wait until the epoch setup phase to force network into EFM
	s.AwaitEpochPhase(s.Ctx, 0, flow.EpochPhaseSetup, 10*time.Second, 500*time.Millisecond)

	// We set the DKG phase view len to 10 which is very short and should cause the network to go into EFM
	// without pausing the collection node. This is not the case, we still need to pause the collection node.
	//TODO(EFM, #6164): Why short DKG phase len of 10 views does not trigger EFM without pausing container ; see https://github.com/onflow/flow-go/issues/6164
	ln := s.GetContainersByRole(flow.RoleCollection)[0]
	require.NoError(s.T(), ln.Pause())
	s.AwaitFinalizedView(s.Ctx, s.GetDKGEndView(), 2*time.Minute, 500*time.Millisecond)
	// start the paused collection node now that we are in EFM
	require.NoError(s.T(), ln.Start())

	// get final view form the latest snapshot
	epoch1FinalView, err := s.Net.BootstrapSnapshot.Epochs().Current().FinalView()
	require.NoError(s.T(), err)

	// wait for at least the first block of the next epoch to be sealed so that we can
	// ensure that we are still in the same epoch after the final view of that epoch indicating we are in EFM
	s.TimedLogf("waiting for epoch transition (finalized view %d)", epoch1FinalView+1)
	s.AwaitFinalizedView(s.Ctx, epoch1FinalView+1, 2*time.Minute, 500*time.Millisecond)
	s.TimedLogf("observed finalized view %d", epoch1FinalView+1)

	// assert transition to second epoch did not happen
	// if counter is still 0, epoch emergency fallback was triggered as expected
	s.AssertInEpoch(s.Ctx, 0)

	// 2. Generate transaction arguments for epoch recover transaction.
	// generate epoch recover transaction args
	collectionClusters := s.NumOfCollectionClusters
	numViewsInRecoveryEpoch := s.EpochLen
	numViewsInStakingAuction := s.StakingAuctionLen
	epochCounter := uint64(1)

	txArgs := s.executeEFMRecoverTXArgsCMD(
		collectionClusters,
		numViewsInRecoveryEpoch,
		numViewsInStakingAuction,
		epochCounter,
		// cruise control is disabled for integration tests
		// targetDuration and targetEndTime will be ignored
		3000,
		// unsafeAllowOverWrite set to false, initialize new epoch
		false,
	)

	// 3. Submit recover epoch transaction to the network.
	// submit the recover epoch transaction
	env := utils.LocalnetEnv()
	result := s.recoverEpoch(env, txArgs)
	require.NoError(s.T(), result.Error)
	require.Equal(s.T(), result.Status, sdk.TransactionStatusSealed)

	// 3. Ensure EpochRecover event was emitted.
	eventType := ""
	for _, evt := range result.Events {
		if strings.Contains(evt.Type, "FlowEpoch.EpochRecover") {
			eventType = evt.Type
			break
		}
	}
	require.NotEmpty(s.T(), eventType, "expected FlowEpoch.EpochRecover event type")
	events, err := s.Client.GetEventsForBlockIDs(s.Ctx, eventType, []sdk.Identifier{result.BlockID})
	require.NoError(s.T(), err)
	require.Equal(s.T(), events[0].Events[0].Type, eventType)

	startViewOfNextEpoch := uint64(txArgs[1].(cadence.UInt64))
	// wait for first view of recovery epoch
	s.TimedLogf("waiting to transition into recovery epoch (finalized view %d)", startViewOfNextEpoch)
	s.AwaitFinalizedView(s.Ctx, startViewOfNextEpoch, 2*time.Minute, 500*time.Millisecond)
	s.TimedLogf("observed finalized first view of recovery epoch %d", startViewOfNextEpoch)

	// ensure we transition into recovery epoch
	s.AssertInEpoch(s.Ctx, 1)
}
