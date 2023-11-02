package version_of_ad_of_bbq_products

import (
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

func TestInsertModel(t *testing.T) {
	diagnostic.SetupMySQL()
	temp, err := InsertModel(&Model{})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(temp)
}

func TestGetMaxId(t *testing.T) {
	diagnostic.SetupMySQL()
	if id, err := GetMaxId(); err == nil {
		t.Log("id: ", id)
	}
}
