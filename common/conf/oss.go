package conf

type OSSConf struct {
	OSS map[string]struct {
		ID       string `toml:"ID"`
		Secret   string `toml:"Secret"`
		Endpoint string `toml:"Endpoint"`
	} `toml:"OSS"`
}
