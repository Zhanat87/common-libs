package freecachelocal_test

import (
	"os"
	"testing"

	"github.com/Zhanat87/common-libs/freecachelocal"
	"github.com/Zhanat87/common-libs/utils"
	"github.com/coocood/freecache"
	. "github.com/smartystreets/goconvey/convey"
)

var freeCache *freecache.Cache

func TestMain(m *testing.M) {
	mySetupFunction()
	retCode := m.Run()
	myTeardownFunction()
	os.Exit(retCode)
}

func mySetupFunction() {
	println("start freecachelocal package testing")
	freeCache = freecachelocal.GetDefaultFreeCache()
}

func myTeardownFunction() {
	println("success end freecachelocal package testing")
}

func TestInMemoryCache(t *testing.T) {
	Convey("Should success in memory cache", t, func() {
		inMemoryCache := freecachelocal.NewInMemoryCache(freeCache)
		utils.TestInMemoryCache(inMemoryCache, freecache.ErrNotFound, freecache.ErrNotFound)
	})
}
