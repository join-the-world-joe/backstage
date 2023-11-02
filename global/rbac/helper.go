package rbac

import (
	"backstage/common/macro/abbreviation"
	role2 "backstage/common/macro/role"
	"backstage/global/config"
	"backstage/global/log"
	"fmt"
	"github.com/spf13/cast"
	"golang.org/x/exp/slices"
	"strings"
)

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

func GetRole(name string) (string, string, string, int, bool) {
	cf := config.RBACConf()
	if cf == nil {
		log.Error("GetRole failure, cf == nil")
		return "", "", "", 0, false
	}
	if !HasRole(name) {
		log.ErrorF("GetRole failure, role(%v) not in Role", name)
		return "", "", "", 0, false
	}
	return name, cf.Role[name].Department, cf.Role[name].Description, cf.Role[name].Rank, true
}

func GetRankOfRole(role string) int {
	rank := 0
	if len(role) <= 0 {
		return rank
	}

	_, _, _, r, b := GetRole(role)
	if b {
		if r > rank {
			rank = r
		}
	}

	return rank
}

func GetTopRankOfRoleList(roleList []string) int {
	rank := 0
	if len(roleList) <= 0 {
		return rank
	}
	for _, v := range roleList {
		_, _, _, r, b := GetRole(v)
		if b {
			if r > rank {
				rank = r
			}
		}
	}
	return rank
}

func GetRoleListLERankInDepartment(rank int, department string) []string {
	cf := config.RBACConf()
	if cf == nil {
		log.Error("GetRoleListBERank failure, cf == nil")
		return nil
	}

	roleList := []string{}

	for roleName, role := range cf.Role {
		if role.Rank <= rank && strings.Compare(role.Department, department) == 0 {
			roleList = append(roleList, roleName)
		}
	}
	return roleList
}

