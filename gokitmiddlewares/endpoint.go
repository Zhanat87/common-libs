package gokitmiddlewares

import (
	"github.com/Zhanat87/common-libs/tracers"
	"github.com/go-kit/kit/endpoint"
	"time"
)

func GetDefaultEndpoint(endPoint endpoint.Endpoint, endPointName string) endpoint.Endpoint {
	return GetCircuitBreakerEndpoint(
		GetRateLimiterEndpoint(
			GetZipkinEndpoint(tracers.ZipkinTracer, endPoint,
				endPointName), 100), endPointName, 30 * time.Second)
}

func GetEndpoint(endPoint endpoint.Endpoint, endPointName string, circuitBreakerTimeout time.Duration) endpoint.Endpoint {
	return GetCircuitBreakerMiddleware(name, timeout)(endPoint)
}
