package fastcachelocal

import (
	"github.com/VictoriaMetrics/fastcache"
	"github.com/Zhanat87/common-libs/utils"
)

func GetDefaultFastCache() *fastcache.Cache {
	return fastcache.New(utils.GetGigabytes(1))
}
