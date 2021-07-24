package contracts

import (
	"context"
	"time"
)

type Cache interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string) (interface{}, error)
	Exists(ctx context.Context, key string) (bool, error)
	Delete(ctx context.Context, key string) (int64, error)
}
