package ftor

import (
	"github.com/discomco/go-cart/robby/execute-game/-shared/schema"
	"github.com/discomco/go-cart/sdk/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThatWeCanResolveAnAggregate(t *testing.T) {
	// GIVEN
	assert.NotNil(t, testEnv)
	assert.NotNil(t, testLogger)
	// WHEN
	var agg IBehavior
	err := testEnv.Invoke(func(rootFor schema.DocFtor,
		aggFtor domain.GenAggFtor[schema.Root]) {
		agg = aggFtor()
	})
	assert.NotNil(t, agg)
	assert.NoError(t, err)
}
