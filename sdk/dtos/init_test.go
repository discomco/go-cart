package dtos

import (
	"github.com/discomco/go-cart/sdk/core"
	"github.com/discomco/go-cart/sdk/test"
)

var (
	testID core.IIdentity
	testId string
)

func init() {
	testID, err := core.NewIdentityFrom(test.TEST_PREFIX, test.TEST_UUID)
	if err != nil {
		panic(err)
	}
	testId = testID.Id()
}

func testDto() (IDto, error) {
	pl := &MyPl{Name: "Hello"}
	// WHEN
	return NewDto(testId, pl)
}

type MyPl struct {
	Name string
}
