package domain

import (
	"github.com/discomco/go-cart/robby/execute-game/-shared/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThatWeCanResolveAnAggregate(t *testing.T) {
	// GIVEN
	assert.NotNil(t, testEnv)
	assert.NotNil(t, testLogger)
	// WHEN
	var agg IAggregate
	err := testEnv.Invoke(func(rootFor model.DocFtor,
		aggFtor AggFtor) {
		agg = aggFtor()
	})
	assert.NotNil(t, agg)
	assert.NoError(t, err)
}
