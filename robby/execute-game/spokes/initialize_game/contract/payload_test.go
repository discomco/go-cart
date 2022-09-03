package contract

import (
	"github.com/discomco/go-cart/robby/execute-game/schema/doc"
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
	xSize := 42
	ySize := 42
	zSize := 42
	nbrOfPlayers := 12
	pl := NewPayload(ID.Id(), gameName, xSize, ySize, zSize, nbrOfPlayers)
	// THEN
	assert.NotNil(t, pl)
	assert.NotNil(t, pl.Details)
	assert.NotNil(t, pl.MapSize)
	assert.Equal(t, gameName, pl.Details.Name)
	assert.Equal(t, xSize, pl.MapSize.X)
	assert.Equal(t, ySize, pl.MapSize.Y)
	assert.Equal(t, zSize, pl.MapSize.Z)
}
