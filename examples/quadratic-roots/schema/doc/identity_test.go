package doc

import (
	"github.com/discomco/go-cart/sdk/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThatWeCanCreateADocID(t *testing.T) {
	// GIVEN
	assert.NotNil(t, testEnv)
	// WHEN
	ID, err := NewDocID()
	// THEN
	assert.NoError(t, err)
	assert.NotNil(t, ID)
}

func TestThatWeCanCreateADocIDFromString(t *testing.T) {
	// GIVEN
	assert.NotNil(t, testEnv)
	id := test.CLEAN_TEST_UUID
	// WHEN
	ID, err := NewDocIDFromString(id)
	// THEN
	assert.NoError(t, err)
	assert.NotNil(t, ID)
	assert.Equal(t, id, ID.Value)
	assert.Equal(t, IdPrefix, ID.Prefix)
}