func GetRoleListLERank(rank int) []string {
	cf := config.RBACConf()
	if cf == nil {
		log.Error("GetRoleListBERank failure, cf == nil")
		return nil
	}

	roleList := []string{}

	for roleName, role := range cf.Role {
		if role.Rank <= rank {
			roleList = append(roleList, roleName)
		}
	}
	return roleList
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
		for _, permission := range cf.RolePermission[role].PermissionList {
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

func GetItemListOfMenu(menu string) ([]string, []string) {
	cf := config.RBACConf()
	if cf == nil {
		log.Error("GetItemListOfMenu failure, cf == nil")
		return nil, nil
	}
	if _, exist := cf.Menu[menu]; !exist {
		log.Error(fmt.Sprintf("GetItemListOfMenu %s not in Menu", menu))
		return nil, nil
	}
	itemList := []string{}
	descriptionList := []string{}
	for itemName, item := range cf.Menu[menu].Item {
		itemList = append(itemList, itemName)
		if len(item.Description) > 0 {
			descriptionList = append(descriptionList, item.Description)
		} else {
			descriptionList = append(descriptionList, abbreviation.NA)
		}

	}
	return itemList, descriptionList
}

func GetMenuList() ([]string, [][]string, [][]string) {
	cf := config.RBACConf()
	if cf == nil {
		log.Error("GetMenuList failure, cf == nil")
		return nil, nil, nil
	}

	menuList := []string{}
	itemListSet := make([][]string, len(cf.Menu))
	descListSet := make([][]string, len(cf.Menu))

	index := 0
	for menuName, menu := range cf.Menu {
		if len(menu.Item) > 0 {
			menuList = append(menuList, menuName)
			itemListSet[index], descListSet[index] = GetItemListOfMenu(menuName)
			index++
		}
	}
	return menuList, itemListSet, descListSet
}

func GetItemListOfRoleMenu(role, menu string) ([]string, []string) {
	cf := config.RBACConf()
	if cf == nil {
		log.Error("GetMenuList failure, cf == nil")
		return nil, nil
	}

	if role == role2.Administrator {
		return GetItemListOfMenu(menu)
	} else {
		//check if role in both Role and RoleMenu
		if !HasRole(role) {
			log.ErrorF("GetItemListOfRoleMenu failure, %v not in Role", role)
			return nil, nil
		}
		if _, exist := cf.RoleMenu[role]; !exist {
			log.ErrorF("GetItemListOfRoleMenu failure, role(%v) not in RoleMenu", role)
			return nil, nil
		}
		if _, exist := cf.Menu[menu]; !exist {
			log.Warn(fmt.Sprintf("GetItemListOfRoleMenu %v not in Menu", menu))
			return nil, nil
		}
		if _, exist := cf.RoleMenu[role].Menu[menu]; !exist {
			log.Warn(fmt.Sprintf("GetItemListOfRoleMenu %s not in RoleMenu", menu))
			return nil, nil
		}
		itemList := []string{}
		descriptionList := []string{}
		for _, itemName1 := range cf.RoleMenu[role].Menu[menu].ItemList {
			for itemName2, item := range cf.Menu[menu].Item {
				//fmt.Println("item1: ", itemName1, ", item2: ", itemName2)
				if strings.Compare(itemName1, itemName2) == 0 {
					itemList = append(itemList, itemName1)
					if len(item.Description) > 0 {
						descriptionList = append(descriptionList, item.Description)
					} else {
						descriptionList = append(descriptionList, abbreviation.NA)
					}

					break
				}
			}
		}
		return itemList, descriptionList
	}
}

func GetMenuListOfRoleList(roleList []string) ([]string, [][]string, [][]string) {
	cf := config.RBACConf()
	if cf == nil {
		log.Error("GetMenuListOfRoleList failure, cf == nil")
		return nil, nil, nil
	}

	itemMap := map[string][]string{}
	descMap := map[string][]string{}
	menuList := []string{}
	itemListSet := make([][]string, len(cf.Menu))
	descListSet := make([][]string, len(cf.Menu))

	for _, role := range roleList {
		if role == role2.Administrator {
			return GetMenuList()
		} else {
			for menuName, _ := range cf.RoleMenu[role].Menu {
				_, exist1 := itemMap[menuName]
				_, exist2 := descMap[menuName]

				if !exist1 && !exist2 {
					itemMap[menuName] = []string{}
					descMap[menuName] = []string{}
				}

				itemList, descList := GetItemListOfRoleMenu(role, menuName)
				if len(itemList) != len(descList) {
					continue
				}

				for i := 0; i < len(itemList); i++ {
					if !slices.Contains(itemMap[menuName], itemList[i]) {
						itemMap[menuName] = append(itemMap[menuName], itemList[i])
						if len(descList[i]) > 0 {
							descMap[menuName] = append(descMap[menuName], descList[i])
						} else {
							descMap[menuName] = append(descMap[menuName], abbreviation.NA)
						}

					}
				}
			}
		}
	}

	index := 0
	for k, _ := range itemMap {
		if len(itemMap[k]) > 0 && len(descMap[k]) > 0 && len(itemMap[k]) == len(descMap[k]) {
			menuList = append(menuList, k)
			itemListSet[index] = itemMap[k]
			descListSet[index] = descMap[k]
			index++
		}
	}

	return menuList, itemListSet, descListSet
}

func GetFieldListOfTable(table string) ([]string, []string) {
	cf := config.RBACConf()
	if cf == nil {
		log.Error("GetFieldListOfTable failure, cf == nil")
		return nil, nil
	}
	if _, exist := cf.Table[table]; !exist {
		log.Error(fmt.Sprintf("GetFieldListOfTable %s not in Table", table))
		return nil, nil
	}
	fieldList := []string{}
	descriptionList := []string{}
	for fieldName, item := range cf.Table[table].Field {
		fieldList = append(fieldList, fieldName)
		if len(item.Description) > 0 {
			descriptionList = append(descriptionList, item.Description)
		} else {
			descriptionList = append(descriptionList, abbreviation.NA)
		}
	}
	return fieldList, descriptionList
}

func GetFieldList() ([]string, [][]string, [][]string) {
	cf := config.RBACConf()
	if cf == nil {
		log.Error("GetFieldList failure, cf == nil")
		return nil, nil, nil
	}

	tableList := []string{}
	fieldListSet := make([][]string, len(cf.Table))
	descListSet := make([][]string, len(cf.Table))

	index := 0
	for tableName, field := range cf.Table {
		if len(field.Field) > 0 {
			tableList = append(tableList, tableName)
			fieldListSet[index], descListSet[index] = GetFieldListOfTable(tableName)
			index++
		}
	}
	return tableList, fieldListSet, descListSet
}

func GetFieldListOfRoleField(role, table string) ([]string, []string) {
	cf := config.RBACConf()
	if cf == nil {
		log.Error("GetFieldListOfRoleField failure, cf == nil")
		return nil, nil
	}

	if role == role2.Administrator {
		return GetFieldListOfTable(table)
	} else {
		//check if role in both Role and RoleMenu
		if !HasRole(role) {
			log.ErrorF("GetFieldListOfRoleField failure, %v not in Role", role)
			return nil, nil
		}
		if _, exist := cf.RoleMenu[role]; !exist {
			log.ErrorF("GetFieldListOfRoleField failure, role(%v) not in RoleTable", role)
			return nil, nil
		}
		if _, exist := cf.Table[table]; !exist {
			log.Warn(fmt.Sprintf("GetFieldListOfRoleField %v not in Table", table))
			return nil, nil
		}
		if _, exist := cf.RoleField[role].Table[table]; !exist {
			log.Warn(fmt.Sprintf("GetFieldListOfRoleField %s not in RoleField", table))
			return nil, nil
		}
		fieldList := []string{}
		descriptionList := []string{}
		for _, fieldName1 := range cf.RoleField[role].Table[table].FieldList {
			for fieldName2, field := range cf.Table[table].Field {
				if strings.Compare(fieldName1, fieldName2) == 0 {
					fieldList = append(fieldList, fieldName1)
					if len(field.Description) > 0 {
						descriptionList = append(descriptionList, field.Description)
					} else {
						descriptionList = append(descriptionList, abbreviation.NA)
					}

					break
				}
			}
		}
		return fieldList, descriptionList
	}
}

func GetFieldListOfRoleList(roleList []string) ([]string, [][]string, [][]string) {
	cf := config.RBACConf()
	if cf == nil {
		log.Error("GetFieldListOfRoleList failure, cf == nil")
		return nil, nil, nil
	}

	fieldMap := map[string][]string{}
	descMap := map[string][]string{}
	tableList := []string{}
	fieldListSet := make([][]string, len(cf.Table))
	descListSet := make([][]string, len(cf.Table))

	for _, role := range roleList {
		if role == role2.Administrator {
			return GetFieldList()
		} else {
			for tableName, _ := range cf.RoleField[role].Table {
				_, exist1 := fieldMap[tableName]
				_, exist2 := descMap[tableName]

				if !exist1 && !exist2 {
					fieldMap[tableName] = []string{}
					descMap[tableName] = []string{}
				}

				fieldList, descList := GetFieldListOfRoleField(role, tableName)
				for i := 0; i < len(fieldList); i++ {
					if !slices.Contains(fieldMap[tableName], fieldList[i]) {
						fieldMap[tableName] = append(fieldMap[tableName], fieldList[i])
						if len(descList[i]) > 0 {
							descMap[tableName] = append(descMap[tableName], descList[i])
						} else {
							descMap[tableName] = append(descMap[tableName], abbreviation.NA)
						}
					}
				}
			}
		}
	}

	index := 0
	for k, _ := range fieldMap {
		if len(fieldMap[k]) > 0 && len(descMap[k]) > 0 && len(fieldMap[k]) == len(descMap[k]) {
			tableList = append(tableList, k)
			fieldListSet[index] = fieldMap[k]
			descListSet[index] = descMap[k]
			index++
		}
	}

	return tableList, fieldListSet, descListSet
}

func GetPermissionList() ([]string, []string, []string, []string) {
	cf := config.RBACConf()
	if cf == nil {
		log.Error("GetPermissionList failure, cf == nil")
		return nil, nil, nil, nil
	}
	nameList := []string{}
	majorList := []string{}
	minorList := []string{}
	descList := []string{}
	for permissionName, permission := range cf.Permission {
		if !slices.Contains(nameList, permissionName) {
			nameList = append(nameList, permissionName)
			if permission.Major > 0 {
				majorList = append(majorList, cast.ToString(permission.Major))
			} else {
				majorList = append(majorList, abbreviation.NA)
			}
			if permission.Minor > 0 {
				minorList = append(minorList, cast.ToString(permission.Minor))
			} else {
				minorList = append(minorList, abbreviation.NA)
			}
			if len(permission.Description) > 0 {
				descList = append(descList, cast.ToString(permission.Description))
			} else {
				descList = append(descList, abbreviation.NA)
			}
		}
	}
	return nameList, majorList, minorList, descList
}

// Notice: this call is without permission check, please check if has permission first
func GetPermissionByName(name string) (string, string, string, string) {
	major := ""
	minor := ""
	desc := ""
	cf := config.RBACConf()
	if cf == nil {
		log.Error("GetPermissionByName failure, cf == nil")
		return "", major, minor, desc
	}
	if len(name) <= 0 {
		return "", major, minor, desc
	}

	if permission, exist := cf.Permission[name]; exist {
		if permission.Major > 0 {
			major = cast.ToString(permission.Major)
		} else {
			major = abbreviation.NA
		}
		if permission.Minor > 0 {
			minor = cast.ToString(permission.Minor)
		} else {
			minor = abbreviation.NA
		}
		if len(permission.Description) > 0 {
			desc = permission.Description
		} else {
			desc = abbreviation.NA
		}
	} else {
		name = ""
	}
	return name, major, minor, desc
}

func GetPermissionListOfRoleList(roleList []string) ([]string, []string, []string, []string) {
	cf := config.RBACConf()
	if cf == nil {
		log.Error("GetPermissionListOfRoleList failure, cf == nil")
		return nil, nil, nil, nil
	}

	nameList := []string{}
	majorList := []string{}
	minorList := []string{}
	descList := []string{}

	for _, role := range roleList {
		if role == role2.Administrator {
			return GetPermissionList()
		} else {
			// check if role in both Role and RolePermission
			if !HasRole(role) {
				log.WarnF("GetPermissionListOfRoleList Warning: role(%v) not in Role", role)
				continue
			}
			if _, exist := cf.RolePermission[role]; !exist {
				log.WarnF("GetPermissionListOfRoleList failure, role(%v) not in RolePermission", role)
				continue
			}

			for _, permissionName := range cf.RolePermission[role].PermissionList {
				// check if permission in both RolePermission and Permission
				if permission, exist := cf.Permission[permissionName]; exist {
					if !slices.Contains(nameList, permissionName) {
						nameList = append(nameList, permissionName)
						if permission.Major > 0 {
							majorList = append(majorList, cast.ToString(permission.Major))
						} else {
							majorList = append(majorList, abbreviation.NA)
						}
						if permission.Minor > 0 {
							minorList = append(minorList, cast.ToString(permission.Minor))
						} else {
							minorList = append(minorList, abbreviation.NA)
						}
						if len(permission.Description) > 0 {
							descList = append(descList, cast.ToString(permission.Description))
						} else {
							descList = append(descList, abbreviation.NA)
						}
					}
				} else {
					log.WarnF("GetRolePermissionListOfRole failure, RolePermission.Role(%v).Permission(%v) not in Permission.Permission", role, permission)
				}
			}
		}
	}
	return nameList, majorList, minorList, descList
}

//	func GetFieldList() []byte {
//		cf := config.RBACConf()
//		if cf == nil {
//			log.Error("GetFieldList failure, cf == nil")
//			return nil
//		}
//		js := json2.New()
//		for table, temp := range cf.Table {
//			fieldList := []string{}
//			fieldDescList := []string{}
//			for fieldName, field := range temp.Field {
//				fieldList = append(fieldList, fieldName)
//				fieldDescList = append(fieldDescList, field.Description)
//			}
//			js.Set(table, fieldList)
//		}
//		bytes, err := js.Encode()
//		if err != nil {
//			log.Error("GetFieldList failure, js.Bytes err: ", err.Error())
//			return nil
//		}
//		return bytes
//	}
//
//	func GetFieldListOfTable(table string) ([]string, []string) {
//		cf := config.RBACConf()
//		if cf == nil {
//			log.Error("GetFieldListOfTable failure, cf == nil")
//			return nil, nil
//		}
//		if _, exist := cf.Table[table]; !exist {
//			log.Error(fmt.Sprintf("GetFieldListOfTable %s not in Table", table))
//			return nil, nil
//		}
//		fieldList := []string{}
//		descList := []string{}
//		for fieldName, field := range cf.Table[table].Field {
//			fieldList = append(fieldList, fieldName)
//			descList = append(descList, field.Description)
//		}
//		return fieldList, descList
//	}
//
//	func GetFieldListOfRoleField(role, table string) ([]string, []string) {
//		cf := config.RBACConf()
//		if cf == nil {
//			log.Error("GetFieldListOfRoleField failure, cf == nil")
//			return nil, nil
//		}
//
//		if role == role2.Administrator {
//			return GetFieldListOfTable(table)
//		} else {
//			//check if role in both Role and RoleField
//			if !HasRole(role) {
//				log.ErrorF("GetFieldListOfRoleField failure, %v not in Role", role)
//				return nil, nil
//			}
//			if _, exist := cf.RoleField[role]; !exist {
//				log.ErrorF("GetFieldListOfRoleField failure, role(%v) not in RoleField", role)
//				return nil, nil
//			}
//			if _, exist := cf.Table[table]; !exist {
//				log.Warn(fmt.Sprintf("GetFieldListOfRoleField %v not in Table", table))
//				return nil, nil
//			}
//			if _, exist := cf.RoleField[role].Table[table]; !exist {
//				log.Warn(fmt.Sprintf("GetFieldListOfRoleField %s not in RoleField", table))
//				return nil, nil
//			}
//			fieldList := []string{}
//			descList := []string{}
//			for _, fieldName1 := range cf.RoleField[role].Table[table].FieldList {
//				for fieldName2, field := range cf.Table[table].Field {
//					if fieldName1 == fieldName2 {
//						fieldList = append(fieldList, fieldName1)
//						descList = append(descList, field.Description)
//						break
//					}
//				}
//			}
//			return fieldList, descList
//		}
//	}
//
//	func GetFieldListOfRoleList(roleList []string) []byte {
//		cf := config.RBACConf()
//		if cf == nil {
//			log.Error("GetFieldListOfRoleList failure, cf == nil")
//			return nil
//		}
//
//		js := json2.New()
//		temp := map[string]set.Set{}
//
//		for _, role := range roleList {
//			if role == role2.Administrator {
//				return GetFieldList()
//			} else {
//				for tableName, _ := range cf.RoleField[role].Table {
//					if _, exist := temp[tableName]; !exist {
//						temp[tableName] = set2.NewSet()
//					}
//
//					fieldList, descList := GetFieldListOfRoleField(role, tableName)
//					if len(fieldList) != len(descList) {
//						continue
//					}
//
//					for i := 0; i < len(fieldList); i++ {
//						temp[tableName].SAdd(fieldList[i])
//					}
//				}
//			}
//		}
//
//		for k, v := range temp {
//			js.Set(k, v.SMembers())
//		}
//		bytes, err := js.Encode()
//		if err != nil {
//			log.Error("GetFieldListOfRoleList failure, js.Bytes err: ", err.Error())
//			return nil
//		}
//		return bytes
//	}
//
//	func GetFieldListOfTableOfRoleList(roleList []string, table string) ([]string, []string) {
//		cf := config.RBACConf()
//		if cf == nil {
//			log.Error("GetFieldListOfTableOfRoleList failure, cf == nil")
//			return nil, nil
//		}
//
//		js := json2.New()
//		temp := map[string]set.Set{}
//
//		for _, role := range roleList {
//			if role == role2.Administrator {
//				return GetFieldList()
//			} else {
//				for tableName, _ := range cf.RoleField[role].Table {
//					if _, exist := temp[tableName]; !exist {
//						temp[tableName] = set2.NewSet()
//					}
//
//					fieldList, descList := GetFieldListOfRoleField(role, tableName)
//					if len(fieldList) != len(descList) {
//						continue
//					}
//
//					for i := 0; i < len(fieldList); i++ {
//						temp[tableName].SAdd(fieldList[i])
//					}
//				}
//			}
//		}
//
//		for k, v := range temp {
//			js.Set(k, v.SMembers())
//		}
//		bytes, err := js.Encode()
//		if err != nil {
//			log.Error("GetFieldListOfRoleList failure, js.Bytes err: ", err.Error())
//			return nil
//		}
//		return bytes
//	}
//
// //
// //func GetFieldListOfRoleList(roleList []string) []byte {
// //	cf := config.RBACConf()
// //	if cf == nil {
// //		log.Error("GetFieldListOfRoleList failure, cf == nil")
// //		return nil
// //	}
// //
// //	js := json2.New()
// //	temp := map[string]set.Set{}
// //
// //	for _, role := range roleList {
// //		if role == role2.Administrator {
// //			return GetFieldList()
// //		} else {
// //			for tableName, _ := range cf.RoleField[role].Table {
// //				if _, exist := temp[tableName]; !exist {
// //					temp[tableName] = set2.NewSet()
// //				}
// //
// //				fieldList, descList := GetFieldListOfRoleField(role, tableName)
// //				if len(fieldList) != len(descList) {
// //					continue
// //				}
// //
// //				for i := 0; i < len(fieldList); i++ {
// //					//desc := "null"
// //					//if len(descList[i]) > 0 {
// //					//	desc = descList[i]
// //					//}
// //					//temp[menuName].SAdd(fmt.Sprintf("%s-%s", itemList[i], desc))
// //					temp[tableName].SAdd(fieldList[i])
// //				}
// //			}
// //		}
// //	}
// //
// //	for k, v := range temp {
// //		//memberList := v.SMembers()
// //		//for _, v2 := range memberList {
// //		//	temp := strings.Split(v2, "-")
// //		//	if len(temp) == 2 {
// //		//		fmt.Println("k: ", k, "temp: ", temp)
// //		//		js.SetPath([]string{k, "Item"}, temp[0])
// //		//		js.SetPath([]string{k, "Description"}, temp[1])
// //		//	}
// //		//}
// //		js.Set(k, v.SMembers())
// //	}
// //	bytes, err := js.Encode()
// //	if err != nil {
// //		log.Error("GetFieldListOfRoleList failure, js.Bytes err: ", err.Error())
// //		return nil
// //	}
// //	return bytes
// //}
//
// //func GetAttributeList() []byte {
// //	cf := config.RBACConf()
// //	if cf == nil {
// //		log.Error("GetAttribute failure, cf == nil")
// //		return nil
// //	}
// //	js := json2.New()
// //	for table, temp := range cf.Table {
// //		attributes := []string{}
// //		for _, attribute := range temp.Attribute {
// //			attributes = append(attributes, attribute)
// //		}
// //		js.Set(table, attributes)
// //	}
// //	bytes, err := js.Encode()
// //	if err != nil {
// //		log.Error("GetAttribute failure, js.Bytes err: ", err.Error())
// //		return nil
// //	}
// //	return bytes
// //}
//
// //func GetMenuListOfRoleList(roleList []string) []byte {
// //	cf := config.RBACConf()
// //	if cf == nil {
// //		log.Error("GetMenuListOfRoleList failure, cf == nil")
// //		return nil
// //	}
// //
// //	temp := map[string]set2.Set{}
// //	for _, v := range roleList {
// //		if v == role2.Administrator {
// //			return GetMenuList()
// //		}
// //		set.NewSet()
// //
// //	}
// //}
// //
// //func GetItemListOfRole(role, menu string) []string {
// //	cf := config.RBACConf()
// //	if cf == nil {
// //		log.Error("GetMenuListOfRole failure, cf == nil")
// //		return nil
// //	}
// //
// //	if role == role2.Administrator {
// //		return GetMenuList()
// //	} else {
// //		// check if role in both Role and RoleMenu
// //		if !HasRole(role) {
// //			log.ErrorF("GetMenuListOfRole failure, role(%v) not in Role.Role", role)
// //			return nil
// //		}
// //		if _, exist := cf.RoleMenu[role]; !exist {
// //			log.ErrorF("GetMenuListOfRole failure, role(%v) not in RoleMenu.Role", role)
// //			return nil
// //		}
// //		js := json2.New()
// //		for menuName, temp := range cf.RoleMenu[role].Menu {
// //			// check if menu in both RoleMenu and Menu
// //			if _, exist := cf.Menu[menuName]; !exist {
// //				log.ErrorF("GetMenuListOfRole failure, RoleMenu.Role(%v).Menu(%v) not in Menu.Menu", role, menu)
// //				continue
// //			}
// //			items := []string{}
// //			for _, item := range temp.Item {
// //				has := false
// //				for _, tempItem := range cf.Menu[menuName].Item {
// //					// check if item in both Menu and RoleMenu
// //					if strings.Compare(tempItem, item) == 0 {
// //						has = true
// //						break
// //					}
// //				}
// //				if has {
// //					items = append(items, item)
// //				} else {
// //					log.ErrorF("GetMenuListOfRole failure, RoleMenu(%v).Menu(%v).Item(%v) not in Menu.Item", role, menu, item)
// //				}
// //			}
// //			js.Set(menuName, items)
// //		}
// //		bytes, err := js.Encode()
// //		if err != nil {
// //			log.Error("GetMenuListOfRole failure, js.Bytes err: ", err.Error())
// //			return nil
// //		}
// //		return bytes
// //	}
// //}
//
//	func GetRoleList() []string {
//		roleList := []string{}
//		cf := config.RBACConf()
//		if cf == nil {
//			log.Error("GetRoleList failure, cf == nil")
//			return roleList
//		}
//		for role, _ := range cf.Role {
//			roleList = append(roleList, role)
//		}
//		return roleList
//	}
//
//	func GetPermissionList() []int {
//		permissionList := []int{}
//		cf := config.RBACConf()
//		if cf == nil {
//			log.Error("GetPermissionList failure, cf == nil")
//			return permissionList
//		}
//		for _, permission := range cf.Permission {
//			permissionList = append(permissionList, permission.Minor)
//		}
//		return permissionList
//	}
//
//	func GetPermissionListOfRole(role string) []int {
//		permissionList := []int{}
//
//		cf := config.RBACConf()
//		if cf == nil {
//			log.Error("GetRolePermissionListOfRole failure, cf == nil")
//			return permissionList
//		}
//
//		if role == role2.Administrator {
//			return GetPermissionList()
//		} else {
//			// check if role in both Role and RolePermission
//			if !HasRole(role) {
//				log.ErrorF("GetPermissionListOfRole failure, role(%v) not in Role.Role", role)
//				return permissionList
//			}
//			if _, exist := cf.RolePermission[role]; !exist {
//				log.ErrorF("GetRolePermissionListOfRole failure, role(%v) not in RolePermission.Role", role)
//				return permissionList
//			}
//			for _, permissionName := range cf.RolePermission[role].PermissionList {
//				// check if permission in both RolePermission and Permission
//				if permission, exist := cf.Permission[permissionName]; exist {
//					permissionList = append(permissionList, permission.Minor)
//				} else {
//					log.ErrorF("GetRolePermissionListOfRole failure, RolePermission.Role(%v).Permission(%v) not in Permission.Permission", role, permission)
//				}
//			}
//			return permissionList
//		}
//	}
//
// //
// //func GetAttributeListOfRole(role string) []byte {
// //	cf := config.RBACConf()
// //	if cf == nil {
// //		log.Error("GetAttributeListOfRole failure, cf == nil")
// //		return nil
// //	}
// //
// //	if role == role2.Administrator {
// //		return GetAttributeList()
// //	} else {
// //		// check if role in both Role and RoleAttribute
// //		if !HasRole(role) {
// //			log.ErrorF("GetMenuListOfRole failure, role(%v) not in Role.Role", role)
// //			return nil
// //		}
// //		if _, exist := cf.RoleAttribute[role]; !exist {
// //			log.ErrorF("GetAttributeListOfRole failure, role(%v) not in RoleAttribute.Role", role)
// //			return nil
// //		}
// //		js := json2.New()
// //		for table, temp := range cf.RoleAttribute[role].Table {
// //			attributes := []string{}
// //			// check if table in both Table and RoleAttribute
// //			if _, exist := cf.Table[table]; !exist {
// //				log.ErrorF("GetMenuListOfRole failure, RoleAttribute.Role(%v).Table(%v) not in Table.Table", role, table)
// //				continue
// //			}
// //			for _, attribute := range temp.Attribute {
// //				has := false
// //				for _, tempAttribute := range cf.Table[table].Attribute {
// //					// check if attribute in both Table and RoleAttribute
// //					if strings.Compare(tempAttribute, attribute) == 0 {
// //						has = true
// //						break
// //					}
// //				}
// //				if has {
// //					attributes = append(attributes, attribute)
// //				} else {
// //					log.ErrorF("GetAttributeListOfRole failure, RoleAttribute(%v).Table(%v).Attribute(%v) not in Table.Attribute", role, table, attribute)
// //				}
// //			}
// //			js.Set(table, attributes)
// //		}
// //		bytes, err := js.Encode()
// //		if err != nil {
// //			log.Error("GetAttributeListOfRole failure, js.Bytes err: ", err.Error())
// //			return nil
// //		}
// //		return bytes
// //	}
// //}

//
//func HasMenu(role, menu, item string) bool {
//	cf := config.RBACConf()
//	if cf == nil {
//		log.Error("HasPermission failure, cf == nil")
//		return false
//	}
//
//	if role == role2.Administrator {
//		return true
//	} else {
//		// check if role in both Role and RoleMenu
//		if !HasRole(role) {
//			log.ErrorF("HasMenu failure, role(%v) not in Role", role)
//			return false
//		}
//		if _, exist := cf.RoleMenu[role]; !exist {
//			log.ErrorF("HasMenu failure, role(%v) not in RoleMenu.Role", role)
//			return false
//		}
//		// check if menu in both Menu and RoleMenu
//		if _, exist := cf.Menu[menu]; !exist {
//			log.ErrorF("HasMenu failure, menu(%v) not in Menu", menu)
//			return false
//		}
//		if _, exist := cf.RoleMenu[role].Menu[menu]; !exist {
//			log.ErrorF("HasMenu failure, menu(%v) not in RoleMenu.Role(%v).Menu", menu, role)
//			return false
//		}
//		// check if item in both Menu and RoleMenu
//		has := false
//		for _, tempItem := range cf.Menu[menu].Item {
//			if strings.Compare(tempItem, item) == 0 {
//				has = true
//			}
//		}
//		if !has {
//			log.ErrorF("HasMenu failure, item(%v) not in Menu.(%v).Item", item, menu)
//		}
//		has = false
//		for _, tempItem := range cf.RoleMenu[role].Menu[menu].Item {
//			if strings.Compare(tempItem, item) == 0 {
//				has = true
//			}
//		}
//		if !has {
//			log.ErrorF("HasMenu failure, item(%v) not in RoleMenu.Role(%v).Menu(%v).Item", item, role, menu)
//			return false
//		}
//
//		return true
//	}
//}
//
//func HasAttribute(role, table, attribute string) bool {
//	cf := config.RBACConf()
//	if cf == nil {
//		log.Error("HasPermission failure, cf == nil")
//		return false
//	}
//
//	if role == role2.Administrator {
//		return true
//	} else {
//		// check if role in both Role and RoleAttribute
//		if !HasRole(role) {
//			log.ErrorF("HasAttribute failure, role(%v) not in Role", role)
//			return false
//		}
//		if _, exist := cf.RoleAttribute[role]; !exist {
//			log.ErrorF("HasAttribute failure, role(%v) not in RoleAttribute.Role", role)
//			return false
//		}
//		// check if table in both Table and RoleAttribute
//		if _, exist := cf.Table[table]; !exist {
//			log.ErrorF("HasAttribute failure, table(%v) not in Table", table)
//			return false
//		}
//		if _, exist := cf.RoleAttribute[role].Table[table]; !exist {
//			log.ErrorF("HasAttribute failure, table(%v) not in RoleAttribute.Role(%v).Table", table, role)
//			return false
//		}
//		// check if attribute in both Table and RoleAttribute
//		has := false
//		for _, tempAttribute := range cf.Table[table].Attribute {
//			if strings.Compare(tempAttribute, attribute) == 0 {
//				has = true
//			}
//		}
//		if !has {
//			log.ErrorF("HasAttribute failure, attribute(%v) not in Table(%v).Attribute", attribute, table)
//			return false
//		}
//		has = false
//		for _, tempAttribute := range cf.RoleAttribute[role].Table[table].Attribute {
//			if strings.Compare(attribute, tempAttribute) == 0 {
//				has = true
//			}
//		}
//		if !has {
//			log.ErrorF("HasAttribute failure, attribute(%v) not in RoleAttribute.Role(%v).Table(%v).Attribute", attribute, role, table)
//			return false
//		}
//
//		return true
//	}
//}
