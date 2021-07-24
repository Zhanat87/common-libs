package contracts

import (
	"context"
	"time"
)

type Cache interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) (err error)
	Get(ctx context.Context, key string) (res interface{}, err error)
	Exists(ctx context.Context, key string) (ok bool, err error)
	Delete(ctx context.Context, key string) (int64, error)
}
