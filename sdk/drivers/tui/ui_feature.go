package tui

import "github.com/discomco/go-cart/sdk/features"

type IMicroApp interface {
	features.IFeature
}

type MicroApp struct {
}
