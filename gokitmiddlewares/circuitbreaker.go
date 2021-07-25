package gokitmiddlewares

import (
	"time"

	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/sony/gobreaker"
)

func GetCircuitBreakerMiddleware(name string, timeout time.Duration) endpoint.Middleware {
	return circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    name,
		Timeout: timeout,
	}))
}

func GetCircuitBreakerEndpoint(endPoint endpoint.Endpoint, name string, timeout time.Duration) endpoint.Endpoint {
	return GetCircuitBreakerMiddleware(name, timeout)(endPoint)
}
