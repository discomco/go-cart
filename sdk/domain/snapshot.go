package domain

import (
	"encoding/json"
	"github.com/discomco/go-cart/sdk/model"
)

// Snapshot Event Sourcing Snapshotting is an optimisation that reduces time spent on reading event from an event store.
type Snapshot struct {
	ID      string        `json:"id"`
	Type    AggregateType `json:"type"`
	State   []byte        `json:"state"`
	Version uint64        `json:"version"`
}

// NewSnapshotFromAggregate create new snapshot from the Aggregate state.
func NewSnapshotFromAggregate[T model.IWriteModel](aggregate IAggregate) (*Snapshot, error) {
	aggregateBytes, err := json.Marshal(aggregate)
	if err != nil {
		return nil, err
	}

	return &Snapshot{
		ID:      aggregate.GetID().Id(),
		Type:    aggregate.GetAggregateType(),
		State:   aggregateBytes,
		Version: uint64(aggregate.GetVersion()),
	}, nil
}
