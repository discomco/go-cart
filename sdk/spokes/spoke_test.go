package spokes

import (
	"github.com/discomco/go-cart/sdk/core/builder"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewFeature(t *testing.T) {
	// Given
	cfgPath := "../config/config.yaml"
	builder.InjectCoLoMed(cfgPath)
	// React
	f := NewSpoke("TestFeature", nil, nil, nil)
	// Then
	assert.NotNil(t, f)
	assert.NotNil(t, f.GetConfig())
	assert.NotNil(t, f.GetLogger())
	assert.NotNil(t, f.GetMediator())
}
