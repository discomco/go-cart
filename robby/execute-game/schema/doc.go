package schema

import (
	"github.com/discomco/go-cart/robby/execute-game/schema/avatar"
	"github.com/discomco/go-cart/robby/execute-game/schema/doc"
	"github.com/discomco/go-cart/sdk/core"
)

type GameDoc struct {
	ID      *core.Identity `json:"id,omitempty"`
	Details *Details       `json:"details"`
	Status  doc.Status     `json:"status"`
}

type Details struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type Avatar struct {
	ID      *core.Identity `json:"id,omitempty"`
	Details *Details       `json:"details,omitempty"`
	Status  avatar.Status  `json:"status,omitempty"`
}

type Target struct {
	ID     *core.Identity `json:"target_id,omitempty"`
	Avatar *Avatar        `json:"avatar,omitempty"`
}

type Action struct {
	Target *Target `json:"target,omitempty"`
}

type Turn struct {
	ID      *core.Identity `json:"id,omitempty"`
	Avatar  Avatar         `json:"avatar,omitempty"`
	Actions []Action
}

type Dimensions struct {
	X int `json:"x,omitempty"`
	Y int `json:"y,omitempty"`
	Z int `json:"z,omitempty"`
}
