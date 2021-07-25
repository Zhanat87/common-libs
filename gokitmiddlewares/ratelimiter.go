package gokitmiddlewares

import (
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/ratelimit"
	"golang.org/x/time/rate"
)

func GetRateLimiterMiddleware(limiterBurst int) endpoint.Middleware {
	return ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), limiterBurst))
}
