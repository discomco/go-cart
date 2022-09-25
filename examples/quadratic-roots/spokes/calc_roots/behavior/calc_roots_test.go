package behavior

import (
	"github.com/discomco/go-cart/examples/quadratic-roots/schema"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThatCalculateRootsReturnsAnOutput(t *testing.T) {
	// GIVEN
	input := schema.RandomInput(10)
	// WHEN
	output := calcRoots(input)
	// THEN
	assert.NotNil(t, output)
}
