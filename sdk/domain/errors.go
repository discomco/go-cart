package domain

import (
	"fmt"
	"github.com/pkg/errors"
)

var (
	ErrCommandCannotBeNil         = errors.New(CommandCannotBeNil)
	ErrPayloadMustNotBeEmpty      = errors.New(PayloadMustNotBeEmpty)
	ErrCommandMustHaveAggregateID = errors.New(CommandMustHaveAggregateID)
	ErrFailedToGetJsonData        = errors.New(FailedToGetJsonData)
	ErrCommandTypeMustNotBeEmpty  = errors.New(CommandTypeMustNotBeEmpty)
	ErrAlreadyExists              = errors.New(AlreadyExists)
	ErrAggregateNotFound          = errors.New(AggregateNotFound)
	ErrInvalidEventType           = errors.New(InvalidEventType)
	ErrInvalidCommandType         = errors.New(InvalidCommandType)
	ErrInvalidAggregate           = errors.New(InvalidAggregate)
	ErrInvalidAggregateID         = errors.New(InvalidAggregateID)
	ErrInvalidEventVersion        = errors.New(InvalidEventVersion)
	ErrEventHasNoAggregateID      = errors.New(EventHasNoAggregateID)
	ErrInvalidAggregateType       = errors.New(InvalidAggregateType)
	ErrHopeCannotBeNil            = errors.New(HopeCannotBeNil)
	ErrNoTopic                    = errors.New(NoTopic)
	ErrTheAggregateHasNoID        = errors.New(TheAggregateHasNoIDPleaseUseSetID)
	ErrAggregateIDCannotBeNil     = errors.New(AggregateIDCannotBeNil)
)

const (
	NoTopic                           = "No topic"
	InvalidEventVersion               = "invalid event version"
	InvalidAggregateID                = "invalid aggregate id"
	InvalidAggregate                  = "invalid aggregate"
	InvalidAggregateType              = "invalid aggregate type"
	InvalidCommandType                = "invalid command type"
	InvalidEventType                  = "invalid event type"
	AggregateNotFound                 = "domain not found"
	AlreadyExists                     = "already exists"
	PayloadMustNotBeEmpty             = "payload cmd_must not be empty"
	CommandCannotBeNil                = "command cannot be nil"
	CommandMustHaveAggregateID        = "command cmd_must have domain id"
	FailedToGetJsonData               = "failed to get json data"
	CommandTypeMustNotBeEmpty         = "command type cmd_must not be empty"
	EventHasNoAggregateID             = "event has no domain id"
	NoApplierForEvent                 = "no applier for event [%+v]"
	NoExecuterForCommand              = "no executer for command [%+v]"
	HopeCannotBeNil                   = "Hope cannot be nil"
	ExecuteDidNotReturnAnEvent        = "[%+v].execute did not return an event"
	TheAggregateHasNoIDPleaseUseSetID = "the aggregate has no Id, please use SetID()"
	AggregateIDCannotBeNil            = "aggregateID cannot be nil"
)

func ErrExecuteDidNotReturnAnEvent(commandType string) string {
	return fmt.Sprintf(ExecuteDidNotReturnAnEvent, commandType)
}
