package factory

import "github.com/Zhanat87/common-libs/logger/service"

type StdoutFactory struct{}

func (sf *StdoutFactory) CreateLogger() service.Logger {
	return service.NewStdout()
}
