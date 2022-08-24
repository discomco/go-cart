package core

import (
	"testing"

	core_mocks "github.com/discomco/go-cart/test"
	"github.com/stretchr/testify/assert"
)

func TestNewMeta(t *testing.T) {
	id, err := NewIdentityFrom(core_mocks.TEST_PREFIX, core_mocks.TEST_UUID)
	assert.NoError(t, err)
	m := NewMeta(id, core_mocks.TEST_TRACE_ID, 42)
	assert.NotNil(t, m)
	assert.NotNil(t, m.ID)
	assert.Equal(t, id.Id(), m.ID.Id())
}
