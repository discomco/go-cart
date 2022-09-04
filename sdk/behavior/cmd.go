package behavior

import (
	"encoding/json"
	"github.com/discomco/go-cart/sdk/contract"
	"github.com/discomco/go-cart/sdk/schema"
)

type Fact2CmdFunc[TFact contract.IFact, TCmd ICmd] func(fact *contract.Dto) (TCmd, error)
type Hope2CmdFunc[THope contract.IHope, TCmd ICmd] func(hope *contract.Dto) (TCmd, error)
type GenData2CmdFunc[TCmd ICmd] func([]byte) (TCmd, error)

// ICmd is an injector that represents a behavior.Cmd
type ICmd interface {
	CmdTypeGetter
	GetAggregateID() schema.IIdentity
	GetPayload() []byte                      //
	GetJsonPayload(pl schema.IPayload) error //
	SetJsonPayload(pl schema.IPayload) error
}

type Cmd struct {
	aggregateID schema.IIdentity
	payload     []byte
	commandType CommandType
}

//SetJsonPayload accepts the payload and serializes it to the payload []byte
func (c *Cmd) SetJsonPayload(pl schema.IPayload) error {
	d, err := json.Marshal(pl)
	if err != nil {
		return err
	}
	c.payload = d
	return nil
}

// GetJsonPayload serializes the payload []byte and returns a reference to an IPayload struct.
func (c *Cmd) GetJsonPayload(pl schema.IPayload) error {
	err := json.Unmarshal(c.payload, pl)
	if err != nil {
		return err
	}
	return nil
}

func (c *Cmd) GetCommandType() CommandType {
	return c.commandType
}

func (c *Cmd) GetPayload() []byte {
	return c.payload
}

func NewCmd(aggregateID schema.IIdentity, commandType CommandType, payload []byte) (*Cmd, error) {
	if aggregateID == nil {
		return nil, ErrAggregateIDCannotBeNil
	}
	if commandType == "" {
		return nil, ErrInvalidCommandType
	}
	return &Cmd{
		aggregateID: aggregateID,
		commandType: commandType,
		payload:     payload,
	}, nil
}

func (c *Cmd) GetAggregateID() schema.IIdentity {
	return c.aggregateID
}
