package features

import (
	"github.com/discomco/go-cart/sdk/domain"
	"golang.org/x/net/context"
)

type IFeature interface {
	IComponent
	IShutdown
	Run(ctx context.Context) func() error
	Inject(plugins ...IFeaturePlugin) IFeature
}

type IShutdown interface {
	Shutdown(ctx context.Context)
}

type IPrjFeature interface {
	IFeature
}

type ICmdFeature interface {
	IMediatorFeature
}

type IMediatorFeature interface {
	IFeature
}

// IGenCmdFeature is a generic CMD ScreamingApp Feature,
// discriminated by TState of type ICmd
type IGenCmdFeature[T domain.ICmd] interface {
	ICmdFeature
}
