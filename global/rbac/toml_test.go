package rbac

import (
	"github.com/BurntSushi/toml"
	"testing"
)

type RBACConf struct {
	Role map[string]struct {
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

type RoleConf struct {
	Role map[string]struct {
		Description string `toml:"Description"`
	} `toml:"Role"`
}

func TestRoleToml(t *testing.T) {
	config := `
	[Role.Role1]
		Description = "Role1角色描述"
	[Role.Role2]
		Description = "Role2角色描述"
	`
	cf := &RoleConf{}
	err := toml.Unmarshal([]byte(config), cf)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(cf)
}

type PermissionConf struct {
	Permission map[string]struct {
		Major       string `toml:"Major"`
		Minor       string `toml:"Minor"`
		Description string `toml:"Description"`
	} `toml:"Permission"`
}

func TestPermissionToml(t *testing.T) {
	config := `
	[Permission.FetchXXX1]
		Major = "Major1"
		Minor = "Minor1"
		Description = "FetchXXX1-Major1-Minor1"
	[Permission.FetchXXX2]
		Major = "Major2"
		Minor = "Minor2"
		Description = "FetchXXX2-Major2-Minor2"
	`
	cf := &PermissionConf{}
	err := toml.Unmarshal([]byte(config), cf)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(cf)
}

type MenuConf struct {
	Menu map[string]struct {
		Item map[string]struct {
			Description string `toml:"Description"`
		} `toml:"Item"`
	} `toml:"Menu"`
}

func TestMenuToml(t *testing.T) {
	config := `
	[Menu.Admission]
		Item.User.Description = "用户管理"
		Item.Role.Description = "角色管理"
		Item.Menu.Description = "菜单管理"
		Item.Permission.Description = "权限管理"
		Item.Field.Description = "字段管理"
		Item.Track.Description = "操作日志"
	`
	cf := &MenuConf{}
	err := toml.Unmarshal([]byte(config), cf)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(cf)
}

type TableConf struct {
	Table map[string]struct {
		Field map[string]struct {
			Description string `toml:"Description"`
		} `toml:"Field"`
	} `toml:"Table"`
}

func TestFieldToml(t *testing.T) {
	config := `
	[Table.user]
		Field.name.Description = "姓名"
		Field.country_code.Description = "电话区域号"
		Field.phone_number.Description = "电话号码"
	`
	cf := &TableConf{}
	err := toml.Unmarshal([]byte(config), cf)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(cf)
}

type RolePermissionConf struct {
	RolePermission map[string]struct {
		Permission string `toml:"Permission"`
	} `toml:"Table"`
}

func TestRolePermissionConf(t *testing.T) {
	config := `
	[RolePermission.Manager]
		Permission = "permission1"
	`
	cf := &RolePermissionConf{}
	err := toml.Unmarshal([]byte(config), cf)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(cf)
}
