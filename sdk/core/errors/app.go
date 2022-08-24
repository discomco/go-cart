package errors

import (
	"fmt"
	"github.com/pkg/errors"
)

const (
	NoLogger                = "no logger"
	NoMediator              = "no mediator"
	NoFeature               = "no feature"
	NoBus                   = "no bus"
	NoAggregateStore        = "no domain store"
	NoElementFor            = "no element for [%+v]"
	NoAggregate             = "no domain"
	NoBusConstructor        = "no bus constructor"
	NoCmdHandlerConstructor = "no command handler constructor"
	NoHope2CmdFunc          = "no hope to command converter"
	NoData2HopeFunc         = "no data to hope converter"
)

var (
	ErrNoLogger                = errors.New(NoLogger)
	ErrNoMediator              = errors.New(NoMediator)
	ErrNoFeature               = errors.New(NoFeature)
	ErrNoBus                   = errors.New(NoBus)
	ErrNoAggregateStore        = errors.New(NoAggregateStore)
	ErrNoAggregate             = errors.New(NoAggregate)
	ErrNoBusConstructor        = errors.New(NoBusConstructor)
	ErrNoCmdHandlerConstructor = errors.New(NoCmdHandlerConstructor)
	ErrNoHope2CmdFunc          = errors.New(NoHope2CmdFunc)
	ErrNoData2HopeFunc         = errors.New(NoData2HopeFunc)
)

func ErrNoElementFor(name string) error {
	return fmt.Errorf(NoElementFor, name)
}

func ErrCfgValidate(error error) error {
	return errors.Wrap(error, "Config Validation")
}
