package redislocal

import (
	"context"
	"time"

	"github.com/Zhanat87/common-libs/contracts"
	"github.com/go-redis/redis/v8"
)

type cache struct {
	client *redis.Client
}

func NewCache(client *redis.Client) contracts.Cache {
	return &cache{client}
}

func (s *cache) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) (err error) {
	return
}

func (s *cache) Get(ctx context.Context, key string) (res interface{}, err error) {
	return
}

func (s *cache) Exists(ctx context.Context, key string) (ok bool, err error) {
	return
}

func (s *cache) Delete(ctx context.Context, key string) (int64, error) {
	return 0, nil
}
