package cache

import (
	"time"
	"utils/cache/util"
)

type Conf struct {
	Config map[string]interface{}
	cache  *util.Cache
}

func NewCache() *Conf {
	gcInterval, _ := time.ParseDuration("5s")
	defaultExpiration, _ := time.ParseDuration("30s")
	conf := new(Conf)
	conf.Config = make(map[string]interface{}, 0)
	conf.cache = util.NewCache(defaultExpiration, gcInterval)
	return conf
}

func (c *Conf) SaveToFile(path string) error {
	return c.cache.SaveToFile(path)
}

func (c *Conf) LoadFromFile(path string) error {
	return c.cache.LoadFromFile(path)
}
