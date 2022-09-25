package comps

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThatWeCanResolveAToRedisListProjection(t *testing.T) {
	// GIVEN
	assert.NotNil(t, testEnv)
	// WHEN
	var toRedisList IToRedisList
	if err := testEnv.Invoke(func(proj IToRedisList) {
		toRedisList = proj
	}); err != nil {
		testLogger.Fatal(err)
		t.Fatal(err)
	}
	// THEN
	assert.NotNil(t, toRedisList)
}
