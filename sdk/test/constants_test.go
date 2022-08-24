package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTEST_TRACE_ID(t *testing.T) {
	assert.Equal(t, "test_trace_id", TEST_TRACE_ID)
}
