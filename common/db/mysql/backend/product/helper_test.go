package product

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

	diagnostic.SetupMySQL()

	_, err := InsertModel(&Model{
		Name:        "蒙牛酸酸乳",
		BuyingPrice: 100,
		Visible:     1,
		Status:      0, // 为0时, 采用数据库设定的默认值
		Vendor:      "汕头市蒙牛奶业有限公司",
		Contact:     "0756-88788371",
		Description: "正规渠道、国产大牌子、大人小孩都爱",
	})
	if err != nil {
		t.Fatal(err)
	}
	_, err = InsertModel(&Model{
		Name:        "恒大冰泉",
		BuyingPrice: 100,
		Visible:     1,
		Status:      0,
		Vendor:      "珠海市水之源有限公司",
		Contact:     "0878-88788371",
		Description: "二级代理、正归渠道、物流优势",
	})
	_, err = InsertModel(&Model{
		Name:        "橡胶袋子",
		BuyingPrice: 100,
		Visible:     1,
		Status:      0,
		Vendor:      "天津市万国集团有限公司",
		Contact:     "020-97718232",
		Description: "一流服务、二流产品、进货便宜",
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetModelListByIdList(t *testing.T) {
	diagnostic.SetupMySQL()
	idList := []int64{1, 2, 3}
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
	visible := 0
	err := UpdateVisibleById(id, visible)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUpdateVisibleByIdList(t *testing.T) {
	diagnostic.SetupMySQL()
	idList := []int64{1, 2, 3}
	visible := 0
	err := UpdateVisibleByIdList(idList, visible)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUpdateFieldListById(t *testing.T) {
	diagnostic.SetupMySQL()
	id := int64(1)
	name := "蒙牛酸酸乳new"
	buyingPrice := 80
	status := 0
	vendor := "汕头市蒙牛奶业有限公司"
	contact := "0829-9982912332"
	description := "正规渠道、国产大牌子、大人小孩都爱"
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
