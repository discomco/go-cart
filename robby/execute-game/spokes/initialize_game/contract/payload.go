package contract

import (
	"github.com/discomco/go-cart/robby/execute-game/schema"
)

// Payload is the structure that will be carried with the initialize command.
type Payload struct {
	GameId       string             `json:"behavior_id"`
	Details      *schema.Details    `json:"details"`
	MapSize      *schema.Dimensions `json:"map_size"`
	NbrOfPlayers int                `json:"nbr_of_players"`
}

//NewPayload create a new payload for the initialize_game module.
func NewPayload(gameId string, name string, x, y, z int, nbrOfPlayers int) *Payload {
	return &Payload{
		GameId:       gameId,
		Details:      schema.NewDetails(name),
		MapSize:      schema.NewDimensions(x, y, z),
		NbrOfPlayers: nbrOfPlayers,
	}
}
