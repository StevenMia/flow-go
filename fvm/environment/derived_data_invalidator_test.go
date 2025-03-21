package environment_test

import (
	"testing"

	"github.com/onflow/cadence/common"
	"github.com/stretchr/testify/require"

	"github.com/onflow/flow-go/fvm"
	"github.com/onflow/flow-go/fvm/environment"
	"github.com/onflow/flow-go/fvm/meter"
	"github.com/onflow/flow-go/fvm/storage"
	"github.com/onflow/flow-go/fvm/storage/derived"
	"github.com/onflow/flow-go/fvm/storage/snapshot"
	"github.com/onflow/flow-go/fvm/storage/state"
	"github.com/onflow/flow-go/model/flow"
	"github.com/onflow/flow-go/utils/unittest"
)

func TestDerivedDataProgramInvalidator(t *testing.T) {

	// create the following dependency graph
	// ```mermaid
	// graph TD
	// 	C-->D
	// 	C-->B
	// 	B-->A
	// ```

	addressA := flow.HexToAddress("0xa")
	cAddressA := common.MustBytesToAddress(addressA.Bytes())
	programALoc := common.AddressLocation{Address: cAddressA, Name: "A"}
	programA2Loc := common.AddressLocation{Address: cAddressA, Name: "A2"}
	programA := &derived.Program{
		Program: nil,
		Dependencies: derived.NewProgramDependencies().
			Add(programALoc),
	}

	addressB := flow.HexToAddress("0xb")
	cAddressB := common.MustBytesToAddress(addressB.Bytes())
	programBLoc := common.AddressLocation{Address: cAddressB, Name: "B"}
	programBDep := derived.NewProgramDependencies()
	programBDep.Add(programALoc)
	programBDep.Add(programBLoc)
	programB := &derived.Program{
		Program: nil,
		Dependencies: derived.NewProgramDependencies().
			Add(programALoc).
			Add(programBLoc),
	}

	addressD := flow.HexToAddress("0xd")
	cAddressD := common.MustBytesToAddress(addressD.Bytes())
	programDLoc := common.AddressLocation{Address: cAddressD, Name: "D"}
	programD := &derived.Program{
		Program: nil,
		Dependencies: derived.NewProgramDependencies().
			Add(programDLoc),
	}

	addressC := flow.HexToAddress("0xc")
	cAddressC := common.MustBytesToAddress(addressC.Bytes())
	programCLoc := common.AddressLocation{Address: cAddressC, Name: "C"}
	programC := &derived.Program{
		Program: nil,
		Dependencies: derived.NewProgramDependencies().
			Add(programALoc).
			Add(programBLoc).
			Add(programCLoc).
			Add(programDLoc),
	}

	t.Run("empty invalidator does not invalidate entries", func(t *testing.T) {
		invalidator := environment.DerivedDataInvalidator{}.ProgramInvalidator()

		require.False(t, invalidator.ShouldInvalidateEntries())
		require.False(t, invalidator.ShouldInvalidateEntry(programALoc, programA, nil))
		require.False(t, invalidator.ShouldInvalidateEntry(programBLoc, programB, nil))
		require.False(t, invalidator.ShouldInvalidateEntry(programCLoc, programC, nil))
		require.False(t, invalidator.ShouldInvalidateEntry(programDLoc, programD, nil))
	})
	t.Run("meter parameters invalidator invalidates all entries", func(t *testing.T) {
		invalidator := environment.DerivedDataInvalidator{
			ExecutionParametersUpdated: true,
		}.ProgramInvalidator()

		require.True(t, invalidator.ShouldInvalidateEntries())
		require.True(t, invalidator.ShouldInvalidateEntry(programALoc, programA, nil))
		require.True(t, invalidator.ShouldInvalidateEntry(programBLoc, programB, nil))
		require.True(t, invalidator.ShouldInvalidateEntry(programCLoc, programC, nil))
		require.True(t, invalidator.ShouldInvalidateEntry(programDLoc, programD, nil))
	})

	t.Run("contract A update invalidation", func(t *testing.T) {
		invalidator := environment.DerivedDataInvalidator{
			ContractUpdates: environment.ContractUpdates{
				Updates: []common.AddressLocation{
					programALoc,
				},
			},
		}.ProgramInvalidator()

		require.True(t, invalidator.ShouldInvalidateEntries())
		require.True(t, invalidator.ShouldInvalidateEntry(programALoc, programA, nil))
		require.True(t, invalidator.ShouldInvalidateEntry(programBLoc, programB, nil))
		require.True(t, invalidator.ShouldInvalidateEntry(programCLoc, programC, nil))
		require.False(t, invalidator.ShouldInvalidateEntry(programDLoc, programD, nil))
	})

	t.Run("contract D update invalidate", func(t *testing.T) {
		invalidator := environment.DerivedDataInvalidator{
			ContractUpdates: environment.ContractUpdates{
				Updates: []common.AddressLocation{
					programDLoc,
				},
			},
		}.ProgramInvalidator()

		require.True(t, invalidator.ShouldInvalidateEntries())
		require.False(t, invalidator.ShouldInvalidateEntry(programALoc, programA, nil))
		require.False(t, invalidator.ShouldInvalidateEntry(programBLoc, programB, nil))
		require.True(t, invalidator.ShouldInvalidateEntry(programCLoc, programC, nil))
		require.True(t, invalidator.ShouldInvalidateEntry(programDLoc, programD, nil))
	})

	t.Run("contract B update invalidate", func(t *testing.T) {
		invalidator := environment.DerivedDataInvalidator{
			ContractUpdates: environment.ContractUpdates{
				Updates: []common.AddressLocation{
					programBLoc,
				},
			},
		}.ProgramInvalidator()

		require.True(t, invalidator.ShouldInvalidateEntries())
		require.False(t, invalidator.ShouldInvalidateEntry(programALoc, programA, nil))
		require.True(t, invalidator.ShouldInvalidateEntry(programBLoc, programB, nil))
		require.True(t, invalidator.ShouldInvalidateEntry(programCLoc, programC, nil))
		require.False(t, invalidator.ShouldInvalidateEntry(programDLoc, programD, nil))
	})

	t.Run("contract invalidator C invalidates C", func(t *testing.T) {
		invalidator := environment.DerivedDataInvalidator{
			ContractUpdates: environment.ContractUpdates{
				Updates: []common.AddressLocation{
					programCLoc,
				},
			},
		}.ProgramInvalidator()

		require.True(t, invalidator.ShouldInvalidateEntries())
		require.False(t, invalidator.ShouldInvalidateEntry(programALoc, programA, nil))
		require.False(t, invalidator.ShouldInvalidateEntry(programBLoc, programB, nil))
		require.True(t, invalidator.ShouldInvalidateEntry(programCLoc, programC, nil))
		require.False(t, invalidator.ShouldInvalidateEntry(programDLoc, programD, nil))
	})

	t.Run("contract invalidator D invalidates C, D", func(t *testing.T) {
		invalidator := environment.DerivedDataInvalidator{
			ContractUpdates: environment.ContractUpdates{
				Updates: []common.AddressLocation{
					programDLoc,
				},
			},
		}.ProgramInvalidator()

		require.True(t, invalidator.ShouldInvalidateEntries())
		require.False(t, invalidator.ShouldInvalidateEntry(programALoc, programA, nil))
		require.False(t, invalidator.ShouldInvalidateEntry(programBLoc, programB, nil))
		require.True(t, invalidator.ShouldInvalidateEntry(programCLoc, programC, nil))
		require.True(t, invalidator.ShouldInvalidateEntry(programDLoc, programD, nil))
	})

	t.Run("new contract deploy on address A", func(t *testing.T) {
		invalidator := environment.DerivedDataInvalidator{
			ContractUpdates: environment.ContractUpdates{
				Deploys: []common.AddressLocation{
					programA2Loc,
				},
			},
		}.ProgramInvalidator()

		require.True(t, invalidator.ShouldInvalidateEntries())
		require.True(t, invalidator.ShouldInvalidateEntry(programALoc, programA, nil))
		require.True(t, invalidator.ShouldInvalidateEntry(programBLoc, programB, nil))
		require.True(t, invalidator.ShouldInvalidateEntry(programCLoc, programC, nil))
		require.False(t, invalidator.ShouldInvalidateEntry(programDLoc, programD, nil))
	})

	t.Run("contract delete on address A", func(t *testing.T) {
		invalidator := environment.DerivedDataInvalidator{
			ContractUpdates: environment.ContractUpdates{
				Deletions: []common.AddressLocation{
					programA2Loc,
				},
			},
		}.ProgramInvalidator()

		require.True(t, invalidator.ShouldInvalidateEntries())
		require.True(t, invalidator.ShouldInvalidateEntry(programALoc, programA, nil))
		require.True(t, invalidator.ShouldInvalidateEntry(programBLoc, programB, nil))
		require.True(t, invalidator.ShouldInvalidateEntry(programCLoc, programC, nil))
		require.False(t, invalidator.ShouldInvalidateEntry(programDLoc, programD, nil))
	})
}

