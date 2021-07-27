package factory

import (
	"os"

	"github.com/Zhanat87/common-libs/logger/service"
	"github.com/sirupsen/logrus"
)

type LogrusFactory struct{}

func (lf *LogrusFactory) CreateLogger() service.Logger {
	logger := logrus.New()
	logger.SetReportCaller(true)
	loggingLevel := os.Getenv("LoggingLevel")
	if loggingLevel != "" {
		loggingLevel, err := logrus.ParseLevel(loggingLevel)
		if err != nil {
			logrus.Fatal("Failed to get logging level: ", err)
		}
		logger.SetLevel(loggingLevel)
	}

	return service.NewLogrus(logger)
}
