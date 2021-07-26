package freecachelocal

import (
	"github.com/Zhanat87/common-libs/utils"
	"github.com/coocood/freecache"
)

func GetDefaultFreeCache() *freecache.Cache {
	return freecache.NewCache(utils.GetGigabytes(1))
}
