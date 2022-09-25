package comps

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThatWeCanResolveAToRedisDocProjection(t *testing.T) {
	// GIVEN
	assert.NotNil(t, testEnv)

	// WHEN
	var toRedisDoc IToRedisDoc
	err := testEnv.Invoke(func(proj IToRedisDoc) {
		toRedisDoc = proj
	})
	if err != nil {
		testLogger.Fatal(err)
		t.Fatal(err)
	}
	// THEN
	assert.NotNil(t, toRedisDoc)
}
