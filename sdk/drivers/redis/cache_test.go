package redis

import (
	"github.com/discomco/go-cart/config"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"testing"
)

type MyModel struct {
	Name   string
	Brand  string
	Year   int
	Color  string
	Status int
}

func myTestModel() *MyModel {
	return &MyModel{
		Name:   "Yaris",
		Brand:  "Toyota",
		Year:   2013,
		Color:  "Pink",
		Status: 42,
	}
}

func (m MyModel) GetStatus() int {
	return m.Status
}

func Test_NewCache(t *testing.T) {
	// GIVEN
	cfg, err := config.AppConfig("../../config/config.yaml")
	assert.Nil(t, err)
	assert.NotNil(t, cfg)
	// WHEN
	c, err := newRedis[MyModel](cfg)
	// THEN
	assert.Nil(t, err)
	assert.NotNil(t, c)
}

func TestThatWeCanDeleteAKey(t *testing.T) {
	// GIVEN
	cfg, err := config.AppConfig("../../config/config.yaml")
	assert.Nil(t, err)
	assert.NotNil(t, cfg)
	// AND
	c, err := newRedis[MyModel](cfg)
	assert.Nil(t, err)
	assert.NotNil(t, c)
	// AND
	key := "test1"
	// AND
	ctx := context.Background()
	// AND
	newModel := myTestModel()
	// AND
	r, err := c.Set(ctx, key, *newModel)
	assert.Nil(t, err)
	assert.NotNil(t, r)
	// AND
	model, err := c.Get(ctx, key)
	assert.Nil(t, err)
	assert.NotNil(t, model)
	// WHEN
	oldModel, err := c.Delete(ctx, key)
	// THEN
	assert.Nil(t, err)
	assert.Equal(t, model, oldModel)
}

func TestThatWeCanWriteAModel(t *testing.T) {
	// GIVEN
	cfg, err := config.AppConfig("../../config/config.yaml")
	assert.Nil(t, err)
	assert.NotNil(t, cfg)
	c, err := newRedis[MyModel](cfg)
	assert.Nil(t, err)
	assert.NotNil(t, c)
	// AND
	m := myTestModel()
	// AND
	key := "test1"
	// AND
	ctx := context.Background()
	// WHEN
	r, err := c.Set(ctx, key, *m)
	// THEN
	assert.Nil(t, err)
	assert.NotNil(t, r)
}

func TestThatWeCanRetrieveAModel(t *testing.T) {
	// GIVEN
	cfg, err := config.AppConfig("../../config/config.yaml")
	assert.Nil(t, err)
	assert.NotNil(t, cfg)
	c, err := newRedis[MyModel](cfg)
	assert.Nil(t, err)
	assert.NotNil(t, c)
	// AND
	ctx := context.Background()
	// AND
	key := "test1"
	// AND
	m := myTestModel()
	// AND
	r, err := c.Set(ctx, key, *m)
	assert.Nil(t, err)
	assert.NotNil(t, r)
	// WHEN
	stored, err := c.Get(ctx, key)
	// THEN
	assert.Nil(t, err)
	assert.NotNil(t, stored)
	assert.Equal(t, m, stored)
	// AND
	deleted, err := c.Delete(ctx, key)
	assert.Nil(t, err)
	assert.Equal(t, m, deleted)
}
