package behavior

import (
	"encoding/json"
)

// Snapshot Event Sourcing Snapshotting is an optimisation that reduces time spent on reading event from an event store.
type Snapshot struct {
	Id      string       `json:"id"`
	Type    BehaviorType `json:"type"`
	State   []byte       `json:"state"`
	Version uint64       `json:"version"`
}

// ToSnapshot create new snapshot from the Aggregate state.
func ToSnapshot(behavior IBehavior) (*Snapshot, error) {
	state, err := json.Marshal(behavior.GetState())
	if err != nil {
		return nil, err
	}
	return &Snapshot{
		Id:      behavior.GetID().Id(),
		Type:    behavior.GetBehaviorType(),
		State:   state,
		Version: uint64(behavior.GetVersion()),
	}, nil
}
