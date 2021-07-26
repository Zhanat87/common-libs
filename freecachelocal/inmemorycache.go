package freecachelocal

import (
	"github.com/Zhanat87/common-libs/contracts"
	"github.com/coocood/freecache"
)

type inMemoryCache struct {
	freeCache *freecache.Cache
}

func NewInMemoryCache(freeCache *freecache.Cache) contracts.InMemoryCache {
	return &inMemoryCache{freeCache: freeCache}
}

func (s *inMemoryCache) Set(key string, value []byte) error {
	s.freeCache.Set([]byte(key), value, 0)

	return nil
}

func (s *inMemoryCache) Get(key string) ([]byte, error) {
	return s.freeCache.Get([]byte(key))
}

func (s *inMemoryCache) Exists(key string) (bool, error) {
	_, err := s.freeCache.Get([]byte(key))
	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *inMemoryCache) Delete(key string) (int64, error) {
	ok := s.freeCache.Del([]byte(key))
	if ok {
		return 1, nil
	}

	return 0, nil
}
