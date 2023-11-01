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
	conf.SMSConf
	conf.GracefulShutdownConf
	ServantConf
}

type ServantConf struct {
	Servant struct {
		Name               string `toml:"Name"`
		LogLevel           string `toml:"LogLevel"`           // changeable
		OTPLenOfRegister   int    `toml:"OTPLenOfRegister"`   // one-time passcode, changeable
		OTPBeginOfRegister int    `toml:"OTPBeginOfRegister"` // OTPBeginOfRegister <= digit
		OTPEndOfRegister   int    `toml:"OTPEndOfRegister"`   // digit >= OTPEndOfRegister
	} `toml:"Servant"`
}
