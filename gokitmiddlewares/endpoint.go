package gokitmiddlewares

import (
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/openzipkin/zipkin-go"
)

func GetDefaultEndpoint(endPoint endpoint.Endpoint, endPointName string, zipkinTracer *zipkin.Tracer) endpoint.Endpoint {
	return GetCircuitBreakerEndpoint(GetRateLimiterEndpoint(GetZipkinEndpoint(
		zipkinTracer, endPoint, endPointName), time.Second, 100), endPointName, 30*time.Second)
}

func GetEndpoint(endPoint endpoint.Endpoint, endPointName string, zipkinTracer *zipkin.Tracer,
	limiterInterval time.Duration, limiterBurst int, circuitBreakerTimeout time.Duration) endpoint.Endpoint {
	return GetCircuitBreakerEndpoint(GetRateLimiterEndpoint(GetZipkinEndpoint(
		zipkinTracer, endPoint, endPointName), limiterInterval, limiterBurst), endPointName, circuitBreakerTimeout)
}
