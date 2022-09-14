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
	Inject(reactions ...comps.ISpokePlugin) ISpoke
}

type IMediatorSpoke interface {
	ISpoke
}

type IProjectionSpoke interface {
	ISpoke
}

type ICommandSpoke interface {
	IMediatorSpoke
}

type IQuerySpoke interface {
	ISpoke
}

// IGenCmdSpoke is a generic CMD ScreamingApp Spoke,
// discriminated by T of type ICmd
type IGenCmdSpoke[T behavior.ICmd] interface {
	ICommandSpoke
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
