package eventstore_db

import (
	"github.com/discomco/go-cart/config"
	"github.com/discomco/go-cart/core/builder"
	"github.com/discomco/go-cart/core/ioc"
)

func EventSourcing(cfgPath config.Path) ioc.IDig {
	ioc := builder.InjectCoLoMed(string(cfgPath))
	return ioc.Inject(ioc,
		SingletonESClient,
		EStore,
		AStore)
}

func AddESProjector(ioc ioc.IDig) ioc.IDig {
	return ioc.Inject(ioc,
		SingletonESClient,
		EvtProjFtor)
}
