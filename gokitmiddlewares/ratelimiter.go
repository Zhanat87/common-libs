package gokitmiddlewares

import (
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/ratelimit"
	"golang.org/x/time/rate"
)

func GetRateLimiterMiddleware(limiterInterval time.Duration, limiterBurst int) endpoint.Middleware {
	return ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(limiterInterval), limiterBurst))
}

func GetRateLimiterEndpoint(endPoint endpoint.Endpoint, limiterInterval time.Duration, limiterBurst int) endpoint.Endpoint {
	return GetRateLimiterMiddleware(limiterInterval, limiterBurst)(endPoint)
}
