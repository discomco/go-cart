package ftor

import (
	"github.com/discomco/go-cart/examples/quadratic-roots/schema"
	sdk_behavior "github.com/discomco/go-cart/sdk/behavior"
	sdk_schema "github.com/discomco/go-cart/sdk/schema"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThatWeCanResolveABehavior(t *testing.T) {
	// GIVEN
	assert.NotNil(t, testEnv)
	assert.NotNil(t, testLogger)
	// WHEN
	var agg IBehavior
	err := testEnv.Invoke(func(rootFor sdk_schema.DocFtor[schema.QuadraticDoc],
		aggFtor sdk_behavior.GenBehaviorFtor[schema.QuadraticDoc]) {
		agg = aggFtor()
	})
	// THEN
	assert.NotNil(t, agg)
	assert.NoError(t, err)
}
