package diagnostic

import (
	"github.com/BurntSushi/toml"
	"go-micro-framework/common/conf"
	"go-micro-framework/global/config"
)

var mysql_conf = `
[MySQL.test.Master]
	Host = "192.168.130.128"
	Port = "12305"
	User = "root"
	Password = "123456"

[MySQL.HongKong.Master]
	Host = "192.168.130.128"
	Port = "12305"
	User = "root"
	Password = "123456"

[MySQL.HongKong1.Master]
	Host = "192.168.130.128"
	Port = "12305"
	User = "root"
	Password = "123456"

[MySQL.HongKong2.Master]
	Host = "192.168.130.128"
	Port = "12306"
	User = "root"
	Password = "123456"

[MySQL.HongKong3.Master]
	Host = "192.168.130.128"
	Port = "12307"
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
