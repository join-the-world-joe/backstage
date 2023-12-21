package conf

import "backstage/common/conf"

const (
	DefaultOTPLenOfRegister   = 4
	DefaultOTPBeginOfRegister = 0
	DefaultOTPEndOfRegister   = 9
)

type ServiceConf struct {
	conf.RouteConf
	conf.RPCServerConf
	conf.BrokerConf
	conf.MySQLConf
	conf.CacheConf
	conf.RBACConf
	conf.MongoConf
	conf.OSSConf
	conf.GracefulShutdownConf
	ServantConf
}

type ServantConf struct {
	Servant struct {
		Name     string `toml:"Name"`
		LogLevel string `toml:"LogLevel"` // changeable
	} `toml:"Servant"`
}
