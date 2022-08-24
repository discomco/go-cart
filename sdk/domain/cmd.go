package domain

import (
	"encoding/json"

	"github.com/discomco/go-cart/core"
	"github.com/discomco/go-cart/dtos"
	"github.com/discomco/go-cart/model"
)

type Fact2CmdFunc[TFact dtos.IFact, TCmd ICmd] func(fact *dtos.Dto) TCmd
type Hope2CmdFunc[THope dtos.IHope, TCmd ICmd] func(hope *dtos.Dto) (TCmd, error)
type GenData2CmdFunc[TCmd ICmd] func([]byte) (TCmd, error)

// ICmd is an injector that represents a domain.Cmd
type ICmd interface {
	CmdTypeGetter
	GetAggregateID() core.IIdentity
	GetPayload() []byte                     //
	GetJsonPayload(pl model.IPayload) error //
	SetJsonPayload(pl model.IPayload) error
}

type Cmd struct {
	aggregateID core.IIdentity
	payload     []byte
	commandType CommandType
}

//SetJsonPayload accepts the payload and serializes it to the payload []byte
func (c *Cmd) SetJsonPayload(pl model.IPayload) error {
	d, err := json.Marshal(pl)
	if err != nil {
		return err
	}
	c.payload = d
	return nil
}

// GetJsonPayload serializes the payload []byte and returns a reference to an IPayload struct.
func (c *Cmd) GetJsonPayload(pl model.IPayload) error {
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

func NewCmd(aggregateID core.IIdentity, commandType CommandType, payload []byte) (*Cmd, error) {
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

func (c *Cmd) GetAggregateID() core.IIdentity {
	return c.aggregateID
}
