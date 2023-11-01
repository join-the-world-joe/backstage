package diagnostic

import (
	"github.com/BurntSushi/toml"
	"go-micro-framework/common/conf"
	"go-micro-framework/global/config"
)

var rbac_conf = `
[Role.Administrator]
[Role.Manager]
	
[Permission.FetchRoleList]
	Minor = 5
[Permission.FetchMenuList]
	Minor = 3
[Permission.FetchPermissionList]
	Minor = 7
[Permission.FetchAttributeList]
	Minor = 9
[Permission.FetchMenuListOfRole]
	Minor = 11
[Permission.FetchPermissionListOfRole]
	Minor = 13
[Permission.FetchAttributeListOfRole]
	Minor = 15

[Menu.Admission]
	Item = ['User', 'Role', 'Permission', 'Menu', 'Track']
[Menu.Menu1]
	Item = ['item1', 'item2', 'item3', 'item4', 'item5']

[Table.user]
	Attribute = ['name', 'id', 'gender', 'account', 'avatar', 'status', 'country']

[RolePermission.Manager]
	Permission = ['FetchMenuList','FetchPermissionList', 'FetchRoleList', 'FetchAttributeList', 'FetchMenuListOfRole', 'FetchPermissionListOfRole', 'FetchAttributeListOfRole']

[RoleMenu.Manager.Menu.Admission]
	Item = ['User', 'Role', 'Permission', 'Menu']

[RoleAttribute.Manager.Table.user]
	Attribute = ['name', 'id', 'gender', 'account']
`

func SetupRBAC() {
	cf := &conf.RBACConf{}
	err := toml.Unmarshal([]byte(rbac_conf), cf)
	if err != nil {
		panic(err)
	}
	config.SetRBACConf(cf)
}
