package features

import (
	"github.com/discomco/go-cart/sdk/behavior"
	"github.com/discomco/go-cart/sdk/reactors"
	"golang.org/x/net/context"
)

type ISpoke interface {
	reactors.IComponent
	reactors.IShutdown
	Run(ctx context.Context) func() error
	Inject(plugins ...reactors.IReactor) ISpoke
}

type IPrjSpoke interface {
	ISpoke
}

type ICmdSpoke interface {
	IMediatorSpoke
}

type IMediatorSpoke interface {
	ISpoke
}

// IGenCmdSpoke is a generic CMD ScreamingApp Spoke,
// discriminated by T of type ICmd
type IGenCmdSpoke[T behavior.ICmd] interface {
	ICmdSpoke
}

//IApp is the Generic Injector for GO-CART applications
type IApp interface {
	reactors.IComponent
	reactors.IShutdown
	Run() error
	Inject(spokes ...ISpoke) IApp
}

type AppFtor func() IApp
type AppBuilder func() IApp
