package operation

import (
	"github.com/dgraph-io/badger/v2"

	"github.com/onflow/flow-go/model/flow"
)

func RetrieveJobLatestIndex(queue string, index *uint64) func(*badger.Txn) error {
	return retrieve(makePrefix(codeJobQueuePointer, queue), index)
}

func InitJobLatestIndex(queue string, index uint64) func(*badger.Txn) error {
	return insert(makePrefix(codeJobQueuePointer, queue), index)
}

func SetJobLatestIndex(queue string, index uint64) func(*badger.Txn) error {
	return update(makePrefix(codeJobQueuePointer, queue), index)
}

// RetrieveJobAtIndex returns the entity at the given index
func RetrieveJobAtIndex(queue string, index uint64, entity *flow.Identifier) func(*badger.Txn) error {
	return retrieve(makePrefix(codeJobQueue, queue, index), entity)
}

// InsertJobAtIndex insert an entity ID at the given index
func InsertJobAtIndex(queue string, index uint64, entity flow.Identifier) func(*badger.Txn) error {
	return insert(makePrefix(codeJobQueue, queue, index), entity)
}
