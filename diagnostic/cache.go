package diagnostic

import (
	"github.com/BurntSushi/toml"
	"go-micro-framework/common/conf"
	"go-micro-framework/global/config"
)

var cache_conf = `
[Redis.test]
	Name = "Redis Server"
	Servers = ["192.168.130.128:16381"]
	Password = "123456"
`

func SetupCache() {
	cf := &conf.CacheConf{}
	if err := toml.Unmarshal([]byte(cache_conf), &cf); err != nil {
		panic(err)
	}
	config.SetCacheConf(cf)
}
