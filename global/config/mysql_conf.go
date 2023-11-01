package config

import "backstage/common/conf"

var _MySQLConf *conf.MySQLConf

func SetMySQLConf(cf *conf.MySQLConf) {
	_MySQLConf = cf
}

func MySQLConf() *conf.MySQLConf {
	return _MySQLConf
}
