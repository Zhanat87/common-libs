package gokitmiddlewares

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	kitoc "github.com/go-kit/kit/tracing/opencensus"
	tracingzipkin "github.com/go-kit/kit/tracing/zipkin"
	"github.com/openzipkin/zipkin-go"
	"github.com/openzipkin/zipkin-go/model"
)

type ZipkinTracing struct {
	zipkinTracer *zipkin.Tracer
	packageName  string
}

func NewZipkinTracing(zipkinTracer *zipkin.Tracer, packageName string) Tracer {
	return &ZipkinTracing{zipkinTracer, packageName}
}

func (s *ZipkinTracing) Trace(ctx context.Context, methodName string) (zipkin.Span, context.Context) {
	return s.zipkinTracer.StartSpanFromContext(
		ctx,
		s.packageName+" "+methodName,
	)
}

const TraceEndpointNamePrefix = "gokit:endpoint "

func GetTraceEndpoint(endPoint endpoint.Endpoint, endPointName string) endpoint.Endpoint {
	return kitoc.TraceEndpoint(TraceEndpointNamePrefix + endPointName)(endPoint)
}

func GetZipkinEndpoint(zipkinTracer *zipkin.Tracer, endPoint endpoint.Endpoint, endPointName string) endpoint.Endpoint {
	return tracingzipkin.TraceEndpoint(zipkinTracer, TraceEndpointNamePrefix+endPointName)(endPoint)
}

func TraceEndpoint(tracer *zipkin.Tracer, endPointName string) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			var sc model.SpanContext
			if parentSpan := zipkin.SpanFromContext(ctx); parentSpan != nil {
				sc = parentSpan.Context()
			}
			sp := tracer.StartSpan(endPointName, zipkin.Parent(sc))
			defer sp.Finish()
			ctx = zipkin.NewContext(ctx, sp)

			return next(ctx, request)
		}
	}
}
