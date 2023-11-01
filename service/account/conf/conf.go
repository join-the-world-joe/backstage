package conf

import "backstage/common/conf"

const ()

type ServiceConf struct {
	conf.RouteConf
	conf.RPCServerConf
	conf.BrokerConf
	conf.MySQLConf
	conf.CacheConf
	conf.RBACConf
	conf.MongoConf
	conf.GracefulShutdownConf
	ServantConf
}

type ServantConf struct {
	Servant struct {
		Name     string `toml:"Name"`
		LogLevel string `toml:"LogLevel"` // changeable
	} `toml:"Servant"`
}
