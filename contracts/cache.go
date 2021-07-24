package interfaces

import (
	"context"
	"time"
)

type CacheInterface interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) (err error)
	Get(ctx context.Context, key string) (res string, err error)
	Exists(ctx context.Context, key string) (ok bool, err error)
	Delete(ctx context.Context, key string) (int64, error)
}
