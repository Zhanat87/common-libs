package gokitmiddlewares

import "time"

type Saver interface {
	Save(err error, begin time.Time, methodName string)
}
