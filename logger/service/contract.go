package service

type Logger interface {
	Error(err error)
	Warning(warning string)
	Info(info string)
	Debug(debug string)
}
