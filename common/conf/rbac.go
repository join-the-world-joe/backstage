package conf

type RBACConf struct {
	Role map[string]struct {
		Description string `toml:"Description"`
	} `toml:"Role"`
	Permission map[string]struct {
		Minor       int    `toml:"Minor"`
		Description string `toml:"Description"`
	} `toml:"Permission"`
	Menu map[string]struct {
		Item []string `toml:"Item"`
	} `toml:"Menu"`
	Table map[string]struct {
		Attribute []string `toml:"Attribute"`
	} `toml:"Table"`
	RolePermission map[string]struct {
		Permission []string `toml:"Permission"`
	} `toml:"RolePermission"`
	RoleMenu map[string]struct {
		Menu map[string]struct {
			Item []string `toml:"Item"`
		} `toml:"Menu"`
	} `toml:"RoleMenu"`
	RoleAttribute map[string]struct {
		Table map[string]struct {
			Attribute []string `toml:"Attribute"`
		} `toml:"Table"`
	} `toml:"RoleAttribute"`
}
