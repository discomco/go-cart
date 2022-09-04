package schema

import (
	"github.com/discomco/go-cart/robby/execute-game/schema/doc"
	"github.com/discomco/go-cart/sdk/core"
)

type GameList struct {
	ID    *core.Identity       `json:"id,omitempty"`
	Items map[string]*GameItem `json:"items"`
}

type GameItem struct {
	Id              string     `json:"id,omitempty"`
	Name            string     `json:"name"`
	NumberOfPlayers int        `json:"numberOfPlayers"`
	Status          doc.Status `json:"status"`
}
