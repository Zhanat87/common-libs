package gokithttp

import (
	"github.com/Zhanat87/common-libs/encoders"
	kitlog "github.com/go-kit/kit/log"
	kitoc "github.com/go-kit/kit/tracing/opencensus"
	zipkinkit "github.com/go-kit/kit/tracing/zipkin"
	kittransport "github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/openzipkin/zipkin-go"
)

func GetServerOptions(logger kitlog.Logger) []kithttp.ServerOption {
	return []kithttp.ServerOption{
		kithttp.ServerErrorHandler(kittransport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encoders.EncodeErrorJSON),
		kitoc.HTTPServerTrace(),
	}
}

func GetServerOptionsWithZipkinTracer(logger kitlog.Logger, tracer *zipkin.Tracer) []kithttp.ServerOption {
	return []kithttp.ServerOption{
		kithttp.ServerErrorHandler(kittransport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encoders.EncodeErrorJSON),
		zipkinkit.HTTPServerTrace(tracer),
	}
}
