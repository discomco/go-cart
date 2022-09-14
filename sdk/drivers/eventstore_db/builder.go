package eventstore_db

import (
	"github.com/discomco/go-cart/sdk/config"
	"github.com/discomco/go-cart/sdk/core/builder"
	"github.com/discomco/go-cart/sdk/core/ioc"
)

func EventSourcing(cfgPath config.Path) ioc.IDig {
	ioc := builder.InjectCoLoMed(string(cfgPath))
	return ioc.Inject(ioc,
		SingletonESClient,
		EventStore,
		BehaviorStore)
}

func AddESProjector(ioc ioc.IDig) ioc.IDig {
	return ioc.Inject(ioc,
		SingletonESClient,
		EvtProjFtor)
}
