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
	return s.client.Publish(ctx, channel, message)
}

func (s *pubSub) Subscribe(ctx context.Context, channels ...string) interface{} {
	return s.client.Subscribe(ctx, channels...)
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

https://www.programmersought.com/article/50064833942/
https://stackoverflow.com/questions/61967369/unsubscribe-from-redis-doesnt-seem-to-work
https://dev.to/jeroendk/how-to-use-redis-pub-sub-in-go-chat-application-part-3-2h4c
*/