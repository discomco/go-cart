package core

import (
	"fmt"
	"strings"
	"testing"

	"github.com/discomco/go-cart/test"
	"github.com/stretchr/testify/assert"
)

func TestNewIdentity(t *testing.T) {
	// Given
	prefix := "prefix"
	// ApplyEvent
	id, err := NewIdentity(prefix)
	assert.NoError(t, err)
	// Then
	assert.NotNil(t, id)
}

func TestNewDefaultIdentity(t *testing.T) {
	id, err := NewIdentity("")
	assert.NoError(t, err)
	assert.NotNil(t, id)
	assert.Equal(t, "id", id.Prefix)
}

func TestNewIdentityFrom(t *testing.T) {
	// Given
	pf := test.TEST_PREFIX
	uuid := test.TEST_UUID
	// ApplyEvent
	id, err := NewIdentityFrom(pf, uuid)
	assert.NoError(t, err)
	// Then
	assert.NotNil(t, id)
	assert.Equal(t, pf, id.Prefix)
	assert.Equal(t, test.CLEAN_TEST_UUID, id.Value)
	assert.Equal(t, fmt.Sprintf("%s-%s", test.TEST_PREFIX, test.CLEAN_TEST_UUID), id.Id())
}

func TestIdentity_Id(t *testing.T) {
	// Given
	testPrefix := test.TEST_PREFIX
	// ApplyEvent
	ID, err := NewIdentity(testPrefix)
	assert.NoError(t, err)
	assert.NotNil(t, ID)
	id := ID.Id()
	// Then
	assert.True(t, strings.HasPrefix(id, testPrefix))
}

func TestImplementsIIdentity(t *testing.T) {
	// Given
	testId, err := NewIdentity(test.TEST_PREFIX)
	assert.NoError(t, err)
	// ApplyEvent
	b := ImplementsIIdentity(testId)
	// Then
	assert.True(t, b)
}

func TestThatWeCanCreateANilIdentity(t *testing.T) {
	// GIVEN
	// WHEN
	nilID, err := NilIdentity()
	assert.NoError(t, err)
	// THEN
	assert.NotNil(t, nilID)

}
