package middleware

import (
	"time"

	"github.com/go-kit/kit/log"
)

type Logging struct {
	logger      log.Logger
	packageName string
}

func NewLogging(logger log.Logger, packageName string) Saver {
	return &Logging{logger, packageName}
}

func (s *Logging) Save(err error, begin time.Time, methodName string) {
	if err != nil {
		_ = s.logger.Log(
			"method", s.packageName+"_" + methodName,
			"took", time.Since(begin),
			"err", err,
		)
	}
}
