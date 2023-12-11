package ad_of_carousel

import (
	"backstage/diagnostic"
	"backstage/global/mysql"
	"encoding/json"
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
	idList := []int64{1, 2, 3}
	bytes, err := json.Marshal(&idList)
	if err != nil {
		t.Fatal(err)
	}
	diagnostic.SetupMySQL()
	temp, err := InsertModel(&Model{AdvertisementIdList: string(bytes)})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(temp)
}

func TestGetModelByVersion(t *testing.T) {
	version := int64(3)
	diagnostic.SetupMySQL()
	model, err := GetModelByVersion(version)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("model: ", *model)
}
