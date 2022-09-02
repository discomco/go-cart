package dtos

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThatWeCanCreateAnInitializePayload(t *testing.T) {
	// GIVEN
	assert.NotNil(t, testEnv)
	assert.NotNil(t, testLogger)
	// WHEN
	gameName := "Hello, Robby"
	xSize := 42
	ySize := 42
	zSize := 42
	nbrOfPlayers := 12
	pl := NewPayload(gameName, xSize, ySize, zSize, nbrOfPlayers)
	// THEN
	assert.NotNil(t, pl)
	assert.NotNil(t, pl.Details)
	assert.NotNil(t, pl.MapSize)
	assert.Equal(t, gameName, pl.Details.Name)
	assert.Equal(t, xSize, pl.MapSize.X)
	assert.Equal(t, ySize, pl.MapSize.Y)
	assert.Equal(t, zSize, pl.MapSize.Z)
}
