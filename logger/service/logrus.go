package service

import log "github.com/sirupsen/logrus"

type logrus struct {
	logger *log.Logger
}

func NewLogrus(logger *log.Logger) Logger {
	return &logrus{
		logger: logger,
	}
}

func (l *logrus) Error(err error) {
	l.logger.Error(err.Error())
}

func (l *logrus) Warning(warning string) {
	l.logger.Warning(warning)
}

func (l *logrus) Info(info string) {
	l.logger.Info(info)
}

func (l *logrus) Debug(debug string) {
	l.logger.Debug()
}
