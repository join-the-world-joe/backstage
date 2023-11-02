package conf

type RBACConf struct {
	Role map[string]struct {
		Rank        int    `toml:"Rank"`
		Department  string `toml:"Department"`
		Description string `toml:"Description"`
	} `toml:"Role"`
	Permission map[string]struct {
		Major       int    `toml:"Major"`
		Minor       int    `toml:"Minor"`
		Description string `toml:"Description"`
	} `toml:"Permission"`
	Menu map[string]struct {
		Item map[string]struct {
			Description string `toml:"Description"`
		} `toml:"Item"`
	} `toml:"Menu"`
	Table map[string]struct {
		Field map[string]struct {
			Description string `toml:"Description"`
		} `toml:"Field"`
	} `toml:"Table"`
	RolePermission map[string]struct {
		PermissionList []string `toml:"PermissionList"`
	} `toml:"RolePermission"`
	RoleMenu map[string]struct {
		Menu map[string]struct {
			ItemList []string `toml:"ItemList"`
		} `toml:"Menu"`
	} `toml:"RoleMenu"`
	RoleField map[string]struct {
		Table map[string]struct {
			FieldList []string `toml:"FieldList"`
		} `toml:"Table"`
	} `toml:"RoleField"`
}
