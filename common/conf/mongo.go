package conf

type MongoConf struct {
	Mongo map[string]struct {
		URI string `toml:"URI"`
	} `toml:"Mongo"`
}
