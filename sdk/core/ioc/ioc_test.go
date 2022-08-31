package ioc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThatSingletonDigsAreIndeedTheSame(t *testing.T) {
	b := SingleIoC()
	c := SingleIoC()
	assert.NotNil(t, b)
	assert.NotNil(t, c)
	assert.Same(t, b, c)
}

type myStructure struct{}

func NewMyStruct() *myStructure { return &myStructure{} }

func TestShouldBeSameStruct(t *testing.T) {
	// Build The Container
	ioc := SingleIoC()
	ioc.Inject(ioc,
		NewMyStruct)

	var agg1 *myStructure
	err := ioc.Invoke(func(a *myStructure) {
		agg1 = a
	})
	assert.NoError(t, err)
	assert.NotNil(t, agg1)
	var agg2 *myStructure
	_ = ioc.Invoke(func(a *myStructure) {
		agg2 = a
	})
	assert.NotNil(t, agg2)
	assert.Same(t, agg2, agg1)
}

func TestThatWeCanNotResolveASeriesOfInjections(t *testing.T) {
	// GIVEN
	assert.NotNil(t, testEnv)
	// WHEN
	err := testEnv.Invoke(func(hc *HumansAndCars) {
		assert.NotNil(t, hc)
		assert.Empty(t, hc.list)
	})
	assert.NoError(t, err)

}
