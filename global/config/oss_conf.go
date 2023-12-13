package config

import "backstage/common/conf"

var _OSSConf *conf.OSSConf

func SetOSSConf(cf *conf.OSSConf) {
	_OSSConf = cf
}

func OSSConf() *conf.OSSConf {
	return _OSSConf
}
