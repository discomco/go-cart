package comps

import (
	"context"
	"fmt"
	"github.com/discomco/go-cart/examples/robby/execute-game/behavior/builder"
	"github.com/discomco/go-cart/examples/robby/execute-game/behavior/ftor"
	"github.com/discomco/go-cart/examples/robby/execute-game/schema"
	"github.com/discomco/go-cart/examples/robby/execute-game/spokes/initialize_game/behavior"
	"github.com/discomco/go-cart/examples/robby/execute-game/spokes/initialize_game/contract"
	sdk_behavior "github.com/discomco/go-cart/sdk/behavior"
	"github.com/discomco/go-cart/sdk/comps"
	"github.com/discomco/go-cart/sdk/container"
	contract2 "github.com/discomco/go-cart/sdk/contract"
	"github.com/discomco/go-cart/sdk/core/ioc"
	"github.com/discomco/go-cart/sdk/core/logger"
	"log"
	"math/rand"
	"time"
)

const (
	ConfigPath = "../../../config/config.yaml"
)

var (
	testEnv           ioc.IDig
	testLogger        logger.IAppLogger
	newTestBehavior   sdk_behavior.BehaviorBuilder
	newTestCmdHandler comps.CmdHandlerFtor
	testRequester     IRequester
	names             = []string{"John", "Paul", "George", "Ringo", "Brian", "Mick", "Charlie", "Linda", "Yoko", "Pattie", "Barbara"}
	games             = []string{"Chess", "Checkers", "Risk", "Monopoly", "Bridge", "Poker", "Whist", "Roulette", "Blackjack", "Go"}
	colors            = []string{"Red", "Orange", "Yellow", "Green", "Blue", "Indigo", "Violet"}
	adjectives        = []string{"awesome", "dreaded", "sweet", "bitter", "sour", "salty", "overwhelming", "boring"}
	happenings        = []string{"game", "marathon", "bonanza", "freak-out", "festival", "sit-in", "morning", "day", "night", "evening"}
)

func init() {
	testEnv = buildTestEnv()
}

func buildTestEnv() ioc.IDig {
	dig := container.DefaultCMD(ConfigPath)
	dig.Inject(dig,
		schema.DocFtor,
	).Inject(dig,
		ftor.BehaviorFtor,
		builder.BehaviorBuilder,
	).Inject(dig,
		Responder,
		behavior.Hope2Cmd,
		Requester,
	)
	return resolveTestEnv(dig)
}

func resolveTestEnv(dig ioc.IDig) ioc.IDig {
	err := dig.Invoke(func(
		appLogger logger.IAppLogger,
		newBehavior sdk_behavior.BehaviorBuilder,
		newCmdHandler comps.CmdHandlerFtor,
		requester IRequester,
	) {
		testLogger = appLogger
		newTestBehavior = newBehavior
		newTestCmdHandler = newCmdHandler
		testRequester = requester
	})
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	return dig
}

func randomGameName() string {
	iN1 := rand.Intn(len(names))
	iN2 := rand.Intn(len(names))
	iG1 := rand.Intn(len(games))
	iG2 := rand.Intn(len(games))
	iC1 := rand.Intn(len(colors))
	iA1 := rand.Intn(len(adjectives))
	iH1 := rand.Intn(len(happenings))
	return fmt.Sprintf(
		"%+v and %+v's %+v %+v & %+v,%+v %+v",
		names[iN1],
		names[iN2],
		colors[iC1],
		games[iG1],
		games[iG2],
		adjectives[iA1],
		happenings[iH1],
	)
}

func requestWorker(done <-chan bool, hopes <-chan contract.IHope, feedbacks chan contract2.IFbk, ctx context.Context) chan<- bool {
	terminated := make(chan<- bool)
	defer close(terminated)
	for {
		select {
		case <-done:
			return terminated
		default:
			hope := <-hopes
			fbk := testRequester.Request(ctx, hope, 10*time.Second)
			feedbacks <- fbk
		}
	}
	return terminated
}
