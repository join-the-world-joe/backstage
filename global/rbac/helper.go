package rbac

import (
	role2 "backstage/common/macro/role"
	"backstage/global/config"
	"backstage/global/log"
	json2 "backstage/utils/json"
	"strings"
)

func GetRoleList() []string {
	roleList := []string{}
	cf := config.RBACConf()
	if cf == nil {
		log.Error("GetRoleList failure, cf == nil")
		return roleList
	}
	for role, _ := range cf.Role {
		roleList = append(roleList, role)
	}
	return roleList
}

func GetPermissionList() []int {
	permissionList := []int{}
	cf := config.RBACConf()
	if cf == nil {
		log.Error("GetPermissionList failure, cf == nil")
		return permissionList
	}
	for _, permission := range cf.Permission {
		permissionList = append(permissionList, permission.Minor)
	}
	return permissionList
}

func GetMenuList() []byte {
	cf := config.RBACConf()
	if cf == nil {
		log.Error("GetMenuList failure, cf == nil")
		return nil
	}
	js := json2.New()
	for menu, v1 := range cf.Menu {
		items := []string{}
		for _, item := range v1.Item {
			items = append(items, item)
		}
		js.Set(menu, items)
	}
	bytes, err := js.Encode()
	if err != nil {
		log.Error("GetMenuList failure, js.Bytes err: ", err.Error())
		return nil
	}
	return bytes
}

func GetAttributeList() []byte {
	cf := config.RBACConf()
	if cf == nil {
		log.Error("GetAttribute failure, cf == nil")
		return nil
	}
	js := json2.New()
	for table, temp := range cf.Table {
		attributes := []string{}
		for _, attribute := range temp.Attribute {
			attributes = append(attributes, attribute)
		}
		js.Set(table, attributes)
	}
	bytes, err := js.Encode()
	if err != nil {
		log.Error("GetAttribute failure, js.Bytes err: ", err.Error())
		return nil
	}
	return bytes
}

func GetPermissionListOfRole(role string) []int {
	permissionList := []int{}

	cf := config.RBACConf()
	if cf == nil {
		log.Error("GetRolePermissionListOfRole failure, cf == nil")
		return permissionList
	}

	if role == role2.Administrator {
		return GetPermissionList()
	} else {
		// check if role in both Role and RolePermission
		if !HasRole(role) {
			log.ErrorF("GetPermissionListOfRole failure, role(%v) not in Role.Role", role)
			return permissionList
		}
		if _, exist := cf.RolePermission[role]; !exist {
			log.ErrorF("GetRolePermissionListOfRole failure, role(%v) not in RolePermission.Role", role)
			return permissionList
		}
		for _, permissionName := range cf.RolePermission[role].Permission {
			// check if permission in both RolePermission and Permission
			if permission, exist := cf.Permission[permissionName]; exist {
				permissionList = append(permissionList, permission.Minor)
			} else {
				log.ErrorF("GetRolePermissionListOfRole failure, RolePermission.Role(%v).Permission(%v) not in Permission.Permission", role, permission)
			}
		}
		return permissionList
	}
}

func GetMenuListOfRole(role string) []byte {
	cf := config.RBACConf()
	if cf == nil {
		log.Error("GetMenuListOfRole failure, cf == nil")
		return nil
	}

	if role == role2.Administrator {
		return GetMenuList()
	} else {
		// check if role in both Role and RoleMenu
		if !HasRole(role) {
			log.ErrorF("GetMenuListOfRole failure, role(%v) not in Role.Role", role)
			return nil
		}
		if _, exist := cf.RoleMenu[role]; !exist {
			log.ErrorF("GetMenuListOfRole failure, role(%v) not in RoleMenu.Role", role)
			return nil
		}
		js := json2.New()
		for menu, temp := range cf.RoleMenu[role].Menu {
			// check if menu in both RoleMenu and Menu
			if _, exist := cf.Menu[menu]; !exist {
				log.ErrorF("GetMenuListOfRole failure, RoleMenu.Role(%v).Menu(%v) not in Menu.Menu", role, menu)
				continue
			}
			items := []string{}
			for _, item := range temp.Item {
				has := false
				for _, tempItem := range cf.Menu[menu].Item {
					// check if item in both Menu and RoleMenu
					if strings.Compare(tempItem, item) == 0 {
						has = true
						break
					}
				}
				if has {
					items = append(items, item)
				} else {
					log.ErrorF("GetMenuListOfRole failure, RoleMenu(%v).Menu(%v).Item(%v) not in Menu.Item", role, menu, item)
				}
			}
			js.Set(menu, items)
		}
		bytes, err := js.Encode()
		if err != nil {
			log.Error("GetMenuListOfRole failure, js.Bytes err: ", err.Error())
			return nil
		}
		return bytes
	}
}

