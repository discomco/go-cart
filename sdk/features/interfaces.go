package features

import (
	"github.com/discomco/go-cart/sdk/config"
	"github.com/discomco/go-cart/sdk/core"
	"github.com/discomco/go-cart/sdk/core/logger"
	"github.com/discomco/go-cart/sdk/core/mediator"
	"github.com/discomco/go-cart/sdk/domain"
	"github.com/discomco/go-cart/sdk/dtos"
	"golang.org/x/net/context"
	"time"
)

type ESFtor func() IEventStore
type ASFtor func() IAggregateStore
type EventProjectorFtor func() IEventProjector
type EventProjectorBuilder func(newProj EventProjectorFtor) IEventProjector

type IComponent interface {
	GetMediator() mediator.IMediator
	GetName() Name
	GetLogger() logger.IAppLogger
	GetConfig() config.IAppConfig
}

type CommandHandler interface {
	Handle(ctx context.Context, cmd domain.ICmd) dtos.IFbk
}

type GenCommandHandler[T domain.ICmd] interface {
	Handle(ctx context.Context, cmd T) dtos.IFbk
}

// ICmdHandler is an interface to a Command Handler. Will be replaced with IGenCmdHandler
type ICmdHandler interface {
	IComponent
	CommandHandler
}

// IGenCmdHandler is a Command Handler for a specific ICmd
type IGenCmdHandler[T domain.ICmd] interface {
	GenCommandHandler[T]
	GetTopic() domain.Topic
	GetAggregateStore() IAggregateStore
	GetAggregate(ID core.IIdentity) domain.IAggregate
	SetTopic(topic domain.Topic)
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

type IAggregateStore interface {
	IClose
	Load(ctx context.Context, aggregate domain.IAggregate) error
	Save(ctx context.Context, aggregate domain.IAggregate) error
	Exists(ctx context.Context, streamID string) error
}

type IEventStore interface {
	SaveEvents(ctx context.Context, streamID string, events []domain.IEvt) error
	LoadEvents(ctx context.Context, streamID string) ([]domain.IEvt, error)
}

type ISnapshotStore interface {
	SaveSnapshot(ctx context.Context, aggregate domain.IAggregate) error
	GetSnapshot(ctx context.Context, id string) (*domain.Snapshot, error)
}

type IActivate interface {
	Activate(ctx context.Context) error
}

type IDeactivate interface {
	Deactivate(ctx context.Context) error
}

type IDomainLink interface {
	IMediatorSubscriber
	IAmDomainLink()
}

type IGenMediatorSubscriber[TEvt domain.IEvt] interface {
	IFeaturePlugin
	domain.EvtTypeGetter
	domain.IGenWhen[TEvt]
}

//IMediatorSubscriber is an Injector for a mediator Subscriber.
//Will be replaced with IGenMediatorSubscriber at some point.
type IMediatorSubscriber interface {
	IFeaturePlugin
	domain.EvtTypeGetter
	domain.Whener
}

type IGenEvtHandler[TEvt domain.IEvt] interface {
	IGenMediatorSubscriber[TEvt]
}

type IEvtHandler interface {
	IMediatorSubscriber
}

type IEventProjector interface {
	IFeaturePlugin
	domain.Whener
	Project(ctx context.Context, prefixes []string, poolSize int) error
	Inject(handlers ...IProjection)
}

type IHopeResponder interface {
	IFeaturePlugin
	IAmHopeResponder()
	GetHopeType() dtos.HopeType
}

//IHopeReqHandler is an Injector to a Hope Request Handler.
type IGenHopeRequester[THope dtos.IHope] interface {
	IHopeRequester
	GenRequest(ctx context.Context, hope THope, timeout time.Duration) dtos.IFbk
	GenRequestAsync(ctx context.Context, hope THope, timeout time.Duration) dtos.IFbk
}

type IHopeRequester interface {
	IComponent
	IAmHopeRequester()
	GetHopeType() dtos.HopeType
	Request(ctx context.Context, hope dtos.IHope, timeout time.Duration) dtos.IFbk
	RequestAsync(ctx context.Context, hope dtos.IHope, timeout time.Duration) dtos.IFbk
}

type HopeRequesterFtor func() (IHopeRequester, error)
type GenHopeRequesterFtor[THope dtos.IHope] func() (IGenHopeRequester[THope], error)

//IFeaturePlugin is a base Injector for Feature plugins
type IFeaturePlugin interface {
	IComponent
	IActivate
	IDeactivate
}

//IFactListener is an injector for all components that listen for Facts on a
//message bus.
type IFactListener interface {
	IFeaturePlugin
	IAmFactListener()
}

type IGenFactListener[TMsg interface{}, TFact dtos.IFact] interface {
	IFactListener
}

// IFactEmitter is the injector for components that emit facts to message brokers.
// It specializes the IMediatorSubscriber as it registers at the mediator,
// where it listens for specific events that must be emitted from the domain to other systems.
type IFactEmitter interface {
	IFeaturePlugin
	IMediatorSubscriber
	IAmEmitter()
}

//IApp is the Generic Injector for GO-SCREAM applications
type IApp interface {
	IComponent
	IShutdown
	Run() error
	Inject(features ...IFeature) IApp
}

type AppFtor func() IApp
type AppBuilder func() IApp
