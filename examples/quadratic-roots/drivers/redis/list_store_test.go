package redis

import (
	"github.com/discomco/go-cart/examples/quadratic-roots/schema"
	"github.com/discomco/go-cart/sdk/behavior"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThatWeCanResolveAListStore(t *testing.T) {
	// GIVEN
	assert.NotNil(t, testEnv)
	assert.NotNil(t, testLogger)
	// WHEN
	var listStore IListStore
	if err := testEnv.Invoke(func(newListStore behavior.StoreFtor[schema.QuadraticList]) {
		listStore = newListStore()
	}); err != nil {
		testLogger.Fatal(err)
		t.Fatal(err)
	}
	// THEN
	assert.NotNil(t, listStore)
}
