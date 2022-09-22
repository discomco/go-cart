package contract

import (
	"github.com/discomco/go-cart/sdk/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThatFbkReturnsFlattenedErrors(t *testing.T) {
	// GIVEN
	f := NewFbk(test.CLEAN_TEST_UUID, -1, "An Error")
	f.SetError("Another Error")
	// WHEN
	s := f.GetFlattenedErrors()
	// THEN
	assert.Contains(t, s, "errors:\nAn Error\nAnother Error")
}
