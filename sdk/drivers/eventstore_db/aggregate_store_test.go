package eventstore_db

import (
	"github.com/discomco/go-cart/sdk/core/logger"
	"github.com/discomco/go-cart/sdk/features"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThatWeCanResolveAnAggregateStore(t *testing.T) {
	// GIVEN
	assert.NotNil(t, testEnv)
	// WHEN
	var lgg logger.IAppLogger
	var as features.IAggregateStore
	err := testEnv.Invoke(func(log logger.IAppLogger, asCtor features.ASFtor) {
		lgg = log
		as = asCtor()
	})
	assert.Nil(t, err)
	assert.NotNil(t, lgg)
	assert.NotNil(t, as)
	// THEN
}
