package behavior

import (
	"encoding/json"
	"github.com/discomco/go-cart/sdk/schema"
)

// Snapshot Event Sourcing Snapshotting is an optimisation that reduces time spent on reading event from an event store.
type Snapshot struct {
	ID      string       `json:"id"`
	Type    BehaviorType `json:"type"`
	State   []byte       `json:"state"`
	Version uint64       `json:"version"`
}

// NewSnapshotFromBehavior create new snapshot from the Aggregate state.
func NewSnapshotFromBehavior[T schema.IWriteSchema](behavior IBehavior) (*Snapshot, error) {
	state, err := json.Marshal(behavior.GetState())
	if err != nil {
		return nil, err
	}
	return &Snapshot{
		ID:      behavior.GetID().Id(),
		Type:    behavior.GetBehaviorType(),
		State:   state,
		Version: uint64(behavior.GetVersion()),
	}, nil
}
