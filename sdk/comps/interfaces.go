package comps

import (
	"github.com/discomco/go-cart/sdk/behavior"
	"github.com/discomco/go-cart/sdk/config"
	"github.com/discomco/go-cart/sdk/contract"
	"github.com/discomco/go-cart/sdk/core/logger"
	"github.com/discomco/go-cart/sdk/core/mediator"
	"github.com/discomco/go-cart/sdk/schema"
	"golang.org/x/net/context"
	"time"
)

type EventStoreFtor func() IEventStore
type BehaviorStoreFtor func() IBehaviorStore
type ProjectorFtor func() IProjector
type ProjectorBuilder func(newProj ProjectorFtor) IProjector
type GenResponderFtor[THope contract.IHope] func() IGenResponder[THope]
type IGenResponder[THope contract.IHope] interface {
	IResponder
}

type IComponent interface {
	GetMediator() mediator.IMediator
	GetName() schema.Name
	GetLogger() logger.IAppLogger
	GetConfig() config.IAppConfig
}

type CommandHandler interface {
	Handle(ctx context.Context, cmd behavior.ICmd) contract.IFbk
}

type GenCommandHandler[T behavior.ICmd] interface {
	Handle(ctx context.Context, cmd T) contract.IFbk
}

// ICmdHandler is an interface to a Command Handler. Will be replaced with IGenCmdHandler
type ICmdHandler interface {
	IComponent
	CommandHandler
}

// IGenCmdHandler is a Command Handler for a specific ICmd
type IGenCmdHandler[T behavior.ICmd] interface {
	GenCommandHandler[T]
	GetTopic() behavior.Topic
	GetBehaviorStore() IBehaviorStore
	GetBehavior(ID schema.IIdentity) behavior.IBehavior
	SetTopic(topic behavior.Topic)
}

type IClose interface {
	Close() error
}

type IGenBus[TConn interface{}, TMsg interface{}] interface {
	IBus
	IConnection[TConn]
	Respond(ctx context.Context, topic string, hopes chan TMsg)
	RespondAsync(ctx context.Context, topic string, hopes chan TMsg) func() error
}

//BusFtor is a functor that returns a simple IBus injector.
type BusFtor func() (IBus, error)

//GenBusFtor is a generic functor that is discriminated by the Type of Connection and Message Type of the bus driver. that returns a IGenBus injector.
type GenBusFtor[TConn interface{}, TMsg interface{}] func() (IGenBus[TConn, TMsg], error)

// IBus is an interface to a Bus. Will be replaced with IGenBus
type IBus interface {
	Close()
	Publish(ctx context.Context, topic string, data []byte) error
	Request(ctx context.Context, topic string, data []byte, timeout time.Duration) ([]byte, error)
	RequestAsync(ctx context.Context, topic string, data []byte, timeout time.Duration, responses chan []byte) func() error
	Listen(ctx context.Context, topic string, facts chan []byte)
	ListenAsync(ctx context.Context, topic string, facts chan []byte) func() error
	Wait() error
}

type IConnection[TConn interface{}] interface {
	Connection() TConn
}

type IStore interface {
	Load(id string) interface{}
	Save(id string, model interface{})
}

type IBehaviorStore interface {
	IClose
	Load(ctx context.Context, behavior behavior.IBehavior) error
	Save(ctx context.Context, behavior behavior.IBehavior) error
	Exists(ctx context.Context, streamID string) error
}

type IEventStore interface {
	SaveEvents(ctx context.Context, streamID string, events []behavior.IEvt) error
	LoadEvents(ctx context.Context, streamID string) ([]behavior.IEvt, error)
}

type ISnapshotStore interface {
	SaveSnapshot(ctx context.Context, aggregate behavior.IBehavior) error
	GetSnapshot(ctx context.Context, id string) (*behavior.Snapshot, error)
}

type IActivate interface {
	Activate(ctx context.Context) error
}

type IDeactivate interface {
	Deactivate(ctx context.Context) error
}

type IBehaviorLink interface {
	IMediatorReaction
	IAmBehaviorLink()
}

type IGenMediatorReaction[TEvt behavior.IEvt] interface {
	ISpokePlugin
	behavior.EvtTypeGetter
	behavior.IGenReacter[TEvt]
}

//IMediatorReaction is an Injector for a mediator Subscriber.
//Will be replaced with IGenMediatorReaction at some point.
type IMediatorReaction interface {
	ISpokePlugin
	behavior.EvtTypeGetter
	behavior.Reacter
}

type IGenEvtReaction[TEvt behavior.IEvt] interface {
	IGenMediatorReaction[TEvt]
}

type IEvtReaction interface {
	IMediatorReaction
}

type IProjector interface {
	ISpokePlugin
	behavior.Reacter
	Project(ctx context.Context, prefixes []string, poolSize int) error
	Inject(handlers ...IProjection)
}

type IResponder interface {
	ISpokePlugin
	IAmResponder()
	GetHopeType() contract.HopeType
}

//IGenRequester is an Injector to a Hope Request Handler.
type IGenRequester[THope contract.IHope] interface {
	IRequester
	GenRequest(ctx context.Context, hope THope, timeout time.Duration) contract.IFbk
	GenRequestAsync(ctx context.Context, hope THope, timeout time.Duration) contract.IFbk
}

type IRequester interface {
	IComponent
	IAmRequester()
	GetHopeType() contract.HopeType
	Request(ctx context.Context, hope contract.IHope, timeout time.Duration) contract.IFbk
	RequestAsync(ctx context.Context, hope contract.IHope, timeout time.Duration) contract.IFbk
}

type RequesterFtor func() (IRequester, error)
type GenRequesterFtor[THope contract.IHope] func() (IGenRequester[THope], error)

//ISpokePlugin is a base Injector for Spoke plugins
type ISpokePlugin interface {
	IComponent
	IActivate
	IDeactivate
}

//IListener is an injector for all components that listen for Facts on a
//message bus.
type IListener interface {
	ISpokePlugin
	IAmFactListener()
}

type IGenListener[TMsg interface{}, TFact contract.IFact] interface {
	IListener
}

// IEmitter is the injector for components that emit facts to message brokers.
// It specializes the IMediatorReaction as it registers at the mediator,
// where it listens for specific events that must be emitted from the domain to other systems.
type IEmitter interface {
	ISpokePlugin
	IMediatorReaction
	IAmEmitter()
}

type IShutdown interface {
	Shutdown(ctx context.Context)
}

type IQueryProvider interface {
	ISpokePlugin
	IAmQueryProvider()
	RunQuery(ctx context.Context, qry contract.IReq) contract.IRsp
}
