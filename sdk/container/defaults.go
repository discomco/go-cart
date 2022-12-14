package container

import (
	"github.com/discomco/go-cart/sdk/comps"
	"github.com/discomco/go-cart/sdk/core/builder"
	"github.com/discomco/go-cart/sdk/core/ioc"
	"github.com/discomco/go-cart/sdk/drivers/eventstore_db"
	"github.com/discomco/go-cart/sdk/drivers/nats"
	"github.com/discomco/go-cart/sdk/drivers/tirol"
)

//DefaultCMD creates a basic Container that injects the infrastructure for
//
//- Configuration (IAppConfig)
//
//- Logging (IAppLogger
//
//- Mediator (IMediator)
//
//- EventStoreDB (for EventStore functionality) and
//
//- NATS (for async integration)
//
// MUST: This is the default drivers, for which the SDK provides built-in drivers.
// It serves as the basis for CMDApps and in order to make it a functioning application,
// you will need to provide Application-specific AggFtor and ConfigureEventAggBuilder injections as well as the relevant CmdFeatures.
//
// OPTIONAL: You may extend this container by injecting additional Backend Drivers according to your requirements.
func DefaultCMD(cfgPath string) ioc.IDig {
	dig := builder.InjectCoLoMed(cfgPath)
	return dig.Inject(dig,
		comps.CmdHandler,
	).Inject(dig,
		eventstore_db.SingletonESClient,
		eventstore_db.EventStore,
		eventstore_db.BehaviorStore,
	).Inject(dig,
		nats.SingleNATS,
	)
}

// DefaultPRJ creates a basic Container, discriminated by the type of the ReadModel that injects the infrastructure for
//
// - Configuration (IAppConfig)
//
// - Logging (IAppLogger)
//
// - Mediator (IMediator)
//
// - EventStoreDB (for EventStore functionality) and
//
// - Redis (for caching)
//
// You may extend this container with additional Backend Drivers according to your requirements.
func DefaultPRJ(cfgPath string) ioc.IDig {
	dig := builder.InjectCoLoMed(cfgPath)
	return dig.Inject(dig,
		eventstore_db.SingletonESClient,
		eventstore_db.EventStore,
		eventstore_db.BehaviorStore,
		eventstore_db.EvtProjFtor,
		eventstore_db.EventProjector,
	)
}

// DefaultQRY creates a basic Container that injects the infrastructure for
// - Configuration (IAppConfig)
//
// - Logging (IAppLogger)
//
// - Mediator (IMediator)
//
// - EventStoreDB (for EventStore functionality)
//
// - NATS (for async integration)
//
// - Redis (for caching)
//
// - ElasticSearch (for free text search)
//
// You may extend this container with additional Backend Drivers according to your requirements.
func DefaultQRY(cfgPath string) ioc.IDig {
	dig := builder.InjectCoLoMed(cfgPath)
	return dig.Inject(dig,
		tirol.NewTirol,
	)
}
