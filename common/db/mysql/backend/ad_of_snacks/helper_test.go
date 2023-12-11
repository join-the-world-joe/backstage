package ad_of_snacks

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
	idList := []int{7, 8, 9}
	bytes, err := json.Marshal(&idList)
	if err != nil {
		t.Fatal(err)
	}
	diagnostic.SetupMySQL()
	temp, err := Insert(&Model{AdvertisementIdList: string(bytes)})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(temp)
}
