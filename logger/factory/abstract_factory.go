package factory

import "github.com/Zhanat87/common-libs/logger/service"

type AbstractFactory interface {
	CreateLogger() service.Logger
}
