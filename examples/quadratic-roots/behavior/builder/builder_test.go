package builder

import (
	"github.com/discomco/go-cart/sdk/behavior"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThatWeCanResolveTheBehaviorBuilder(t *testing.T) {
	// GIVEN
	assert.NotNil(t, testEnv)
	// WHEN
	var agg behavior.IBehavior
	err := testEnv.Invoke(func(buildAgg behavior.BehaviorBuilder) {
		agg = buildAgg()
	})
	assert.NoError(t, err)
	assert.NotNil(t, agg)
}
