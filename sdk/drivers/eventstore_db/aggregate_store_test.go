package eventstore_db

import (
	"github.com/discomco/go-cart/sdk/comps"
	"github.com/discomco/go-cart/sdk/core/logger"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThatWeCanResolveAnAggregateStore(t *testing.T) {
	// GIVEN
	assert.NotNil(t, testEnv)
	// WHEN
	var lgg logger.IAppLogger
	var as comps.IBehaviorStore
	err := testEnv.Invoke(func(log logger.IAppLogger, asFtor comps.BehSFtor) {
		lgg = log
		as = asFtor()
	})
	assert.Nil(t, err)
	assert.NotNil(t, lgg)
	assert.NotNil(t, as)
	// THEN
}
