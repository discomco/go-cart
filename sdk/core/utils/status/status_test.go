package status

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThatICanBeBorn(t *testing.T) {
	// GIVEN
	me = Unknown
	assert.True(t, HasFlag(me, Unknown))
	// WHEN
	ForceFlag(&me, Born)
	// THEN
	assert.True(t, HasFlag(me, Born))
	assert.True(t, NotHasFlag(me, Unknown))
}

func TestThatICanDie(t *testing.T) {
	// GIVEN
	me = Unknown
	assert.True(t, HasFlag(me, Unknown))
	// WHEN
	SetFlag(&me, Born)
	assert.True(t, HasFlag(me, Born))
	// AND
	SetFlag(&me, Living)
	assert.True(t, HasFlag(me, Living))
	// AND
	UnsetFlag(&me, Born)
	assert.True(t, NotHasFlag(me, Born))
	// AND
	ForceFlag(&me, Dead)
	// THEN
	assert.True(t, NotHasFlag(me, Born))
	assert.True(t, NotHasFlag(me, Living))
	assert.True(t, HasFlag(me, Dead))
}

func TestWeCanSetAndUnsetMultipleFlagsAtOnce(t *testing.T) {
	// GIVEN
	me = Unknown
	assert.True(t, HasFlag(me, Unknown))
	// WHEN
	SetFlags(&me, Born, Living, Dead)
	// THEN
	assert.True(t, HasFlags(me, Born, Living, Dead))
	// AND WHEN
	UnsetFlags(&me, Born, Dead)
	assert.True(t, !HasFlags(me, Born, Dead))

}
