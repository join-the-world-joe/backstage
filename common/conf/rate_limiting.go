package conf

type RateLimitingConf struct {
	RateLimiting map[string]struct {
		Major  int `toml:"Major"`
		Minor  int `toml:"Minor"`
		Period int `toml:"Period"` // in millisecond
	} `toml:"RateLimiting"`
}
