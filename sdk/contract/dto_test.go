package contract

import (
	"github.com/discomco/go-cart/sdk/schema"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestThatWeCanCreateADto(t *testing.T) {
	// GIVEN
	assert.True(t, testId != "")
	// WHEN
	d := newDto(testId)
	// THEN
	assert.NotNil(t, d)
}

func TestThatWeCanCreateAnIDto(t *testing.T) {
	// GIVEN
	assert.True(t, testId != "")
	// AND
	car := &MyPl{
		Name: "Yaris",
	}
	// WHEN
	d, err := NewDto(testId, car)
	// THEN
	assert.Nil(t, err)
	assert.NotNil(t, d)
	var ok bool
	switch IDto(d).(type) {
	case IDto:
		ok = true
	default:
		ok = false
	}
	assert.True(t, ok)
}

func TestThatWeCanGetTheAggregateIdFromADto(t *testing.T) {
	// GIVEN
	d, err := testDto()
	assert.Nil(t, err)
	assert.NotNil(t, d)
	// WHEN
	id := d.GetId()
	// THEN
	assert.Equal(t, testId, id)

}

func TestThatWeCanGetTheDataFromAdto(t *testing.T) {
	// GIVEN
	assert.NotNil(t, testID)
}

func TestThatWeCanSetTheDtoDataFromAnObject(t *testing.T) {
	// GIVEN
	d, err := testDto()
	assert.NotNil(t, d)
	assert.Nil(t, err)
	//AND
	newCar := &MyPl{Name: "Yaris"}
	// WHEN
	err = d.SetPayload(newCar)
	// THEN
	assert.Nil(t, err)
}

func TestThatWeCanGetTheDtoDataAsAnObject(t *testing.T) {
	// GIVEN
	d, err := testDto()
	assert.NotNil(t, d)
	assert.Nil(t, err)
	// WHEN
	var c *MyPl
	err = d.GetPayload(&c)
	assert.Nil(t, err)
	// THEN
	assert.NotNil(t, c)
	assert.Equal(t, "Hello", c.Name)
}

func TestThatWeCanGetTheDtoAggregateIIdentity(t *testing.T) {
	// GIVEN
	d, err := testDto()
	assert.NotNil(t, d)
	assert.Nil(t, err)
	//WHEN
	ID, err := d.GetID()
	assert.NoError(t, err)
	//THEN
	assert.NotNil(t, ID)
	ty := reflect.TypeOf(ID).Elem()
	typeName := ty.Name()
	assert.Equal(t, "Identity", typeName)
}

func TestThatWeCanInitializeADtoWithNilPayload(t *testing.T) {
	// GIVEN
	// WHEN
	ID, err := schema.NilIdentity()
	assert.NoError(t, err)
	d, err := NewDto(ID.Id(), nil)
	assert.Nil(t, err)
	assert.NotNil(t, d)
}
