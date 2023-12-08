package selling_point_of_advertisement

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

func TestInsertModel1(t *testing.T) {
	advertisementId := int64(1)
	sellingPoint := "优质奶源"
	diagnostic.SetupMySQL()
	temp, err := InsertModel(&Model{
		AdvertisementId: advertisementId,
		SellingPoint:    sellingPoint,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(temp)
}

func TestInsertModel(t *testing.T) {
	advertisementId := int64(1)
	sellingPoint := "含钙"
	diagnostic.SetupMySQL()
	temp, err := InsertModel(&Model{
		AdvertisementId: advertisementId,
		SellingPoint:    sellingPoint,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(temp)
}

func TestUpdateVisibleByIdList(t *testing.T) {
	idList := []int64{1, 2, 3}
	diagnostic.SetupMySQL()
	err := UpdateVisibleByIdList(idList, 1)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUpdateVisibleByAdvertisementIdAndSellingPoint(t *testing.T) {
	advertisementId := int64(1)
	sellingPoint := "解渴"
	visible := 0
	diagnostic.SetupMySQL()
	err := UpdateVisibleByAdvertisementIdAndSellingPoint(advertisementId, sellingPoint, visible)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetModelListByAdvertisementId(t *testing.T) {
	advertisementId := int64(1)
	diagnostic.SetupMySQL()
	ml, err := GetModelListByAdvertisementId(advertisementId)
	if err != nil {
		t.Fatal(err)
	}
	bytes, err := json.Marshal(ml)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))
}
