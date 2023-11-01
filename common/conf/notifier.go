package conf

type NotifierConf struct {
	Notify map[string]struct {
		Id map[string]struct {
			CMD string `toml:"CMD"`
		} `toml:"Id"`
	} `toml:"Notify"`
}
