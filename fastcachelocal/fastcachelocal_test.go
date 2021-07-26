package fastcachelocal_test

import (
	"os"
	"testing"

	"github.com/VictoriaMetrics/fastcache"
	"github.com/Zhanat87/common-libs/fastcachelocal"
	"github.com/Zhanat87/common-libs/utils"
	. "github.com/smartystreets/goconvey/convey"
)

var fastCache *fastcache.Cache

func TestMain(m *testing.M) {
	mySetupFunction()
	retCode := m.Run()
	myTeardownFunction()
	os.Exit(retCode)
}

func mySetupFunction() {
	println("start fastcachelocal package testing")
	fastCache = fastcachelocal.GetDefaultFastCache()
}

func myTeardownFunction() {
	println("success end fastcachelocal package testing")
}

func TestInMemoryCache(t *testing.T) {
	Convey("Should success in memory cache", t, func() {
		inMemoryCache := fastcachelocal.NewInMemoryCache(fastCache)
		utils.TestInMemoryCache(inMemoryCache, nil, nil)
	})
}
