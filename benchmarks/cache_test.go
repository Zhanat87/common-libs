package benchmarks_test

import (
	"context"
	"testing"
	"time"

	"github.com/Zhanat87/common-libs/bigcachelocal"
	"github.com/Zhanat87/common-libs/contracts"
	"github.com/Zhanat87/common-libs/redislocal"
	"github.com/Zhanat87/common-libs/utils"
)

func cacheSGED(cache contracts.Cache, value interface{}) {
	key := "test"
	duration := 1 * time.Second
	ctx := context.Background()
	cache.Set(ctx, key, value, duration)
	cache.Get(ctx, key)
	cache.Exists(ctx, key)
	cache.Delete(ctx, key)
}

// https://habr.com/ru/post/268585/
// go test -bench=. -benchmem -benchtime=4s -timeout 30m benchmarks/cache_test.go
// go test -bench=. benchmarks/cache_test.go
// cd benchmarks && go test -bench CacheRedis -run=^$
// go test -bench=. -benchmem -benchtime=4s -timeout 30m benchmarks/*_test.go
func BenchmarkCacheRedis(b *testing.B) {
	startedAt := time.Now()
	defer utils.PrintBenchReport(b, startedAt, "cache redis")
	cache := redislocal.NewCache(redislocal.GetDefaultInstance())
	// задать пропускную способность за одну итерацию в байтах
	// при помощи метода b.SetBytes(n int64)
	// b.SetBytes(2)
	// чтобы не учитывать не нужный код в бенчмарке, надо запустить b.ResetTimer() и
	// перед кодом, который нужно пропустить надо поставить b.StopTimer()
	// а перед кодом, который нужно измерить надо поставить b.StartTimer()
	/*
		измерить прирост (или падение) производительности
		go test -bench=. -benchmem bench_test.go > new.txt
		git stash
		go test -bench=. -benchmem bench_test.go > old.txt
		сравнить эти результаты с помощью утилиты benchcmp.
		установить ее, выполнив команду go get golang.org/x/tools/cmd/benchcmp.
		Вот результаты сравнения: # benchcmp old.txt new.txt
	*/
	/*
		записать cpu и memory профили во время выполнения бенчмарков:
		go test -bench=. -benchmem -cpuprofile=cpu.out -memprofile=mem.out bench_test.go
		про анализ профилей вы можете прочитать отличный пост в официальном блоге Go:
		https://blog.golang.org/pprof
	*/
	//b.ResetTimer()
	value := "test value"
	for i := 0; i < b.N; i++ {
		// b.StopTimer()
		// time.Sleep(22 * time.Millisecond)
		// b.StartTimer()
		cacheSGED(cache, value)
		// показать количество байт и аллокаций памяти за итерацию
		// b.ReportAllocs()
	}
	/*
		BenchmarkCache
		Elapsed: 1.209692821s (всего времени на тест)
		SGED per second: 5845.285578
		    7071 (всего операций за тест в цикле)
		171666 ns/op (среднее время на одну операцию в цикле)
		170 B/op (количество байт памяти за итерацию)
		220 allocs/op (количество аллокаций памяти за итерацию)
	*/
}

func BenchmarkCacheBigCache(b *testing.B) {
	startedAt := time.Now()
	defer utils.PrintBenchReport(b, startedAt, "cache bigCache")
	bigCache, err := bigcachelocal.GetDefaultBigCache()
	if err != nil {
		b.Error(err)
	}
	cache := bigcachelocal.NewCache(bigCache)
	value := []byte("test value")
	for i := 0; i < b.N; i++ {
		cacheSGED(cache, value)
	}
}
