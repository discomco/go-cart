package guards

import (
	"testing"

	"github.com/nauyey/guard"
	"github.com/nauyey/guard/validators"
	"github.com/stretchr/testify/assert"
)

func TestGuard(t *testing.T) {
	// Given
	name := "User Name"
	age := 10
	gender := ""

	// ApplyEvent
	// validate data
	err := guard.Validate(
		&validators.StringNotBlank{Value: name},
		&validators.IntGreaterThan{Value: age, Target: 16},                                  // invalid age
		&validators.StringInclusion{Value: gender, In: []string{"female", "male", "other"}}, // invalid gender
	)
	errs, ok := err.(guard.Errors)

	// Then
	assert.Equal(t, 2, len(errs.ValidationErrors()))
	assert.True(t, ok)
}
