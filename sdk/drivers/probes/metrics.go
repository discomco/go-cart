package probes

import (
	"context"
	"fmt"

	"github.com/discomco/go-cart/config"
	"github.com/discomco/go-cart/features"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"golang.org/x/sync/errgroup"
)

const (
	stackSize = 1 << 10 // 1KB
)

type IMetrics interface {
	features.IFeature
	IAmMetrics()
	RegisterCounters(counters ...Counter)
}

type metrics struct {
	*features.AppComponent
	chDone   chan struct{}
	errGroup *errgroup.Group
	config   config.IProbesConfig
	counters map[features.Name]Counter
}

func (m *metrics) Shutdown(ctx context.Context) {
	m.chDone <- struct{}{}
	m.GetLogger().Infof("%+v has stopped.", m.GetName())
}

func (m *metrics) Run(ctx context.Context) func() error {
	m.GetLogger().Infof("%+v running on port %v", m.GetName(), m.config.GetPort())
	return func() error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-m.chDone:
			return ctx.Err()
		default:
			metricsServer := echo.New()
			metricsServer.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
				StackSize:         stackSize,
				DisablePrintStack: true,
				DisableStackAll:   true,
			}))
			metricsServer.GET(m.config.GetPrometheusPath(), echo.WrapHandler(promhttp.Handler()))
			if err := metricsServer.Start(m.config.GetPort()); err != nil {
				m.GetLogger().Errorf("metricsServer.Start: {%v}", err)
				m.chDone <- struct{}{}
				return err
			}
		}
		return nil
	}
}

func (m *metrics) Inject(plugins ...features.IFeaturePlugin) features.IFeature {
	for _, plugin := range plugins {
		switch plugin.(type) {
		case IMetricsCounter:
			m.RegisterCounters(plugin.(Counter))
		}
	}
	return m
}

func (m *metrics) IAmMetrics() {}

func (m *metrics) RegisterCounters(counters ...Counter) {
	for _, c := range counters {
		_, ok := m.counters[c.Name]
		if !ok {
			m.counters[c.Name] = c
		}
	}
}

const (
	NameFmt = "[%+v].Metrics"
)

func NewMetrics(config config.IAppConfig) IMetrics {
	name := fmt.Sprintf(NameFmt, config.GetServiceConfig().GetServiceName())
	return newMetrics(features.Name(name), config.GetProbesConfig())
}

func newMetrics(name features.Name, probesConfig config.IProbesConfig) IMetrics {
	b := features.NewAppComponent(name)
	m := &metrics{
		config:   probesConfig,
		counters: make(map[features.Name]Counter),
	}
	m.AppComponent = b
	return m
}

type IMetricsCounter interface {
	features.IFeaturePlugin
	IAmMetricsCounter()
}

func NewCounter(name features.Name,
	nameSpace string,
	subSystem string,
	help string,
	constLabels map[string]string) IMetricsCounter {
	opts := prometheus.CounterOpts{
		Namespace:   nameSpace,
		Subsystem:   subSystem,
		Name:        string(name),
		Help:        help,
		ConstLabels: constLabels,
	}
	return newCounter(name, promauto.NewCounter(opts))
}

func newCounter(name features.Name, pCounter prometheus.Counter) *Counter {
	comp := features.NewAppComponent(name)
	c := &Counter{
		PCounter: pCounter,
	}
	c.AppComponent = comp
	return c
}

type Counter struct {
	*features.AppComponent
	PCounter prometheus.Counter
}

func (c Counter) Activate(ctx context.Context) error {
	return nil
}

func (c Counter) Deactivate(ctx context.Context) error {
	return nil
}

func (c Counter) IAmMetricsCounter() {}
