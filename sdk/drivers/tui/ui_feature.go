package tui

import "github.com/discomco/go-cart/sdk/features"

type IMicroApp interface {
	features.ISpoke
}

type MicroApp struct {
}
