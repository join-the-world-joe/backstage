package advertisement

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
	diagnostic.SetupMySQL()

	temp, err := InsertModel(&Model{
		Name:          "春季牛奶",
		Title:         "健康的乳",
		PlaceOFOrigin: "内蒙古",
		Image:         "牛奶外观图片网址",
		SellingPrice:  200,
		Visible:       1,
		Stock:         5,
		ProductId:     1,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(temp)
}

func TestInsertModel2(t *testing.T) {
	diagnostic.SetupMySQL()

	temp, err := InsertModel(&Model{
		Name:          "夏季矿物水广告",
		Title:         "健康的水",
		PlaceOFOrigin: "四川",
		Image:         "水瓶外观图片网址",
		SellingPrice:  100,
		Visible:       1,
		Stock:         5,
		ProductId:     1,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(temp)
}

func TestInsertModel3(t *testing.T) {
	diagnostic.SetupMySQL()

	temp, err := InsertModel(&Model{
		Name:          "太阳灯",
		Title:         "自然光、无紫外线、无辐射",
		PlaceOFOrigin: "吸杂",
		Image:         "水瓶外观图片网址2",
		SellingPrice:  300,
		Visible:       1,
		Stock:         5,
		ProductId:     1,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(temp)
}

func TestGetIdListInTable(t *testing.T) {
	diagnostic.SetupMySQL()
	idList, err := GetIdListInTable()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("id list: ", idList)
}

func TestGetModelListByIdList(t *testing.T) {
	idList := []int64{1}
	diagnostic.SetupMySQL()
	ml, err := GetModelListByIdList(idList)
	if err != nil {
		t.Fatal(err)
	}
	bytes, err := json.Marshal(ml)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("model list: ", string(bytes))
}

func TestUpdateFieldListById(t *testing.T) {
	id := int64(1)
	name := "夏季矿物水广告new"
	title := "健康的水 new"
	placeOfOrigin := "四川 new"
	url := "水瓶外观图片网址 new"
	sellingPrice := 800
	status := 0
	stock := 100
	productId := int64(1)
	description := "关联商品名称 new"
	filedList := map[string]interface{}{
		"name":            name,
		"title":           title,
		"place_of_origin": placeOfOrigin,
		"url":             url,
		"selling_price":   sellingPrice,
		"status":          status,
		"stock":           stock,
		"product_id":      productId,
		"description":     description,
	}
	diagnostic.SetupMySQL()
	err := UpdateFieldListById(id, filedList)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUpdateVisibleByIdList(t *testing.T) {
	idList := []int64{1, 2}
	visible := 0
	diagnostic.SetupMySQL()
	err := UpdateVisibleByIdList(idList, visible)
	if err != nil {
		t.Fatal(err)
	}
}
