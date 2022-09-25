package contract

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThatWeCanCreateAHopePayload(t *testing.T) {
	// GIVEN
	// WHEN
	pl := NewHopePayload()
	// THEN
	assert.NotNil(t, pl)
}
