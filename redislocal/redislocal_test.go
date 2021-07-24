package redislocal_test

import (
	"os"
	"testing"

	"github.com/Zhanat87/common-libs/redislocal"
	"github.com/go-redis/redis/v8"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMain(m *testing.M) {
	mySetupFunction()
	retCode := m.Run()
	myTeardownFunction()
	os.Exit(retCode)
}

func mySetupFunction() {
	println("start redislocal package testing")
}

func myTeardownFunction() {
	println("success end redislocal package testing")
}

func TestInitClient(t *testing.T) {
	Convey("Should init redis connection", t, func() {
		err := redislocal.InitDefaultClient()
		So(err, ShouldBeNil)
		So(redislocal.Client, ShouldHaveSameTypeAs, &redis.Client{})
	})
}
