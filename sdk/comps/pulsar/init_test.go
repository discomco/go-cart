package pulsar

import (
	"github.com/discomco/go-cart/sdk/comps"
	"github.com/discomco/go-cart/sdk/core/builder"
	"github.com/discomco/go-cart/sdk/core/ioc"
	"github.com/discomco/go-cart/sdk/core/logger"
	"golang.org/x/net/context"
	"log"
)

const (
	CfgPath = "../../config/config.yaml"
)

var (
	testEnv            ioc.IDig
	testSecondPulsar   ISecondPulsar
	testMediatorLogger comps.IMediatorLogger
	testLogger         logger.IAppLogger
)

func init() {
	testEnv = buildTestEnv()
	err := testEnv.Invoke(func(
		sp ISecondPulsar,
		ml comps.IMediatorLogger,
		al logger.IAppLogger,
	) {
		testMediatorLogger = ml
		testSecondPulsar = sp
		testLogger = al
	})
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}

func buildTestEnv() ioc.IDig {
	dig := builder.InjectCoLoMed(CfgPath)
	return dig.Inject(dig,
		SecondPulsar,
		comps.GeneralMediatorLogger,
	)
}

func runMediatorLogger(ctx context.Context) func() error {
	return func() error {
		return testMediatorLogger.Activate(ctx)
	}
}

func runPulsar(ctx context.Context) func() error {
	return func() error {
		return testSecondPulsar.Activate(ctx)
	}
}
