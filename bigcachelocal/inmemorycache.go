package bigcachelocal

import (
	"github.com/Zhanat87/common-libs/contracts"
	"github.com/allegro/bigcache/v3"
)

type inMemoryCache struct {
	bigCache *bigcache.BigCache
}

func NewInMemoryCache(bigCache *bigcache.BigCache) contracts.InMemoryCache {
	return &inMemoryCache{bigCache}
}

func (s *inMemoryCache) Set(key string, value []byte) error {
	return s.bigCache.Set(key, value)
}

func (s *inMemoryCache) Get(key string) ([]byte, error) {
	return s.bigCache.Get(key)
}

func (s *inMemoryCache) Exists(key string) (bool, error) {
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

func (s *inMemoryCache) Delete(key string) (int64, error) {
	err := s.bigCache.Delete(key)
	if err != nil {
		return 0, err
	}

	return 1, nil
}
