package ioccontainer

import (
	loggerservice "github.com/Zhanat87/common-libs/logger/service"
	"github.com/golobby/container"
)

func GetLogger() loggerservice.Logger {
	var logger loggerservice.Logger
	container.Make(&logger)

	return logger
}
