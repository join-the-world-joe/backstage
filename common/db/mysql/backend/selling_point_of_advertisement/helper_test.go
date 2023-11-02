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

func TestInsertModel(t *testing.T) {
	productId := int64(3)
	sellingPoint := "解渴3"
	diagnostic.SetupMySQL()
	temp, err := InsertModel(&Model{
		ProductId:    productId,
		SellingPoint: sellingPoint,
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

func TestUpdateVisibleByProductIdAndSellingPoint(t *testing.T) {
	productId := int64(1)
	sellingPoint := "解渴"
	visible := 0
	diagnostic.SetupMySQL()
	err := UpdateVisibleByProductIdAndSellingPoint(productId, sellingPoint, visible)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetModelListByProductId(t *testing.T) {
	productId := int64(1)
	diagnostic.SetupMySQL()
	ml, err := GetModelListByProductId(productId)
	if err != nil {
		t.Fatal(err)
	}
	bytes, err := json.Marshal(ml)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))
}
