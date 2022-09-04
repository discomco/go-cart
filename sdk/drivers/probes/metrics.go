package probes

import (
	"context"
	"fmt"
	"github.com/discomco/go-cart/sdk/comps"
	"github.com/discomco/go-cart/sdk/config"
	"github.com/discomco/go-cart/sdk/features"
	"github.com/discomco/go-cart/sdk/schema"
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
	features.ISpoke
	IAmMetrics()
	RegisterCounters(counters ...Counter)
}

type metrics struct {
	*comps.Component
	chDone   chan struct{}
	errGroup *errgroup.Group
	config   config.IProbesConfig
	counters map[schema.Name]Counter
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

func (m *metrics) Inject(plugins ...comps.IReactor) features.ISpoke {
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
	return newMetrics(schema.Name(name), config.GetProbesConfig())
}

func newMetrics(name schema.Name, probesConfig config.IProbesConfig) IMetrics {
	b := comps.NewComponent(name)
	m := &metrics{
		config:   probesConfig,
		counters: make(map[schema.Name]Counter),
	}
	m.Component = b
	return m
}

type IMetricsCounter interface {
	comps.IReactor
	IAmMetricsCounter()
}

func NewCounter(name schema.Name,
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

func newCounter(name schema.Name, pCounter prometheus.Counter) *Counter {
	comp := comps.NewComponent(name)
	c := &Counter{
		PCounter: pCounter,
	}
	c.Component = comp
	return c
}

type Counter struct {
	*comps.Component
	PCounter prometheus.Counter
}

func (c Counter) Activate(ctx context.Context) error {
	return nil
}

func (c Counter) Deactivate(ctx context.Context) error {
	return nil
}

func (c Counter) IAmMetricsCounter() {}
