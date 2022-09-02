package dtos

import "github.com/discomco/go-cart/robby/execute-game/-shared/model"

// Payload is the structure that will be carried with the initialize command.
type Payload struct {
	Details      *model.Details    `json:"details"`
	MapSize      *model.Dimensions `json:"map_size"`
	NbrOfPlayers int               `json:"nbr_of_players"`
}

func NewPayload(name string, x, y, z int, nbrOfPlayers int) *Payload {
	return &Payload{
		Details:      model.NewDetails(name),
		MapSize:      model.NewDimensions(x, y, z),
		NbrOfPlayers: nbrOfPlayers,
	}
}
