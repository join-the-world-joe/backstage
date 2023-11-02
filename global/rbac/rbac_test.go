package rbac

import (
	role2 "backstage/common/macro/role"
	"backstage/diagnostic"
	"backstage/utils/json"
	"testing"
)

func TestGetPermissionByName(t *testing.T) {
	diagnostic.SetupLogger()
	diagnostic.SetupRBAC()
	name := "FetchMenuListOfCondition"
	name, major, minor, desc := GetPermissionByName(name)
	t.Log("Name: ", name)
	t.Log("Major: ", major)
	t.Log("Minor: ", minor)
	t.Log("Description: ", desc)
}

func TestHasPermission(t *testing.T) {
	diagnostic.SetupLogger()
	diagnostic.SetupRBAC()

	role := "Administrator"
	//role = role2.Administrator
	minor := 5

	b := HasPermission(role, minor)
	if b {
		t.Logf("b(%v)", b)
	} else {
		t.Logf("b(%v)", b)
	}
}

func TestGetItemListOfMenu(t *testing.T) {
	menu := "Admission"
	diagnostic.SetupLogger()
	diagnostic.SetupRBAC()
	itemList, descList := GetItemListOfMenu(menu)
	t.Log("Item List: ", itemList)
	t.Log("Desc List: ", descList)
}

func TestGetMenuList(t *testing.T) {
	diagnostic.SetupLogger()
	diagnostic.SetupRBAC()
	menuList, itemList, descList := GetMenuList()
	t.Log("Menu List: ", menuList)
	t.Log("Item List: ", itemList)
	t.Log("Description List: ", descList)

	js := json.New()
	for i, length := 0, len(menuList); i < length; i++ {
		js.SetPath([]string{menuList[i], "Item"}, itemList[i])
		js.SetPath([]string{menuList[i], "Description"}, descList[i])
	}

	bytes, err := js.Encode()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(bytes))
}

func TestGetItemListOfRoleMenu(t *testing.T) {
	role := "Manager"
	//role := role2.Administrator
	menu := "Admission"
	diagnostic.SetupLogger()
	diagnostic.SetupRBAC()
	itemList, descList := GetItemListOfRoleMenu(role, menu)
	t.Log("Item List: ", itemList)
	t.Log("Desc List: ", descList)
}

func TestGetMenuListOfRoleList(t *testing.T) {
	roleList := []string{role2.Administrator}
	//roleList := []string{"Role1", "Role2"}
	//roleList := []string{"Role2"}
	diagnostic.SetupLogger()
	diagnostic.SetupRBAC()
	//t.Log(string(GetMenuListOfRoleList(roleList)))
	menuList, itemList, descList := GetMenuListOfRoleList(roleList)
	t.Log("Menu List: ", menuList)
	t.Log("Item List: ", itemList)
	t.Log("Description List: ", descList)

	js := json.New()
	for i, length := 0, len(menuList); i < length; i++ {
		js.SetPath([]string{menuList[i], "Item"}, itemList[i])
		js.SetPath([]string{menuList[i], "Description"}, descList[i])
	}

	bytes, err := js.Encode()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(bytes))
}

func TestGetFieldListOfTable(t *testing.T) {
	table := "user"
	diagnostic.SetupLogger()
	diagnostic.SetupRBAC()
	fieldList, descList := GetFieldListOfTable(table)
	t.Log("Field List: ", fieldList)
	t.Log("Desc List: ", descList)
}

func TestGetFieldList(t *testing.T) {
	diagnostic.SetupLogger()
	diagnostic.SetupRBAC()
	tableList, fieldList, descList := GetFieldList()
	t.Log("Table List: ", tableList)
	t.Log("Field List: ", fieldList)
	t.Log("Description List: ", descList)

	js := json.New()
	for i, length := 0, len(tableList); i < length; i++ {
		js.SetPath([]string{tableList[i], "Item"}, fieldList[i])
		js.SetPath([]string{tableList[i], "Description"}, descList[i])
	}

	bytes, err := js.Encode()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(bytes))
}

func TestGetFieldListOfRoleField(t *testing.T) {
	role := "Manager"
	//role := role2.Administrator
	table := "user"
	diagnostic.SetupLogger()
	diagnostic.SetupRBAC()
	fieldList, descList := GetFieldListOfRoleField(role, table)
	t.Log("Field List: ", fieldList)
	t.Log("Desc List: ", descList)
}

func TestGetFieldListOfRoleList(t *testing.T) {
	//roleList := []string{"Role1", "Role2", "Manager"}
	//roleList := []string{"Role1", "Role2"}
	roleList := []string{"Manager"}
	diagnostic.SetupLogger()
	diagnostic.SetupRBAC()
	//t.Log(string(GetMenuListOfRoleList(roleList)))
	tableList, fieldList, descList := GetFieldListOfRoleList(roleList)
	t.Log("Table List: ", tableList)
	t.Log("Field List: ", fieldList)
	t.Log("Description List: ", descList)

	js := json.New()
	for i, length := 0, len(tableList); i < length; i++ {
		js.SetPath([]string{"table_list", tableList[i], "Field"}, fieldList[i])
		js.SetPath([]string{"table_list", tableList[i], "Description"}, descList[i])
	}

	bytes, err := js.Encode()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(bytes))
}

