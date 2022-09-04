package behavior

import (
	"context"
)

type IGetBehavior interface {
	GetBehavior() IBehavior
}

type ILoadEvents interface {
	Load(events []Event) error
}

type Reacter interface {
	React(ctx context.Context, evt IEvt) error
}

type GenReacter[TEvt IEvt] interface {
	React(ctx context.Context, evt TEvt) error
}

type IGenReacter[TEvt IEvt] interface {
	GenReacter[TEvt]
}
