package conf

import (
	"backstage/common/conf"
)

const (
	DefaultP2PChannelSize       = 1024
	DefaultWaitForCloseInterval = 30 // unit in second
)

type ServiceConf struct {
	conf.RateLimitingConf
	conf.RouteConf
	conf.BrokerConf
	conf.RPCServerConf
	conf.CacheConf
	conf.RBACConf
	conf.GracefulShutdownConf
	conf.MongoConf
	ServantConf
}

type ServantConf struct {
	Servant struct {
		QPS int `toml:"QPS"` // changeable

		Name     string `toml:"Name"`
		Endpoint string `toml:"Endpoint"`

		IPLimit bool `toml:"EnableIPLimit"` // changeable

		LogLevel string `toml:"LogLevel"` // changeable

		WebsocketEndpoint         string `toml:"WebsocketEndpoint"`
		WebsocketAuthReadDeadline int    `toml:"WebsocketAuthReadDeadline"` // changeable
		WebsocketReadDeadline     int    `toml:"WebsocketReadDeadline"`     // changeable
		WebsocketIdleTimeout      int    `toml:"WebsocketIdleTimeout"`      // changeable
		WebsocketReadLimit        int64  `toml:"WebsocketReadLimit"`        // changeable
		WebsocketReadBufferSize   int    `toml:"WebsocketReadBufferSize"`   // changeable
		WebsocketWriteBufferSize  int    `toml:"WebsocketWriteBufferSize"`  // changeable
		WaitForCloseInterval      int    `toml:"WaitForCloseInterval"`      // changeable
	} `toml:"Servant"`
	Encryption struct {
		Algorithm string `toml:"Algorithm"` // changeable
		Enable    bool   `toml:"Enable"`    // changeable
	} `toml:"Encryption"`
	Feedback struct { // TODO: 未验证、已验证的用户处理方式(断开、继续服务)
		Enable  bool   `toml:"Enable"` // changeable
		Major   string `toml:"Major"`
		Minor   string `toml:"Minor"`
		Message string `toml:"Message"`
	} `toml:"Feedback"`
}
