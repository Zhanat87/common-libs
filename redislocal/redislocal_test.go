package redislocal_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/Zhanat87/common-libs/contracts"
	"github.com/Zhanat87/common-libs/redislocal"
	"github.com/go-redis/redis/v8"
	. "github.com/smartystreets/goconvey/convey"
)

var ctx context.Context

func TestMain(m *testing.M) {
	mySetupFunction()
	retCode := m.Run()
	myTeardownFunction()
	os.Exit(retCode)
}

func mySetupFunction() {
	println("start redislocal package testing")
	ctx = context.Background()
}

func myTeardownFunction() {
	println("success end redislocal package testing")
}

func TestGetInstance(t *testing.T) {
	Convey("Should return redis client", t, func() {
		redisClient := redislocal.GetDefaultInstance()
		So(redisClient, ShouldHaveSameTypeAs, &redis.Client{})

		redisClient2 := redislocal.GetDefaultInstance()
		So(redisClient2, ShouldPointTo, redisClient)

		redisClient3 := redislocal.GetDefaultInstance()
		So(redisClient3, ShouldPointTo, redisClient)
		So(redisClient3, ShouldPointTo, redisClient2)
	})
}

func TestInitClient(t *testing.T) {
	Convey("Should init redis connection", t, func() {
		err := redislocal.InitDefaultClient()
		So(err, ShouldBeNil)
		So(redislocal.Client, ShouldHaveSameTypeAs, &redis.Client{})
	})
}

func TestCache(t *testing.T) {
	Convey("Should success cache", t, func() {
		redisClient := redislocal.GetDefaultInstance()
		So(redisClient, ShouldHaveSameTypeAs, &redis.Client{})

		cache := redislocal.NewCache(redisClient)
		So(cache, ShouldImplement, (*contracts.Cache)(nil))

		key := "test"
		duration := 1 * time.Second
		value := "test value"

		res, err := cache.Get(ctx, key)
		So(res, ShouldBeEmpty)
		So(err, ShouldEqual, redis.Nil)

		err = cache.Set(ctx, key, value, duration)
		So(err, ShouldBeNil)

		res, err = cache.Get(ctx, key)
		So(err, ShouldBeNil)
		So(res, ShouldEqual, value)

		ok, err := cache.Exists(ctx, key)
		So(err, ShouldBeNil)
		So(ok, ShouldBeTrue)

		time.Sleep(duration + time.Millisecond*22)

		ok, err = cache.Exists(ctx, key)
		So(err, ShouldBeNil)
		So(ok, ShouldBeFalse)

		deletedCount, err := cache.Delete(ctx, key)
		So(err, ShouldBeNil)
		So(deletedCount, ShouldBeZeroValue)

		err = cache.Set(ctx, key, value, duration)
		So(err, ShouldBeNil)

		deletedCount, err = cache.Delete(ctx, key)
		So(err, ShouldBeNil)
		So(deletedCount, ShouldEqual, 1)
	})
}

func TestPubSub(t *testing.T) {
	Convey("Should success pubSub", t, func() {
		redisClient := redislocal.GetDefaultInstance()
		So(redisClient, ShouldHaveSameTypeAs, &redis.Client{})

		pubSub := redislocal.NewPubSub(redisClient)
		So(pubSub, ShouldImplement, (*contracts.PubSub)(nil))
	})
}
