package jaeger

import (
	"io"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/server"
	internal_opentracing "github.com/kitex-contrib/tracer-opentracing"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

type ClientTracer struct{}

func NewClientTracer() *ClientTracer {
	return &ClientTracer{}
}

func (ClientTracer) Init(serviceName string) (client.Suite, io.Closer) {
	cfg := config.Configuration{
		ServiceName: serviceName,
		Disabled:    jeagerDisabled,
		Sampler: &config.SamplerConfig{
			Type:  jeagerSamplerType,
			Param: float64(jeagerSamplerParam),
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           jeagerReporterLogSpans,
			LocalAgentHostPort: jeagerAgentHostPort,
		},
	}
	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(err)
	}
	opentracing.InitGlobalTracer(tracer)
	return internal_opentracing.NewDefaultClientSuite(), closer
}

type ServerTracer struct{}

func NewServerSuite() *ServerTracer {
	return &ServerTracer{}
}

func (ServerTracer) Init(serviceName string) (server.Suite, io.Closer) {
	cfg := config.Configuration{
		ServiceName: serviceName,
		Disabled:    jeagerDisabled,
		Sampler: &config.SamplerConfig{
			Type:  jeagerSamplerType,
			Param: float64(jeagerSamplerParam),
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           jeagerReporterLogSpans,
			LocalAgentHostPort: jeagerAgentHostPort,
		},
	}
	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(err)
	}
	opentracing.InitGlobalTracer(tracer)
	return internal_opentracing.NewDefaultServerSuite(), closer
}
