package product

import (
	"backstage/diagnostic"
	"backstage/global/mysql"
	"encoding/json"
	"fmt"
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
	n := 20
	product := "product"
	vendor := "vendor"
	contact := "contact"
	desc := "desc"

	diagnostic.SetupMySQL()

	for i := 0; i < n; i++ {

		temp, err := Insert(&Model{
			Name:        fmt.Sprintf("%s%d", product, i+1),
			BuyingPrice: 100,
			Visible:     1,
			Status:      0,
			Vendor:      fmt.Sprintf("%s%d", vendor, i+1),
			Contact:     fmt.Sprintf("%s%d", contact, i+1),
			Description: fmt.Sprintf("%s%d", desc, i+1),
		})
		if err != nil {
			t.Fatal(err)
		}

		t.Log(temp)
	}
}

func TestGetModelListByIdList(t *testing.T) {
	diagnostic.SetupMySQL()
	idList := []int64{3, 4, 5}
	ml, err := GetModelListByIdList(idList)
	if err != nil {
		t.Fatal(err)
	}
	bytes, err := json.Marshal(ml)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("bytes: ", string(bytes))
}

func TestGetIdListInTable(t *testing.T) {
	diagnostic.SetupMySQL()
	idList, err := GetIdListInTable()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("id list: ", idList)
}

func TestUpdateVisibleById(t *testing.T) {
	diagnostic.SetupMySQL()
	id := int64(1)
	visible := 1
	err := UpdateVisibleById(id, visible)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUpdateVisibleByIdList(t *testing.T) {
	diagnostic.SetupMySQL()
	idList := []int64{2, 3, 4}
	visible := 1
	err := UpdateVisibleByIdList(idList, visible)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUpdateFieldListById(t *testing.T) {
	diagnostic.SetupMySQL()
	id := int64(1)
	name := "xxx"
	buyingPrice := 1
	status := 1
	vendor := "vendor"
	contact := "contact"
	description := "description"
	fieldList := map[string]interface{}{}

	if len(name) > 0 {
		fieldList["name"] = name
	}

	if buyingPrice > 0 {
		fieldList["buying_price"] = buyingPrice
	}

	fieldList["status"] = status

	if len(vendor) > 0 {
		fieldList["vendor"] = vendor
	}

	if len(contact) > 0 {
		fieldList["contact"] = contact
	}

	if len(description) > 0 {
		fieldList["description"] = description
	}
	err := UpdateFieldListById(id, fieldList)
	if err != nil {
		t.Fatal(err)
	}
}