func GetAttributeListOfRole(role string) []byte {
	cf := config.RBACConf()
	if cf == nil {
		log.Error("GetAttributeListOfRole failure, cf == nil")
		return nil
	}

	if role == role2.Administrator {
		return GetAttributeList()
	} else {
		// check if role in both Role and RoleAttribute
		if !HasRole(role) {
			log.ErrorF("GetMenuListOfRole failure, role(%v) not in Role.Role", role)
			return nil
		}
		if _, exist := cf.RoleAttribute[role]; !exist {
			log.ErrorF("GetAttributeListOfRole failure, role(%v) not in RoleAttribute.Role", role)
			return nil
		}
		js := json2.New()
		for table, temp := range cf.RoleAttribute[role].Table {
			attributes := []string{}
			// check if table in both Table and RoleAttribute
			if _, exist := cf.Table[table]; !exist {
				log.ErrorF("GetMenuListOfRole failure, RoleAttribute.Role(%v).Table(%v) not in Table.Table", role, table)
				continue
			}
			for _, attribute := range temp.Attribute {
				has := false
				for _, tempAttribute := range cf.Table[table].Attribute {
					// check if attribute in both Table and RoleAttribute
					if strings.Compare(tempAttribute, attribute) == 0 {
						has = true
						break
					}
				}
				if has {
					attributes = append(attributes, attribute)
				} else {
					log.ErrorF("GetAttributeListOfRole failure, RoleAttribute(%v).Table(%v).Attribute(%v) not in Table.Attribute", role, table, attribute)
				}
			}
			js.Set(table, attributes)
		}
		bytes, err := js.Encode()
		if err != nil {
			log.Error("GetAttributeListOfRole failure, js.Bytes err: ", err.Error())
			return nil
		}
		return bytes
	}
}

func HasRole(role string) bool {
	cf := config.RBACConf()
	if cf == nil {
		log.Error("HasRole failure, cf == nil")
		return false
	}
	if _, exist := cf.Role[role]; !exist {
		return false
	}
	return true
}

func HasPermission(role string, minor int) bool {
	cf := config.RBACConf()
	if cf == nil {
		log.Error("HasPermission failure, cf == nil")
		return false
	}

	if role == role2.Administrator {
		return true
	} else {
		// check if role in both Role and RolePermission
		if !HasRole(role) {
			log.ErrorF("HasPermission failure, role(%v) not in Role", role)
			return false
		}
		if _, exist := cf.RolePermission[role]; !exist {
			log.ErrorF("HasPermission failure, role(%v) not in RolePermission", role)
			return false
		}

		// check if permission in both Permission and RolePermission
		for _, permission := range cf.RolePermission[role].Permission {
			if _, exist := cf.Permission[permission]; !exist {
				log.ErrorF("HasPermission failure, permission(%v) not in Permission", permission)
				continue
			}
			// check if endpoint in permission
			if cf.Permission[permission].Minor == minor {
				return true
			}
		}
		log.ErrorF("HasPermission failure, permission(%v) not in both Permission and RolePermission.Role(%v).Permission", minor, role)
		return false
	}
}

func HasMenu(role, menu, item string) bool {
	cf := config.RBACConf()
	if cf == nil {
		log.Error("HasPermission failure, cf == nil")
		return false
	}

	if role == role2.Administrator {
		return true
	} else {
		// check if role in both Role and RoleMenu
		if !HasRole(role) {
			log.ErrorF("HasMenu failure, role(%v) not in Role", role)
			return false
		}
		if _, exist := cf.RoleMenu[role]; !exist {
			log.ErrorF("HasMenu failure, role(%v) not in RoleMenu.Role", role)
			return false
		}
		// check if menu in both Menu and RoleMenu
		if _, exist := cf.Menu[menu]; !exist {
			log.ErrorF("HasMenu failure, menu(%v) not in Menu", menu)
			return false
		}
		if _, exist := cf.RoleMenu[role].Menu[menu]; !exist {
			log.ErrorF("HasMenu failure, menu(%v) not in RoleMenu.Role(%v).Menu", menu, role)
			return false
		}
		// check if item in both Menu and RoleMenu
		has := false
		for _, tempItem := range cf.Menu[menu].Item {
			if strings.Compare(tempItem, item) == 0 {
				has = true
			}
		}
		if !has {
			log.ErrorF("HasMenu failure, item(%v) not in Menu.(%v).Item", item, menu)
		}
		has = false
		for _, tempItem := range cf.RoleMenu[role].Menu[menu].Item {
			if strings.Compare(tempItem, item) == 0 {
				has = true
			}
		}
		if !has {
			log.ErrorF("HasMenu failure, item(%v) not in RoleMenu.Role(%v).Menu(%v).Item", item, role, menu)
			return false
		}

		return true
	}
}

func HasAttribute(role, table, attribute string) bool {
	cf := config.RBACConf()
	if cf == nil {
		log.Error("HasPermission failure, cf == nil")
		return false
	}

	if role == role2.Administrator {
		return true
	} else {
		// check if role in both Role and RoleAttribute
		if !HasRole(role) {
			log.ErrorF("HasAttribute failure, role(%v) not in Role", role)
			return false
		}
		if _, exist := cf.RoleAttribute[role]; !exist {
			log.ErrorF("HasAttribute failure, role(%v) not in RoleAttribute.Role", role)
			return false
		}
		// check if table in both Table and RoleAttribute
		if _, exist := cf.Table[table]; !exist {
			log.ErrorF("HasAttribute failure, table(%v) not in Table", table)
			return false
		}
		if _, exist := cf.RoleAttribute[role].Table[table]; !exist {
			log.ErrorF("HasAttribute failure, table(%v) not in RoleAttribute.Role(%v).Table", table, role)
			return false
		}
		// check if attribute in both Table and RoleAttribute
		has := false
		for _, tempAttribute := range cf.Table[table].Attribute {
			if strings.Compare(tempAttribute, attribute) == 0 {
				has = true
			}
		}
		if !has {
			log.ErrorF("HasAttribute failure, attribute(%v) not in Table(%v).Attribute", attribute, table)
			return false
		}
		has = false
		for _, tempAttribute := range cf.RoleAttribute[role].Table[table].Attribute {
			if strings.Compare(attribute, tempAttribute) == 0 {
				has = true
			}
		}
		if !has {
			log.ErrorF("HasAttribute failure, attribute(%v) not in RoleAttribute.Role(%v).Table(%v).Attribute", attribute, role, table)
			return false
		}

		return true
	}
}
