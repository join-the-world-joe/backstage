package conf

type BrokerConf struct {
	Broker map[string]struct { // map[which one]config or map[identity]config or map[signature]config
		Category string                 `toml:"Category"` // as in common/broker/category.go
		Servers  []string               `toml:"Servers"`
		User     string                 `toml:"User"`
		Password string                 `toml:"Password"`
		Param    map[string]interface{} `toml:"Param"`
	} `toml:"Broker"`
}
