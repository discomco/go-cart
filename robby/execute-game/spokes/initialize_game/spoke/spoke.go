package spoke

import (
	"github.com/discomco/go-cart/robby/execute-game/-shared/behavior/builder"
	"github.com/discomco/go-cart/robby/execute-game/-shared/behavior/ftor"
	"github.com/discomco/go-cart/robby/execute-game/-shared/schema"
	"github.com/discomco/go-cart/robby/execute-game/spokes/initialize_game/actors"
	"github.com/discomco/go-cart/robby/execute-game/spokes/initialize_game/behavior"
	"github.com/discomco/go-cart/sdk/config"
	"github.com/discomco/go-cart/sdk/container"
	"github.com/discomco/go-cart/sdk/core/ioc"
	"github.com/discomco/go-cart/sdk/features"
	"log"
)

type ISpoke interface {
	features.ICmdFeature
}

func newSpoke() ISpoke {
	return features.NewCmdFeature(behavior.CMD_TOPIC)
}

func Spoke(cfgPath config.Path) ISpoke {
	dig := container.DefaultCMD(string(cfgPath))
	dig.Inject(dig,
		schema.RootFtor,
	).Inject(dig,
		ftor.BehaviorFtor,
		builder.BehaviorBuilder,
	).Inject(dig,
		actors.Responder,
		behavior.Hope2Cmd,
	)
	return resolve(dig)
}

func resolve(dig ioc.IDig) ISpoke {
	m := newSpoke()
	var responder actors.IResponder
	err := dig.Invoke(func(
		r actors.IResponder,
	) {
		responder = r
	})
	m.Inject(
		responder,
	)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	return m
}
