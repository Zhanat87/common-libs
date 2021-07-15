package middleware

import (
	"time"

	"github.com/go-kit/kit/metrics"
)

type Instrumenting struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	requestError   metrics.Counter
	packageName    string
}

func NewInstrumenting(counter metrics.Counter, latency metrics.Histogram,
	counterE metrics.Counter, packageName string) Saver {
	return &Instrumenting{
		requestCount:   counter,
		requestLatency: latency,
		requestError:   counterE,
		packageName:    packageName,
	}
}

func (s *Instrumenting) Save(err error, begin time.Time, methodName string) {
	labels := []string{"method", s.packageName + "_" + methodName}
	s.requestCount.With(labels...).Add(1)
	if err != nil {
		s.requestError.With(labels...).Add(1)
	}
	s.requestLatency.With(labels...).Observe(time.Since(begin).Seconds())
}
