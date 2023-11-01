package conf

type RPCServerConf struct {
	RPCServer map[string]struct { // map[server_name-server_id]config
		Enable      bool   `toml:"Enable"`
		Description string `toml:"Description"`
	} `toml:"RPCServer"`
}
