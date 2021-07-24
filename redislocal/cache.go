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

func (s *cache) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return s.client.Set(ctx, key, value, expiration).Err()
}

func (s *cache) Get(ctx context.Context, key string) (interface{}, error) {
	return s.client.Get(ctx, key).Result()
}

func (s *cache) Exists(ctx context.Context, key string) (bool, error) {
	_, err := s.client.Get(ctx, key).Result()
	switch err {
	case redis.Nil:
		return false, nil
	case nil:
		return true, nil
	default:
		return false, err
	}
}

func (s *cache) Delete(ctx context.Context, key string) (int64, error) {
	return s.client.Del(ctx, key).Result()
}
