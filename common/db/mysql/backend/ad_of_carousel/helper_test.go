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

func TestInsert(t *testing.T) {
	imagePathList := []string{"path7", "path8", "path9"}
	bytes, err := json.Marshal(&imagePathList)
	if err != nil {
		t.Fatal(err)
	}
	diagnostic.SetupMySQL()
	temp, err := Insert(&Model{ImagePathList: string(bytes)})
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

func TestGetLatestVersionModel(t *testing.T) {
	diagnostic.SetupMySQL()
	if model, err := GetLatestVersionModel(); err == nil {
		t.Log("model: ", model)
	} else {
		t.Fatal(err)
	}
}
