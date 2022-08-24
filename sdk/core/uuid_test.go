package core

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNullId(t *testing.T) {
	// Given
	id := NullId()
	// ApplyEvent
	newId := strings.Replace(id, "0", "A", -1)
	// Then
	assert.Equal(t, "AAAAAAAA-AAAA-AAAA-AAAA-AAAAAAAAAAAA", newId)
}

func TestCleanUuid(t *testing.T) {
	// Given
	// ApplyEvent
	uid, _ := CleanUuid()
	// Then
	assert.True(t, !strings.Contains(uid, "-"))
	assert.Equal(t, 32, len(uid))
}

func TestNullUuid(t *testing.T) {
	// Given
	// ApplyEvent
	uid := CleanNullId()
	// Then
	assert.True(t, !strings.Contains(uid, "-"))
	assert.Equal(t, 32, len(uid))
}

func TestNewUuid(t *testing.T) {
	uid, _ := NewUuid()
	if !strings.Contains(uid, "-") {
		t.Errorf("%+v  does not contain '-'", uid)
	}
	if len(uid) != 36 {
		t.Errorf("%+v is not 36 characters long. Length is: %+v", uid, len(uid))
	}
}
