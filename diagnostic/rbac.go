package diagnostic

import (
	"backstage/common/conf"
	"backstage/global/config"
	"github.com/BurntSushi/toml"
)

var rbac_conf = `
[Role.Administrator]
	Rank = 3
	Department = "Board of Directors"
	Description = "The top role of this organization"

[Role.RD_Manager]
	Rank = 2
	Department = "Research and Development"
	Description = "The manager of Research and Development department"
[Role.Software_Engineer]
	Rank = 1
	Department = "Research and Development"
	Description = "A Software Engineer of Research and Development department"
[Role.Hardware_Engineer]
	Rank = 1
	Department = "Research and Development"
	Description = "A Hardware Engineer of Research and Development department"

[Role.Finance_Manger]
	Rank = 2
	Department = "Finance Department"
	Description = "The manager of Finance Department"
[Role.Purchasing_Specialist]
	Rank = 1
	Department = "Finance Department"
	Description = "A Purchasing Specialist of Finance Department"
[Role.Accounting_Specialist]
	Rank = 1
	Department = "Finance Department"
	Description = "A Accounting Specialist of Finance Department"

[Role.HR_Manger]
	Rank = 2
	Department = "Human Resources Department"
	Description = "The manager of Human Resources Department"
[Role.HR_Specialist]
	Rank = 1
	Department = "Human Resources Department"
	Description = "A Human Resources Specialist of Human Resources Department"

[Role.Marketing_Manger]
	Rank = 2
	Department = "Marketing Department"
	Description = "The manager of Marketing Department"
[Role.Sales_Specialist]
	Rank = 1
	Department = "Marketing Department"
	Description = "A Sales Specialist of Marketing Department"

[Role.Manufacturing_Manger]
	Rank = 2
	Department = "Manufacturing Department"
	Description = "The manager of Manufacturing Department"
[Role.Production_Specialist]
	Rank = 1
	Department = "Manufacturing Department"
	Description = "A Production Specialist of Manufacturing Department"	

[Permission.FetchMenuListOfCondition]
	Major = 5
	Minor = 3
	Description = "FetchMenuListOfCondition"
[Permission.FetchUserListOfCondition]
	Major = 5
	Minor = 5
	Description = "FetchUserListOfCondition"	
[Permission.FetchRoleListOfCondition]
	Major = 5
	Minor = 7
	Description = "FetchRoleListOfCondition"
[Permission.FetchPermissionListOfCondition]
	Major = 5
	Minor = 9
	Description = "FetchPermissionListOfCondition"
[Permission.InsertUserRecord]
	Major = 5
	Minor = 11
	Description = "InsertUserRecord"
[Permission.SoftDeleteUserRecord]
	Major = 5
	Minor = 13
	Description = "SoftDeleteUserRecord"
[Permission.UpdateUserRecord]
	Major = 5
	Minor = 15
	Description = "UpdateUserRecord"	

[Menu.Admission]
	Item.User.Description = "用户管理"
	Item.Role.Description = "角色管理"
	Item.Menu.Description = "菜单管理"
	Item.Permission.Description = "权限管理"
	Item.Field.Description = "字段管理"
	Item.Track.Description = "操作日志"

[Table.user]
	Field.id.Description = "用户ID"
	Field.name.Description = "姓名"
	Field.country_code.Description = "国家地区码"
	Field.phone_number.Description = "电话号码"
	Field.status.Description = "状态"
	Field.created_at.Description = "创建日期"

[RolePermission.RD_Manager]
	PermissionList = ['FetchMenuListOfCondition', 'FetchUserListOfCondition', 'FetchRoleListOfCondition', 'UpdateUserRecord']
[RolePermission.Finance_Manger]
	PermissionList = ['FetchMenuListOfCondition', 'FetchUserListOfCondition', 'FetchRoleListOfCondition', 'UpdateUserRecord']
[RolePermission.Manufacturing_Manger]
	PermissionList = ['FetchMenuListOfCondition', 'FetchUserListOfCondition', 'FetchRoleListOfCondition', 'UpdateUserRecord']
[RolePermission.Worker]
	PermissionList = ['FetchMenuListOfCondition', 'FetchUserListOfCondition']
[RolePermission.Sales]
	PermissionList = ['FetchUserListOfCondition']

[RoleMenu.RD_Manager.Menu.Admission]
	ItemList = ['User']
[RoleMenu.Finance_Manger.Menu.Admission]
	ItemList = ['User']
[RoleMenu.Manufacturing_Manger.Menu.Admission]
	ItemList = ['User']

[RoleField.RD_Manager.Table.user]
	FieldList = ['name', 'id', 'gender', 'account', 'country_code', 'phone_number', 'status', 'created_at']
[RoleField.Finance_Manger.Table.user]
	FieldList = ['name', 'id', 'gender', 'account', 'country_code', 'phone_number', 'status']
[RoleField.Manufacturing_Manger.Table.user]
	FieldList = ['name', 'id', 'gender', 'account', 'country_code', 'phone_number', 'status']


[RoleField.High.Table.user]
	FieldList = ['name', 'id', 'gender', 'account', 'country_code', 'phone_number', 'status', 'created_at']

[RoleField.Middle.Table.user]
	FieldList = ['name', 'id', 'gender', 'account', 'country_code', 'phone_number', 'status', 'created_at']
`

func SetupRBAC() {
	cf := &conf.RBACConf{}
	err := toml.Unmarshal([]byte(rbac_conf), cf)
	if err != nil {
		panic(err)
	}
	config.SetRBACConf(cf)
}
