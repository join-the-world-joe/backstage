package conf

type RouteConf struct {
	Upstream struct {
		Protocol string `toml:"Protocol"`
	} `toml:"Upstream"`
	Downstream struct {
		Protocol string `toml:"Protocol"`
	} `toml:"Downstream"`
}
