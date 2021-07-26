package utils

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/Zhanat87/common-libs/contracts"
	"github.com/allegro/bigcache/v3"
	. "github.com/smartystreets/goconvey/convey" // nolint
)

func PrintBenchReport(b *testing.B, startedAt time.Time, cacheBackend string) {
	elapsed := time.Since(startedAt)
	fmt.Printf("\n%s\nElapsed: %s\nSGED per second: %f\n",
		cacheBackend, elapsed, float64(b.N)/elapsed.Seconds())
}

func TestCache(cache contracts.Cache, value interface{}) {
	So(cache, ShouldImplement, (*contracts.Cache)(nil))

	ctx := context.Background()
	key := "test"
	duration := 1 * time.Second

	res, err := cache.Get(ctx, key)
	So(res, ShouldBeEmpty)
	So(err, ShouldEqual, bigcache.ErrEntryNotFound)

	err = cache.Set(ctx, key, value, duration)
	So(err, ShouldBeNil)

	res, err = cache.Get(ctx, key)
	So(err, ShouldBeNil)
	So(res, ShouldResemble, value)

	ok, err := cache.Exists(ctx, key)
	So(err, ShouldBeNil)
	So(ok, ShouldBeTrue)

	deletedCount, err := cache.Delete(ctx, key)
	So(err, ShouldBeNil)
	So(deletedCount, ShouldEqual, 1)

	ok, err = cache.Exists(ctx, key)
	So(err, ShouldBeNil)
	So(ok, ShouldBeFalse)

	err = cache.Set(ctx, key, value, duration)
	So(err, ShouldBeNil)

	deletedCount, err = cache.Delete(ctx, key)
	So(err, ShouldBeNil)
	So(deletedCount, ShouldEqual, 1)
}

func TestInMemoryCache(inMemoryCache contracts.InMemoryCache, getError, existsError error) {
	So(inMemoryCache, ShouldImplement, (*contracts.InMemoryCache)(nil))

	key := "test"
	value := []byte("test value")

	res, err := inMemoryCache.Get(key)
	So(res, ShouldBeEmpty)
	So(err, ShouldEqual, getError)

	err = inMemoryCache.Set(key, value)
	So(err, ShouldBeNil)

	res, err = inMemoryCache.Get(key)
	So(err, ShouldBeNil)
	So(res, ShouldResemble, value)

	ok, err := inMemoryCache.Exists(key)
	So(err, ShouldBeNil)
	So(ok, ShouldBeTrue)

	deletedCount, err := inMemoryCache.Delete(key)
	So(err, ShouldBeNil)
	So(deletedCount, ShouldEqual, 1)

	ok, err = inMemoryCache.Exists(key)
	So(err, ShouldEqual, existsError)
	So(ok, ShouldBeFalse)

	err = inMemoryCache.Set(key, value)
	So(err, ShouldBeNil)

	deletedCount, err = inMemoryCache.Delete(key)
	So(err, ShouldBeNil)
	So(deletedCount, ShouldEqual, 1)
}
