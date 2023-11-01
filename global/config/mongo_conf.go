package config

import "backstage/common/conf"

var _MongoConf *conf.MongoConf

func SetMongoConf(cf *conf.MongoConf) {
	_MongoConf = cf
}

func MongoConf() *conf.MongoConf {
	return _MongoConf
}
