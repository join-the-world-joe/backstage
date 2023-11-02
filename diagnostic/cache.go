package diagnostic

import (
	"backstage/common/conf"
	"backstage/global/config"
	"github.com/BurntSushi/toml"
)

var cache_conf = `
[Redis.test]
	Name = "Redis Server"
	Servers = ["119.23.224.221:16381"]
	Password = "123456"
`

func SetupCache() {
	cf := &conf.CacheConf{}
	if err := toml.Unmarshal([]byte(cache_conf), &cf); err != nil {
		panic(err)
	}
	config.SetCacheConf(cf)
}
