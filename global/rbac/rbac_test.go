package rbac

import (
	role2 "go-micro-framework/common/macro/role"
	"go-micro-framework/diagnostic"
	"testing"
)

func TestGetRoleList(t *testing.T) {
	diagnostic.SetupLogger()
	diagnostic.SetupRBAC()
	t.Log("Role List: ", GetRoleList())
}

func TestGetPermissionList(t *testing.T) {
	diagnostic.SetupLogger()
	diagnostic.SetupRBAC()
	t.Log("Permission List: ", GetPermissionList())
}

func TestGetMenuList(t *testing.T) {
	diagnostic.SetupLogger()
	diagnostic.SetupRBAC()
	t.Log("Menu List: ", string(GetMenuList()))
}

func TestGetAttribute(t *testing.T) {
	diagnostic.SetupLogger()
	diagnostic.SetupRBAC()
	t.Log("Attribute List: ", string(GetAttributeList()))
}

func TestGetPermissionListOfRole(t *testing.T) {
	diagnostic.SetupLogger()
	diagnostic.SetupRBAC()
	roleList := []string{role2.Administrator, "Manager"}
	for _, role := range roleList {
		t.Logf("Permission List(%v): %v", role, GetPermissionListOfRole(role))
	}
}

func TestGetMenuListOfRole(t *testing.T) {
	diagnostic.SetupLogger()
	diagnostic.SetupRBAC()
	roleList := []string{role2.Administrator, "Manager"}
	for _, role := range roleList {
		t.Logf("Menu List(%v): %v", role, string(GetMenuListOfRole(role)))
	}
}

func TestGetAttributeListOfRole(t *testing.T) {
	diagnostic.SetupLogger()
	diagnostic.SetupRBAC()
	roleList := []string{role2.Administrator, "Manager"}
	for _, role := range roleList {
		t.Logf("Attribute List(%v): %v", role, string(GetAttributeListOfRole(role)))
	}
}

func TestHasRole(t *testing.T) {
	diagnostic.SetupLogger()
	diagnostic.SetupRBAC()

	role := "Manager"
	b := HasRole(role)
	if b {
		t.Logf("b(%v)", b)
	} else {
		t.Logf("b(%v)", b)
	}
}

func TestHasPermission(t *testing.T) {
	diagnostic.SetupLogger()
	diagnostic.SetupRBAC()

	role := "Manager"
	//role = role2.Administrator
	minor := 7

	b := HasPermission(role, minor)
	if b {
		t.Logf("b(%v)", b)
	} else {
		t.Logf("b(%v)", b)
	}
}

func TestHasMenu(t *testing.T) {
	diagnostic.SetupLogger()
	diagnostic.SetupRBAC()

	role := "Manager"
	menu := "Admission"
	item := "Track"

	b := HasMenu(role, menu, item)
	if b {
		t.Logf("b(%v)", b)
	} else {
		t.Logf("b(%v)", b)
	}
}

func TestHasAttribute(t *testing.T) {
	diagnostic.SetupLogger()
	diagnostic.SetupRBAC()

	//role := "Manager"
	role := role2.Administrator
	table := "user"
	//name := "country"
	name := "account"

	b := HasAttribute(role, table, name)
	if b {
		t.Logf("b(%v)", b)
	} else {
		t.Logf("b(%v)", b)
	}
}
