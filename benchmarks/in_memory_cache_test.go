// go test -bench=. -benchmem -benchtime=4s -timeout 30m benchmarks/in_memory_cache_test.go
package benchmarks_test

import (
	"testing"
	"time"

	"github.com/Zhanat87/common-libs/bigcachelocal"
	"github.com/Zhanat87/common-libs/contracts"
	"github.com/Zhanat87/common-libs/fastcachelocal"
	"github.com/Zhanat87/common-libs/freecachelocal"
	"github.com/Zhanat87/common-libs/utils"
)

func inMemoryCacheSGED(cache contracts.InMemoryCache) {
	key := "test"
	cache.Set(key, []byte("test value")) // nolint
	cache.Get(key) // nolint
	cache.Exists(key) // nolint
	cache.Delete(key) // nolint
}

func BenchmarkInMemoryCacheFastCache(b *testing.B) { // 1
	startedAt := time.Now()
	defer utils.PrintBenchReport(b, startedAt, "InMemoryCache fastCache")
	fastCache := fastcachelocal.GetDefaultFastCache()
	inMemoryCache := fastcachelocal.NewInMemoryCache(fastCache)
	for i := 0; i < b.N; i++ {
		inMemoryCacheSGED(inMemoryCache)
	}
}

func BenchmarkInMemoryCacheBigCache(b *testing.B) { // 2
	startedAt := time.Now()
	defer utils.PrintBenchReport(b, startedAt, "InMemoryCache bigCache")
	bigCache, err := bigcachelocal.GetDefaultBigCache()
	if err != nil {
		b.Error(err)
	}
	inMemoryCache := bigcachelocal.NewInMemoryCache(bigCache)
	for i := 0; i < b.N; i++ {
		inMemoryCacheSGED(inMemoryCache)
	}
}

func BenchmarkInMemoryCacheFreeCache(b *testing.B) { // 3
	startedAt := time.Now()
	defer utils.PrintBenchReport(b, startedAt, "InMemoryCache freeCache")
	freeCache := freecachelocal.GetDefaultFreeCache()
	inMemoryCache := freecachelocal.NewInMemoryCache(freeCache)
	for i := 0; i < b.N; i++ {
		inMemoryCacheSGED(inMemoryCache)
	}
}
