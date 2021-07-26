package bigcachelocal_test

import (
	"os"
	"testing"

	"github.com/Zhanat87/common-libs/bigcachelocal"
	"github.com/Zhanat87/common-libs/utils"
	"github.com/allegro/bigcache/v3"
	. "github.com/smartystreets/goconvey/convey"
)

var bigCache *bigcache.BigCache

func TestMain(m *testing.M) {
	mySetupFunction()
	retCode := m.Run()
	myTeardownFunction()
	os.Exit(retCode)
}

func mySetupFunction() {
	println("start bigcachelocal package testing")
	var err error
	bigCache, err = bigcachelocal.GetDefaultBigCache()
	if err != nil {
		panic(err)
	}
}

func myTeardownFunction() {
	println("success end bigcachelocal package testing")
}

func TestCache(t *testing.T) {
	Convey("Should success bigCache", t, func() {
		cache := bigcachelocal.NewCache(bigCache)
		utils.TestCache(cache, []byte("test value"))
	})
}

func TestInMemoryCache(t *testing.T) {
	Convey("Should success in memory bigCache", t, func() {
		inMemoryCache := bigcachelocal.NewInMemoryCache(bigCache)
		utils.TestInMemoryCache(inMemoryCache, bigcache.ErrEntryNotFound, nil)
	})
}
