package tui

import "github.com/discomco/go-cart/sdk/spokes"

type IMicroApp interface {
	spokes.ISpoke
}

type MicroApp struct {
}
