package contracts

import "context"

type PubSub interface {
	Publish(ctx context.Context, channel string, message interface{}) interface{}
	Subscribe(ctx context.Context, channels ...string) interface{}
}
