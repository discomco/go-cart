package kafka

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThatWeCanCreateANewKafkaBus(t *testing.T) {
	// GIVEN
	assert.NotNil(t, testEnv)
	// WHEN
	bus := newKafkaBus()
	// THEN
	assert.NotNil(t, bus)
}
