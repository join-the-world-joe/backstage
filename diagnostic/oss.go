package diagnostic

import (
	"backstage/common/conf"
	"backstage/global/config"
	"github.com/BurntSushi/toml"
)

var oss_conf = `
[OSS.AliYun]
	ID = ""
	Secret = ""
	Endpoint = "oss-cn-shenzhen.aliyuncs.com"
`

func SetupOSS() {
	cf := &conf.OSSConf{}
	if err := toml.Unmarshal([]byte(oss_conf), &cf); err != nil {
		panic(err)
	}
	config.SetOSSConf(cf)
}
