package container

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/discomco/go-cart/config"
	"github.com/discomco/go-cart/core/errors"
	"github.com/discomco/go-cart/drivers/jaeger"
	"github.com/discomco/go-cart/features"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
	"golang.org/x/sync/errgroup"
)

type (
	DownAppFunc func(ctx context.Context)
	RunAppFunc  func() error
)

const (
	waitShutdownDuration = 3
)

var (
	cMutex    = &sync.Mutex{}
	singleApp features.IApp
)

// App is the receiver that serves as the -base for all GO-SCREAM Apps
//
// Definition:  CPQRS: "Command GenProjection Query Segregation"
//   Extending on the principle of CQRS, we separate the responsibility of
//   Projecting Events from the EventStream to a separate application
//
// Definition: Screaming Architecture:
//   An architectural style in which all artifacts "scream" their purpose to the observer.
//   This means that there cmd_must be a clear alignment between discovery and implementation artifacts.
//
// Definition: GO-SCREAM Application:
//   a GO-SCREAM Applications is conceived as a Screaming Monolith (a form of Majestic Monolith),
//   where "Features" are the building blocks of the application.
//   Following the principles of Clean Architecture,  all side effect functionality is addressed via interfaces.
//
// There are 3 types of GO-SCREAM apps:
//   1) *-CMD apps that implement pure COMMAND (write) functionality (the "C" in CPQRS)
//   2) *-PRJ apps that implement pure PROJECTION (ETL) functionality. (the "P" in CPQRS)
//   3) *-QRY apps that implement pure QUERY (read) functionality (the "Q" in CPQRS)
//
type App struct {
	*features.AppComponent
	echo     *echo.Echo
	features map[features.Name]features.IFeature
	doneCh   chan struct{}
	runApp   RunAppFunc
	downApp  DownAppFunc
}

func (a *App) calcShutdownDuration() int {
	return waitShutdownDuration * (len(a.features) + 1)
}

func (a *App) waitShutdown(duration int) {
	go func(dur int) {
		d := time.Duration(dur) * 1000 * 1000 * 1000
		time.Sleep(d)
		a.doneCh <- struct{}{}
	}(duration)
}

func NewApp(
	config config.IAppConfig,
	run RunAppFunc,
	down DownAppFunc,
) *App {
	name := config.GetServiceConfig().GetServiceName()

	base := features.NewAppComponent(features.Name(name))
	a := &App{
		features: make(map[features.Name]features.IFeature, 0),
		runApp:   run,
		downApp:  down,
	}
	a.AppComponent = base
	return a
}

func (a *App) regFeature(feature features.IFeature) {
	if feature == nil {
		return
	}
	if a.features[feature.GetName()] == nil {
		a.features[feature.GetName()] = feature
	}
}

//Run the App
func (a *App) Run() error {
	defer func() {
		if err := recover(); err != nil {
			a.GetLogger().Errorf("panic occurred:", err)
		}
	}()
	// Create Context
	v := validator.New()
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer stop()
	if err := v.StructCtx(ctx, a.GetConfig()); err != nil {
		return errors.ErrCfgValidate(err)
	}
	// Setup Tracing
	cfg := a.GetConfig().(*config.Config)
	if cfg.Jaeger != nil {
		tracerCfg := a.GetConfig().GetJaegerConfig()
		if tracerCfg != nil {
			if tracerCfg.IsEnabled() {
				tracer, closer, err := jaeger.NewJaegerTracer(tracerCfg)
				if err != nil {
					return err
				}
				defer closer.Close()
				opentracing.SetGlobalTracer(tracer)
			}
		}
	}
	g, ctx := errgroup.WithContext(ctx)
	for _, f := range a.features {
		g.Go(f.Run(ctx))
	}
	if a.runApp != nil {
		g.Go(a.runApp)
	}
	a.GetLogger().Infof("Application [%+v] is RUNNING!", a.GetConfig().GetServiceConfig().GetServiceName())
	<-ctx.Done()
	a.waitShutdown(a.calcShutdownDuration())
	a.Shutdown(ctx)
	<-a.doneCh
	a.GetLogger().Infof("Application [%+v] EXITED PROPERLY!", a.GetConfig().GetServiceConfig().GetServiceName())
	return g.Wait()
}

//Shutdown shuts down the application
func (a *App) Shutdown(ctx context.Context) {
	for _, f := range a.features {
		f.Shutdown(ctx)
	}
	if a.downApp != nil {
		a.downApp(ctx)
	}
}

func (a *App) Inject(features ...features.IFeature) features.IApp {
	if len(features) == 0 {
		return a
	}
	for _, feature := range features {
		a.regFeature(feature)
	}
	return a
}
