package ftor

import (
	"github.com/discomco/go-cart/examples/robby/execute-game/schema"
	"github.com/discomco/go-cart/sdk/behavior"
	sdk_schema "github.com/discomco/go-cart/sdk/schema"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThatWeCanResolveAnAggregate(t *testing.T) {
	// GIVEN
	assert.NotNil(t, testEnv)
	assert.NotNil(t, testLogger)
	// WHEN
	var agg IBehavior
	err := testEnv.Invoke(func(rootFor sdk_schema.DocFtor[schema.GameDoc],
		aggFtor behavior.GenBehaviorFtor[schema.GameDoc]) {
		agg = aggFtor()
	})
	// THEN
	assert.NotNil(t, agg)
	assert.NoError(t, err)
}