func TestMeterParamOverridesInvalidator(t *testing.T) {
	invalidator := environment.DerivedDataInvalidator{}.
		ExecutionParametersInvalidator()

	require.False(t, invalidator.ShouldInvalidateEntries())
	require.False(t, invalidator.ShouldInvalidateEntry(
		struct{}{},
		derived.StateExecutionParameters{},
		nil))

	invalidator = environment.DerivedDataInvalidator{
		ContractUpdates:            environment.ContractUpdates{},
		ExecutionParametersUpdated: true,
	}.ExecutionParametersInvalidator()

	require.True(t, invalidator.ShouldInvalidateEntries())
	require.True(t, invalidator.ShouldInvalidateEntry(
		struct{}{},
		derived.StateExecutionParameters{},
		nil))
}

func TestMeterParamOverridesUpdated(t *testing.T) {
	memoryLimit := uint64(666)

	compKind := common.ComputationKind(12345)
	compWeight := uint64(10)
	computationWeights := meter.ExecutionEffortWeights{
		compKind: compWeight,
	}

	memKind := common.MemoryKind(23456)
	memWeight := uint64(20000)
	memoryWeights := meter.ExecutionMemoryWeights{
		memKind: memWeight,
	}

	snapshotTree := snapshot.NewSnapshotTree(nil)

	ctx := fvm.NewContext(fvm.WithChain(flow.Emulator.Chain()))

	vm := fvm.NewVirtualMachine()
	executionSnapshot, _, err := vm.Run(
		ctx,
		fvm.Bootstrap(
			unittest.ServiceAccountPublicKey,
			fvm.WithExecutionMemoryLimit(memoryLimit),
			fvm.WithExecutionEffortWeights(computationWeights),
			fvm.WithExecutionMemoryWeights(memoryWeights)),
		snapshotTree)
	require.NoError(t, err)

	blockDatabase := storage.NewBlockDatabase(
		snapshotTree.Append(executionSnapshot),
		0,
		nil)

	txnState, err := blockDatabase.NewTransaction(0, state.DefaultParameters())
	require.NoError(t, err)

	computer := fvm.NewExecutionParametersComputer(
		ctx.Logger,
		ctx,
		txnState,
	)

	overrides, err := computer.Compute(txnState, struct{}{})
	require.NoError(t, err)

	// Sanity check.  Note that bootstrap creates additional computation /
	// memory weight entries.  We'll only check the entries we added.
	require.NotNil(t, overrides.MemoryLimit)
	require.Equal(t, memoryLimit, *overrides.MemoryLimit)
	require.Equal(t, compWeight, overrides.ComputationWeights[compKind])
	require.Equal(t, memWeight, overrides.MemoryWeights[memKind])

	//
	// Actual test
	//

	ctx.TxBody = &flow.TransactionBody{}

	meterStateRead := &snapshot.ExecutionSnapshot{
		ReadSet: map[flow.RegisterID]struct{}{
			flow.NewRegisterID(ctx.Chain.ServiceAddress(), "meter"): {},
		},
	}

	checkForUpdates := func(id flow.RegisterID, expected bool) {
		snapshot := &snapshot.ExecutionSnapshot{
			WriteSet: map[flow.RegisterID]flow.RegisterValue{
				id: flow.RegisterValue("blah"),
			},
		}

		invalidator := environment.NewDerivedDataInvalidator(
			environment.ContractUpdates{},
			snapshot,
			meterStateRead)
		require.Equal(t, expected, invalidator.ExecutionParametersUpdated)
	}

	executionSnapshot, err = txnState.FinalizeMainTransaction()
	require.NoError(t, err)

	owner := ctx.Chain.ServiceAddress()
	otherOwner := unittest.RandomAddressFixtureForChain(ctx.Chain.ChainID())

	for _, registerId := range executionSnapshot.AllRegisterIDs() {
		checkForUpdates(registerId, false)
		checkForUpdates(
			flow.NewRegisterID(otherOwner, registerId.Key),
			false)
	}

	checkForUpdates(flow.NewRegisterID(owner, "meter2"), false)
	checkForUpdates(flow.NewRegisterID(owner, "meter"), true)
	checkForUpdates(flow.NewRegisterID(otherOwner, "meter2"), false)
	checkForUpdates(flow.NewRegisterID(otherOwner, "meter"), false)
}
