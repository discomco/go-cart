package redis

import (
	"github.com/discomco/go-cart/examples/quadratic-roots/schema"
	"github.com/discomco/go-cart/sdk/behavior"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThatWeCanResolveADocStore(t *testing.T) {
	// GIVEN
	assert.NotNil(t, testEnv)
	assert.NotNil(t, testLogger)
	// WHEN
	var docStore IDocStore
	err := testEnv.Invoke(func(newDocStore behavior.StoreFtor[schema.QuadraticDoc]) {
		docStore = newDocStore()
	})
	// AND
	if err != nil {
		testLogger.Fatal(err)
		t.Fatal(err)
	}
	// THEN
	assert.NotNil(t, docStore)

}
