package fastcachelocal

import (
	"github.com/VictoriaMetrics/fastcache"
	"github.com/Zhanat87/common-libs/contracts"
)

// https://pkg.go.dev/github.com/VictoriaMetrics/fastcache
type inMemoryCache struct {
	fastCache *fastcache.Cache
	dst       []byte
}

func NewInMemoryCache(fastCache *fastcache.Cache) contracts.InMemoryCache {
	return &inMemoryCache{fastCache: fastCache}
}

func (s *inMemoryCache) SetDst(dst []byte) contracts.InMemoryCache {
	s.dst = dst

	return s
}

func (s *inMemoryCache) Set(key string, value []byte) error {
	s.fastCache.SetBig([]byte(key), value)

	return nil
}

func (s *inMemoryCache) Get(key string) ([]byte, error) {
	return s.fastCache.GetBig(s.dst, []byte(key)), nil
}

func (s *inMemoryCache) Exists(key string) (bool, error) {
	_, ok := s.fastCache.HasGet(s.dst, []byte(key))

	return ok, nil
}

func (s *inMemoryCache) Delete(key string) (int64, error) {
	s.fastCache.Del([]byte(key))

	return 1, nil
}
