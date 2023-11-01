package conf

type SMSConf struct {
	SMS struct {
		Behavior map[string]struct {
			Template string `toml:"Template"`
		} `toml:"Behavior"`
	} `toml:"SMS"`
}
