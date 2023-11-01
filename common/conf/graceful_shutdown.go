package conf

type GracefulShutdownConf struct {
	GracefulShutdown struct {
		Timeout       int `toml:"Timeout"`       // changeable
		CheckInterval int `toml:"CheckInterval"` // changeable
	} `toml:"GracefulShutdown"`
}
