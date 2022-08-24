package jaeger

import (
	"io"

	sdk_config "github.com/discomco/go-cart/config"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaeger_config "github.com/uber/jaeger-client-go/config"
)

func NewJaegerTracer(tracerConfig sdk_config.IJaegerConfig) (opentracing.Tracer, io.Closer, error) {
	cfg := &jaeger_config.Configuration{
		ServiceName: tracerConfig.GetServiceName(),

		// "const" sampler is a binary sampling strategy: 0=never sample, 1=always sample.
		Sampler: &jaeger_config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},

		// Log the emitted spans to stdout.
		Reporter: &jaeger_config.ReporterConfig{
			LogSpans:           tracerConfig.UseLogSpans(),
			LocalAgentHostPort: tracerConfig.GetHostPort(),
		},
	}
	return cfg.NewTracer(jaeger_config.Logger(jaeger.StdLogger))
}
