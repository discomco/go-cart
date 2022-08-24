package bogus

import (
	"github.com/discomco/go-cart/core"
)

const (
	BOGUS_PREFIX     = "bogus"
	START_EVT_TOPIC  = "bogus:started"
	START_CMD_TOPIC  = "bogus:start"
	START_FACT_TOPIC = "bogus.started"
	START_HOPE_TOPIC = "bogus.start"
)

type Root struct {
	ID     *core.Identity
	Status Status
	Car    *Car
}

func (m *Root) GetStatus() int {
	return int(m.Status)
}

func NewRootIdentity() (*core.Identity, error) {
	return core.NewIdentity(BOGUS_PREFIX)
}

func NewRoot(ID *core.Identity) *Root {
	return &Root{
		ID:     ID,
		Status: Unknown,
	}
}
