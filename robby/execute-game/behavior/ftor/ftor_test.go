package ftor

import (
	"github.com/discomco/go-cart/robby/execute-game/schema"
	"github.com/discomco/go-cart/sdk/domain"
	"github.com/discomco/go-cart/sdk/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThatWeCanResolveAnAggregate(t *testing.T) {
	// GIVEN
	assert.NotNil(t, testEnv)
	assert.NotNil(t, testLogger)
	// WHEN
	var agg IBehavior
	err := testEnv.Invoke(func(rootFor model.DocFtor[schema.GameDoc],
		aggFtor domain.GenAggFtor[schema.GameDoc]) {
		agg = aggFtor()
	})
	// THEN
	assert.NotNil(t, agg)
	assert.NoError(t, err)
}
