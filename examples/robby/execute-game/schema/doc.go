package schema

import (
	"github.com/discomco/go-cart/examples/robby/execute-game/schema/avatar"
	"github.com/discomco/go-cart/examples/robby/execute-game/schema/doc"
	"github.com/discomco/go-cart/sdk/schema"
)

type GameDoc struct {
	ID       *schema.Identity `json:"id,omitempty"`
	Details  *Details         `json:"details"`
	Status   doc.Status       `json:"status"`
	Settings *Settings        `json:"settings"`
}

type Settings struct {
	MapSize      *Dimensions
	NbrOfPlayers int
}

type Details struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type Avatar struct {
	ID      *schema.Identity `json:"id,omitempty"`
	Details *Details         `json:"details,omitempty"`
	Status  avatar.Status    `json:"status,omitempty"`
}

type Target struct {
	ID     *schema.Identity `json:"target_id,omitempty"`
	Avatar *Avatar          `json:"avatar,omitempty"`
}

type Action struct {
	Target *Target `json:"target,omitempty"`
}

type Turn struct {
	ID      *schema.Identity `json:"id,omitempty"`
	Avatar  Avatar           `json:"avatar,omitempty"`
	Actions []Action
}

type Dimensions struct {
	X int `json:"x,omitempty"`
	Y int `json:"y,omitempty"`
	Z int `json:"z,omitempty"`
}
