package root

import (
	"github.com/discomco/go-cart/sdk/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThatWeCanCreateARootID(t *testing.T) {
	// GIVEN
	assert.NotNil(t, testEnv)
	// WHEN
	ID, err := NewRootID()
	// THEN
	assert.NoError(t, err)
	assert.NotNil(t, ID)
}

func TestThatWeCanCreateARootIDFromString(t *testing.T) {
	// GIVEN
	assert.NotNil(t, testEnv)
	id := test.CLEAN_TEST_UUID
	// WHEN
	ID, err := NewRootIDFromString(id)
	// THEN
	assert.NoError(t, err)
	assert.NotNil(t, ID)
	assert.Equal(t, id, ID.Value)
	assert.Equal(t, ID_PREFIX, ID.Prefix)
}
