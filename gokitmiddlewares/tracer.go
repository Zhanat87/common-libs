package gokitmiddlewares

import "context"

type Tracer interface {
	Trace(ctx context.Context, methodName string)
}
