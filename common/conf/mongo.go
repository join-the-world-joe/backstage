package conf

type MongoConf struct {
	Mongo map[string]struct { // map[which one]config or map[identity]config or map[signature]config
		Servers  []string `toml:"Servers"`
		User     string   `toml:"User"`
		Password string   `toml:"Password"`
	} `toml:"Mongo"`
}
