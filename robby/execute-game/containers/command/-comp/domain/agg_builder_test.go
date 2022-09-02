package domain

import (
	"github.com/discomco/go-cart/sdk/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThatWeCanResolveTheAggregateBuilder(t *testing.T) {
	// GIVEN
	assert.NotNil(t, testEnv)
	// WHEN
	var agg domain.IAggregate
	err := testEnv.Invoke(func(buildAgg domain.AggBuilder) {
		agg = buildAgg()
	})
	assert.NoError(t, err)
	assert.NotNil(t, agg)
}
