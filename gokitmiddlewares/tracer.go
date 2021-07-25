package gokitmiddlewares

import (
	"context"

	"github.com/openzipkin/zipkin-go"
)

type Tracer interface {
	Trace(ctx context.Context, methodName string) (zipkin.Span, context.Context)
}
