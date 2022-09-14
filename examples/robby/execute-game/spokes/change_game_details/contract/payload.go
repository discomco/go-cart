package contract

import (
	"github.com/discomco/go-cart/robby/execute-game/schema"
)

// Payload is the structure that will be carried with the initialize command.
type Payload struct {
	Details *schema.Details `json:"details"`
}

//NewPayload create a new payload for the change_game_details module.
func NewPayload(details *schema.Details) *Payload {
	return &Payload{
		Details: details,
	}
}
