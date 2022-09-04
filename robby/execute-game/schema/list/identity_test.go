package list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThatWeCanGetTheDefaultListID(t *testing.T) {
	// GIVEN
	// WHEN
	ID, err := DefaultID()
	// THEN
	assert.NoError(t, err)
	assert.NotNil(t, ID)
}
