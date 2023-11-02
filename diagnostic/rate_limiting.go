package diagnostic

import (
	"backstage/common/conf"
	"backstage/global/config"
	"github.com/BurntSushi/toml"
)

var rate_limiting_conf = `
[RateLimiting.SigIn]
	Major = 1
	Minor = 2
	Period = 200
`

func SetupRateLimiting() {
	cf := &conf.RateLimitingConf{}
	err := toml.Unmarshal([]byte(rate_limiting_conf), cf)
	if err != nil {
		panic(err)
	}
	config.SetRateLimiting(cf)
}
