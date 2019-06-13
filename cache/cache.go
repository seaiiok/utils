package cache

import (
	"time"
	"utils/cache/util"
)

type Conf struct {
	Config map[string]interface{}
	Cache  *util.Cache
}

func NewCache() *Conf {
	gcInterval, _ := time.ParseDuration("10s")
	conf := new(Conf)
	conf.Config = make(map[string]interface{}, 0)
	conf.Cache = util.NewCache(util.NoExpiration, gcInterval)
	return conf
}
