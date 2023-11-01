package config

import "backstage/common/conf"

var _BrokerConf *conf.BrokerConf

func SetBrokerConf(cf *conf.BrokerConf) {
	_BrokerConf = cf
}

func BrokerConf() *conf.BrokerConf {
	return _BrokerConf
}

func BrokerEnable() bool {
	if len(_BrokerConf.Broker) > 0 {
		return true
	}
	return false
}
