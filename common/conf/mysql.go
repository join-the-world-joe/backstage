package conf

type MySQLConf struct {
	MySQL map[string]struct { // map[which one]config or map[identity]config or map[signature]config
		Master struct {
			Host     string `toml:"Host"`
			Port     string `toml:"Port"`
			User     string `toml:"User"`
			Password string `toml:"Password"`
		} `toml:"Master"`
		Sources []struct {
			Host     string `toml:"Host"`
			Port     string `toml:"Port"`
			User     string `toml:"User"`
			Password string `toml:"Password"`
		} `toml:"Sources"`
		Replicas []struct {
			Host     string `toml:"Host"`
			Port     string `toml:"Port"`
			User     string `toml:"User"`
			Password string `toml:"Password"`
		} `toml:"Replicas"`
	} `toml:"MySQL"`
}
