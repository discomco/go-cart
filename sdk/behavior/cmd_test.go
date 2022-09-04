package behavior

import "github.com/discomco/go-cart/sdk/schema"

// ITestCmd is the injector for our test domain ICmd
type ITestCmd interface {
	ICmd
}

func newTestCmd(aggregateID *schema.Identity) (ITestCmd, error) {
	return NewCmd(aggregateID, A_CMD_TOPIC, nil)
}
