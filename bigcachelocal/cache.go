package bigcachelocal

import (
	"context"
	"time"

	"github.com/Zhanat87/common-libs/contracts"
	"github.com/allegro/bigcache/v3"
)

type cache struct {
	bigCache *bigcache.BigCache
}

func NewCache(bigCache *bigcache.BigCache) contracts.Cache {
	return &cache{bigCache}
}

func (s *cache) Set(_ context.Context, key string, value interface{}, _ time.Duration) error {
	return s.bigCache.Set(key, value.([]byte))
}

func (s *cache) Get(_ context.Context, key string) (interface{}, error) {
	return s.bigCache.Get(key)
}

func (s *cache) Exists(_ context.Context, key string) (bool, error) {
	_, err := s.bigCache.Get(key)
	switch err {
	case bigcache.ErrEntryNotFound:
		return false, nil
	case nil:
		return true, nil
	default:
		return false, err
	}
}

func (s *cache) Delete(_ context.Context, key string) (int64, error) {
	err := s.bigCache.Delete(key)
	if err != nil {
		return 0, err
	}

	return 1, nil
}
