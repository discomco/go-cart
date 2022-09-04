package spokes

import (
	"github.com/discomco/go-cart/sdk/behavior"
	"github.com/discomco/go-cart/sdk/comps"
	"golang.org/x/net/context"
)

type ISpoke interface {
	comps.IComponent
	comps.IShutdown
	Run(ctx context.Context) func() error
	Inject(plugins ...comps.IReactor) ISpoke
}

type IMediatorSpoke interface {
	ISpoke
}

type IPrjSpoke interface {
	ISpoke
}

type ICmdSpoke interface {
	IMediatorSpoke
}

type IQrySpoke interface {
	ISpoke
}

// IGenCmdSpoke is a generic CMD ScreamingApp Spoke,
// discriminated by T of type ICmd
type IGenCmdSpoke[T behavior.ICmd] interface {
	ICmdSpoke
}

//IApp is the Generic Injector for GO-CART applications
type IApp interface {
	comps.IComponent
	comps.IShutdown
	Run() error
	Inject(spokes ...ISpoke) IApp
}

type AppFtor func() IApp
type AppBuilder func() IApp
