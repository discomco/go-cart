package behavior

import (
	"fmt"
	"github.com/pkg/errors"
)

var (
	ErrCommandCannotBeNil        = errors.New(CommandCannotBeNil)
	ErrPayloadMustNotBeEmpty     = errors.New(PayloadMustNotBeEmpty)
	ErrCommandMustHaveBehaviorID = errors.New(CommandMustHaveBehaviorID)
	ErrFailedToGetJsonData       = errors.New(FailedToGetJsonData)
	ErrCommandTypeMustNotBeEmpty = errors.New(CommandTypeMustNotBeEmpty)
	ErrAlreadyExists             = errors.New(AlreadyExists)
	ErrBehaviorNotFound          = errors.New(BehaviorNotFound)
	ErrInvalidEventType          = errors.New(InvalidEventType)
	ErrInvalidCommandType        = errors.New(InvalidCommandType)
	ErrInvalidBehavior           = errors.New(InvalidBehavior)
	ErrInvalidBehaviorID         = errors.New(InvalidBehaviorID)
	ErrInvalidEventVersion       = errors.New(InvalidEventVersion)
	ErrEventHasNoBehaviorID      = errors.New(EventHasNoBehaviorID)
	ErrInvalidBehaviorType       = errors.New(InvalidBehaviorType)
	ErrHopeCannotBeNil           = errors.New(HopeCannotBeNil)
	ErrNoTopic                   = errors.New(NoTopic)
	ErrTheBehaviorHasNoID        = errors.New(TheBehaviorHasNoIDPleaseUseSetID)
	ErrBehaviorIDCannotBeNil     = errors.New(BehaviorIDCannotBeNil)
)

const (
	NoTopic                          = "No topic"
	InvalidEventVersion              = "invalid event version"
	InvalidBehaviorID                = "invalid behavior id"
	InvalidBehavior                  = "invalid behavior"
	InvalidBehaviorType              = "invalid behavior type"
	InvalidCommandType               = "invalid command type"
	InvalidEventType                 = "invalid event type"
	BehaviorNotFound                 = "behavior not found"
	AlreadyExists                    = "already exists"
	PayloadMustNotBeEmpty            = "payload cmd_must not be empty"
	CommandCannotBeNil               = "command cannot be nil"
	CommandMustHaveBehaviorID        = "command cmd_must have behavior id"
	FailedToGetJsonData              = "failed to get json data"
	CommandTypeMustNotBeEmpty        = "command type cmd_must not be empty"
	EventHasNoBehaviorID             = "event has no behavior id"
	NoApplierForEvent                = "no F(apply) for event [%+v]"
	NoExecuterForCommand             = "no F(try) for command [%+v]"
	HopeCannotBeNil                  = "Hope cannot be nil"
	ExecuteDidNotReturnAnEvent       = "[%+v].execute did not return an event"
	TheBehaviorHasNoIDPleaseUseSetID = "the behavior has no Id, please use SetID()"
	BehaviorIDCannotBeNil            = "behaviorID cannot be nil"
)

func ErrExecuteDidNotReturnAnEvent(commandType string) string {
	return fmt.Sprintf(ExecuteDidNotReturnAnEvent, commandType)
}
