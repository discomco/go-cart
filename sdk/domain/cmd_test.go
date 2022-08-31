package domain

import "github.com/discomco/go-cart/sdk/core"

// ITestCmd is the injector for our test domain ICmd
type ITestCmd interface {
	ICmd
}

func newTestCmd(aggregateID *core.Identity) (ITestCmd, error) {
	return NewCmd(aggregateID, A_CMD_TOPIC, nil)
}
