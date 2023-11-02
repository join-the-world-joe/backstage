package diagnostic

import (
	"backstage/common/conf"
	"backstage/global/config"
	"github.com/BurntSushi/toml"
)

var mysql_conf = `
#[MySQL.test.Master]
#	Host = "119.23.224.221"
#	Port = "13306"
#	User = "root"
#	Password = "123456"

[MySQL.test.Master]
	Host = "119.23.224.221"
	Port = "13306"
	User = "root"
	Password = "123456"
[[MySQL.test.Replicas]]
	Host = "119.23.224.221"
	Port = "13307"
	User = "root"
	Password = "123456"
`

func SetupMySQL() {
	cf := &conf.MySQLConf{}
	if err := toml.Unmarshal([]byte(mysql_conf), &cf); err != nil {
		panic(err)
	}
	config.SetMySQLConf(cf)
}
