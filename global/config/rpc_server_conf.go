package config

import (
	"backstage/common/conf"
)

var _RPCServerConf *conf.RPCServerConf

func SetRPCServerConf(cf *conf.RPCServerConf) {
	_RPCServerConf = cf
}

func RPCServerConf() *conf.RPCServerConf {
	return _RPCServerConf
}

func RPCEnable(service string) bool {
	if temp, exist := _RPCServerConf.RPCServer[service]; exist {
		return temp.Enable
	}
	return false
}
