package model

import (
	"github.com/discomco/go-cart/sdk/core"
)

type Root struct {
	ID      *core.Identity `json:"id,omitempty"`
	Details *Details
}

type Details struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type Avatar struct {
	ID      *core.Identity `json:"id,omitempty"`
	Details *Details       `json:"details,omitempty"`
}

type Turn struct {
	ID     *core.Identity `json:"id,omitempty"`
	Avatar Avatar         `json:"avatar,omitempty"`
}
