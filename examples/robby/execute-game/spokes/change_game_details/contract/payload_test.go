package contract

import (
	"github.com/discomco/go-cart/examples/robby/execute-game/schema"
	"github.com/discomco/go-cart/examples/robby/execute-game/schema/doc"
	"github.com/discomco/go-cart/sdk/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThatWeCanCreateAnInitializePayload(t *testing.T) {
	// GIVEN
	assert.NotNil(t, testEnv)
	assert.NotNil(t, testLogger)
	// WHEN
	ID, err := doc.NewGameIDFromString(test.CLEAN_TEST_UUID)
	assert.NoError(t, err)
	assert.NotNil(t, ID)
	gameName := "Hello, Robby"
	gameDescription := "Welcome to a new game of Robby!"
	details := schema.NewDetails(gameName)
	details.Description = gameDescription
	pl := NewPayload(details)
	// THEN
	assert.NotNil(t, pl)
	assert.NotNil(t, pl.Details)
	assert.Equal(t, gameName, pl.Details.Name)
	assert.Equal(t, gameDescription, pl.Details.Description)
}
