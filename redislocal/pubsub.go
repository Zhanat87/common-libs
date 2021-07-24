package redislocal

import (
	"context"

	"github.com/Zhanat87/common-libs/contracts"
	"github.com/go-redis/redis/v8"
)

type pubSub struct {
	client *redis.Client
}

func NewPubSub(client *redis.Client) contracts.PubSub {
	return &pubSub{client}
}

func (s *pubSub) Publish(ctx context.Context, channel string, message interface{}) interface{} {
	return nil
}

func (s *pubSub) Subscribe(ctx context.Context, channels ...string) interface{} {
	return nil
}

func (s *pubSub) Unsubscribe(ctx context.Context, channels ...string) error {
	return nil
}

/*
https://pkg.go.dev/github.com/go-redis/redis/v8#Client.Publish

sub := client.Subscribe(queryResp)
iface, err := sub.Receive()
if err != nil {
    // handle error
}

// Should be *Subscription, but others are possible if other actions have been
// taken on sub since it was created.
switch iface.(type) {
case *Subscription:
    // subscribe succeeded
case *Message:
    // received first message
case *Pong:
    // pong received
default:
    // handle error
}

ch := sub.Channel()
*/