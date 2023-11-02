package user_role

import (
	role2 "backstage/common/macro/role"
	"backstage/diagnostic"
	"backstage/global/mysql"
	"testing"
)

func TestAutoMigrate(t *testing.T) {
	diagnostic.SetupMySQL()
	err := mysql.AutoMigrate(GetWhich(), GetDbName(), GetTableName(), &Model{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestDropTable(t *testing.T) {
	diagnostic.SetupMySQL()
	err := mysql.DropTable(GetWhich(), GetDbName(), GetTableName())
	if err != nil {
		t.Fatal(err)
	}
}

func TestInsertModel(t *testing.T) {
	userId := int64(1)
	role := role2.Administrator
	diagnostic.SetupMySQL()
	temp, err := InsertModel(&Model{UserId: userId, Role: role})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(temp)
}

func TestGetModelByUserId(t *testing.T) {
	userId := int64(5)
	diagnostic.SetupMySQL()
	m, err := GetModelByUserId(userId)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("m: ", m)
}

func TestGetRoleListByUserId(t *testing.T) {
	userId := int64(5)
	diagnostic.SetupMySQL()
	t.Log("Role List: ", GetRoleListByUserId(userId))
}

func TestGetUserIdListByRole(t *testing.T) {
	role := "Manager"
	diagnostic.SetupMySQL()
	t.Log("User Id List: ", GetUserIdListByRole(role))
}

func TestGetUserIdList(t *testing.T) {
	diagnostic.SetupMySQL()
	t.Log(GetUserIdList())
}