func TestGetPermissionList(t *testing.T) {
	diagnostic.SetupLogger()
	diagnostic.SetupRBAC()
	nameList, majorList, minorList, descList := GetPermissionList()
	t.Log("nameList List: ", nameList)
	t.Log("majorList List: ", majorList)
	t.Log("minorList List: ", minorList)
	t.Log("desc List: ", descList)
}

func TestGetPermissionListOfRoleList(t *testing.T) {
	roleList := []string{"Manager", "Role1"}
	diagnostic.SetupLogger()
	diagnostic.SetupRBAC()
	nameList, majorList, minorList, descList := GetPermissionListOfRoleList(roleList)
	t.Log("nameList List: ", nameList)
	t.Log("majorList List: ", majorList)
	t.Log("minorList List: ", minorList)
	t.Log("desc List: ", descList)

	js := json.New()
	for i, length := 0, len(nameList); i < length; i++ {
		js.SetPath([]string{nameList[i], "Major"}, majorList[i])
		js.SetPath([]string{nameList[i], "Minor"}, minorList[i])
		js.SetPath([]string{nameList[i], "Description"}, descList[i])
	}

	bytes, err := js.Encode()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(bytes))
}

func TestGetTopLevelOfRoleList(t *testing.T) {
	diagnostic.SetupLogger()
	diagnostic.SetupRBAC()

	roleList := []string{"RD_Manager"}
	t.Log("rank: ", GetTopRankOfRoleList(roleList))
}

func TestGetRoleListBERank(t *testing.T) {
	diagnostic.SetupLogger()
	diagnostic.SetupRBAC()

	roleList := []string{role2.RDManager}
	rank := GetTopRankOfRoleList(roleList)
	t.Log("Role List: ", GetRoleListLERank(rank))
}

//
//func TestGetFieldList(t *testing.T) {
//	diagnostic.SetupLogger()
//	diagnostic.SetupRBAC()
//	t.Log("Field List: ", string(GetFieldList()))
//}
//
//func TestGetFieldListOfTable(t *testing.T) {
//	table := "user"
//	diagnostic.SetupLogger()
//	diagnostic.SetupRBAC()
//	itemList, descList := GetFieldListOfTable(table)
//	t.Log("Field List: ", itemList)
//	t.Log("Desc List: ", descList)
//}
//
//func TestGetFieldListOfRoleField(t *testing.T) {
//	//role := "Manager"
//	table := "user"
//	role := role2.Administrator
//	diagnostic.SetupLogger()
//	diagnostic.SetupRBAC()
//	fieldList, descList := GetFieldListOfRoleField(role, table)
//	t.Log("Field List: ", fieldList)
//	t.Log("Desc List: ", descList)
//}
//
//func TestGetFieldListOfRoleList(t *testing.T) {
//	roleList := []string{"Role1", "Role2", role2.Administrator}
//	diagnostic.SetupLogger()
//	diagnostic.SetupRBAC()
//	t.Log(string(GetFieldListOfRoleList(roleList)))
//}
//
//func TestGetRoleList(t *testing.T) {
//	diagnostic.SetupLogger()
//	diagnostic.SetupRBAC()
//	t.Log("Role List: ", GetRoleList())
//}
//
//func TestGetPermissionList(t *testing.T) {
//	diagnostic.SetupLogger()
//	diagnostic.SetupRBAC()
//	t.Log("Permission List: ", GetPermissionList())
//}
//
//func TestGetPermissionListOfRole(t *testing.T) {
//	diagnostic.SetupLogger()
//	diagnostic.SetupRBAC()
//	roleList := []string{role2.Administrator, "Manager"}
//	for _, role := range roleList {
//		t.Logf("Permission List(%v): %v", role, GetPermissionListOfRole(role))
//	}
//}
//
////func TestGetMenuListOfRole(t *testing.T) {
////	diagnostic.SetupLogger()
////	diagnostic.SetupRBAC()
////	roleList := []string{role2.Administrator, "Manager"}
////	for _, role := range roleList {
////		t.Logf("Menu List(%v): %v", role, string(GetMenuListOfRole(role)))
////	}
////}
////
////func TestGetAttributeListOfRole(t *testing.T) {
////	diagnostic.SetupLogger()
////	diagnostic.SetupRBAC()
////	roleList := []string{role2.Administrator, "Manager"}
////	for _, role := range roleList {
////		t.Logf("Attribute List(%v): %v", role, string(GetAttributeListOfRole(role)))
////	}
////}
//
//func TestHasRole(t *testing.T) {
//	diagnostic.SetupLogger()
//	diagnostic.SetupRBAC()
//
//	role := "Manager"
//	b := HasRole(role)
//	if b {
//		t.Logf("b(%v)", b)
//	} else {
//		t.Logf("b(%v)", b)
//	}
//}
//

//func TestHasMenu(t *testing.T) {
//	diagnostic.SetupLogger()
//	diagnostic.SetupRBAC()
//
//	role := "Manager"
//	menu := "Admission"
//	item := "Track"
//
//	b := HasMenu(role, menu, item)
//	if b {
//		t.Logf("b(%v)", b)
//	} else {
//		t.Logf("b(%v)", b)
//	}
//}
//
//func TestHasAttribute(t *testing.T) {
//	diagnostic.SetupLogger()
//	diagnostic.SetupRBAC()
//
//	//role := "Manager"
//	role := role2.Administrator
//	table := "user"
//	//name := "country"
//	name := "account"
//
//	b := HasAttribute(role, table, name)
//	if b {
//		t.Logf("b(%v)", b)
//	} else {
//		t.Logf("b(%v)", b)
//	}
//}
