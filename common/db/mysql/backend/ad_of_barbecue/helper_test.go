package ad_of_barbecue

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
	idList := []int{1, 2, 3}
	bytes, err := json.Marshal(&idList)
	if err != nil {
		t.Fatal(err)
	}
	version := int64(1)
	diagnostic.SetupMySQL()
	temp, err := InsertModel(
		&Model{
			Version:             version,
			AdvertisementIdList: string(bytes),
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(temp)
}
