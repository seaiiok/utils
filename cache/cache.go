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
	gcInterval, _ := time.ParseDuration("1s")
	defaultExpiration, _ := time.ParseDuration("3s")
	conf := new(Conf)
	conf.Config = make(map[string]interface{}, 0)
	conf.Cache = util.NewCache(defaultExpiration, gcInterval)
	return conf
}
