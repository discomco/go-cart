package features

import (
	"github.com/discomco/go-cart/core/builder"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewProjectionFeature(t *testing.T) {
	// Given
	cfgPath := "../config/config.yaml"
	ioc := builder.InjectCoLoMed(cfgPath)
	//
	//proj := eventstore_db.NewProjector(ioc, "container", "groupName")
	//// When
	//pf := NewProjectionAppFeature("testProjectorFeature", proj, nil)
	// Then
	assert.NotNil(t, ioc)
}
