package avatar

import (
	"github.com/discomco/go-cart/sdk/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThatWeCanCreateAnAvatarID(t *testing.T) {
	// GIVEN
	assert.NotNil(t, testEnv)
	// WHEN
	ID, err := NewAvatarID()
	// THEN
	assert.NoError(t, err)
	assert.NotNil(t, ID)
	assert.Equal(t, ID_PREFIX, ID.Prefix)
}

func TestThatWeCanCreateAnAvatarIDFromAString(t *testing.T) {
	// GIVEN
	assert.NotNil(t, testEnv)
	// WHEN
	Id := test.CLEAN_TEST_UUID
	ID, err := NewAvatarIDFrom(Id)
	// THEN
	assert.NoError(t, err)
	assert.NotNil(t, ID)
	assert.Equal(t, ID_PREFIX, ID.Prefix)
	assert.Equal(t, Id, ID.Value)
}
