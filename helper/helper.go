package helper

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/transport"
)

func NewJaegerTracer(domain, username, password, serviceName string) opentracing.Tracer {
	sender := transport.NewHTTPTransport(
		domain,
		transport.HTTPBasicAuth(username, password),
	)
	tracer, _ := jaeger.NewTracer(serviceName,
		jaeger.NewConstSampler(true),
		jaeger.NewRemoteReporter(sender))
	return tracer
}