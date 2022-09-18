package contract

import (
	"github.com/discomco/go-cart/examples/robby/execute-game/schema"
)

// Payload is the structure that will be carried with the initialize command.
type Payload struct {
	Settings *schema.Settings `json:"settings"`
}

//NewPayload create a new payload for the change_game_settings module.
func NewPayload(settings *schema.Settings) *Payload {
	return &Payload{
		Settings: settings,
	}
}
