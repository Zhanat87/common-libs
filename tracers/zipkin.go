package tracers

import (
	"strconv"

	oczipkin "contrib.go.opencensus.io/exporter/zipkin"
	opentracing "github.com/opentracing/opentracing-go"
	zipkinot "github.com/openzipkin-contrib/zipkin-go-opentracing"
	"github.com/openzipkin/zipkin-go"
	"github.com/openzipkin/zipkin-go/model"
	zipkinreporter "github.com/openzipkin/zipkin-go/reporter"
	httpreporter "github.com/openzipkin/zipkin-go/reporter/http"
	"go.opencensus.io/trace"
)

const endpointURL = "http://localhost:9411/api/v2/spans"

var (
	ZipkinTracer   *zipkin.Tracer
	ZipkinReporter zipkinreporter.Reporter
)

func SetZipkinTracerAsOpentracingGlobalTracer(serviceName, hostPort string) (zipkinreporter.Reporter, error) {
	// Set-up our OpenCensus instrumentation with Zipkin backend
	reporter := httpreporter.NewReporter(endpointURL)
	localEndpoint, _ := zipkin.NewEndpoint(serviceName, hostPort) // hello, ":0"
	exporter := oczipkin.NewExporter(reporter, localEndpoint)
	// Always sample our traces for this demo.
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
	// Register our trace exporter.
	trace.RegisterExporter(exporter)
	tracer, err := zipkin.NewTracer(reporter)
	opentracing.SetGlobalTracer(zipkinot.Wrap(tracer))

	return reporter, err
}

func NewZipkinTracer(serviceName, port string) (*zipkin.Tracer, error) {
	portInt, _ := strconv.Atoi(port)
	portUint16 := uint16(portInt)
	// The reporter sends traces to zipkin server
	reporter := httpreporter.NewReporter(endpointURL)
	// Local endpoint represent the local service information
	localEndpoint := &model.Endpoint{ServiceName: serviceName, Port: portUint16}
	// Sampler tells you which traces are going to be sampled or not. In this case we will record 100% (1.00) of traces.
	sampler, err := zipkin.NewCountingSampler(1)
	if err != nil {
		return nil, err
	}
	t, err := zipkin.NewTracer(
		reporter,
		zipkin.WithSampler(sampler),
		zipkin.WithLocalEndpoint(localEndpoint),
	)
	if err != nil {
		return nil, err
	}

	return t, err
}

func NewZipkinTracerAndHTTPReporter(serviceName, hostPort string) (*zipkin.Tracer, zipkinreporter.Reporter, error) {
	// Set-up our OpenCensus instrumentation with Zipkin backend
	reporter := httpreporter.NewReporter(endpointURL)
	localEndpoint, _ := zipkin.NewEndpoint(serviceName, hostPort) // hello, ":0"
	exporter := oczipkin.NewExporter(reporter, localEndpoint)
	// Always sample our traces for this demo.
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
	// Register our trace exporter.
	trace.RegisterExporter(exporter)
	tracer, err := zipkin.NewTracer(reporter)

	return tracer, reporter, err
}

func InitZipkinTracerAndZipkinHTTPReporter(serviceName, hostPort string) error {
	ZipkinReporter = httpreporter.NewReporter(endpointURL)
	localEndpoint, _ := zipkin.NewEndpoint(serviceName, hostPort)
	exporter := oczipkin.NewExporter(ZipkinReporter, localEndpoint)
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
	trace.RegisterExporter(exporter)
	var zipkinTracerError error
	ZipkinTracer, zipkinTracerError = zipkin.NewTracer(ZipkinReporter)

	return zipkinTracerError
}

func InitZipkinTracerAsOpentracingGlobalTracerAndZipkinHTTPReporter(serviceName, hostPort string) error {
	ZipkinReporter = httpreporter.NewReporter(endpointURL)
	localEndpoint, _ := zipkin.NewEndpoint(serviceName, hostPort)
	exporter := oczipkin.NewExporter(ZipkinReporter, localEndpoint)
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
	trace.RegisterExporter(exporter)
	var zipkinTracerError error
	ZipkinTracer, zipkinTracerError = zipkin.NewTracer(ZipkinReporter)
	opentracing.SetGlobalTracer(zipkinot.Wrap(ZipkinTracer))

	return zipkinTracerError
}
