package eventstore_db

import (
	"github.com/discomco/go-cart/sdk/core/logger"
	"github.com/discomco/go-cart/sdk/reactors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThatWeCanResolveAnAggregateStore(t *testing.T) {
	// GIVEN
	assert.NotNil(t, testEnv)
	// WHEN
	var lgg logger.IAppLogger
	var as reactors.IBehaviorStore
	err := testEnv.Invoke(func(log logger.IAppLogger, asFtor reactors.BehSFtor) {
		lgg = log
		as = asFtor()
	})
	assert.Nil(t, err)
	assert.NotNil(t, lgg)
	assert.NotNil(t, as)
	// THEN
}
