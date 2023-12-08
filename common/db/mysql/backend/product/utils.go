package product

import (
	"backstage/global/mysql"
	"backstage/utils/strings"
	"fmt"
)

func InsertModel(model *Model) (*Model, error) {
	temp, err := mysql.Insert(GetWhich(), GetDbName(), GetTableName(), model)
	if err != nil {
		return nil, err
	}
	model, ok := temp.(*Model)
	if !ok {
		return nil, fmt.Errorf("product.insert failure, convert to Model fail")
	}
	if model.Id == 0 {
		return nil, fmt.Errorf("product.insert failure, model.Id == 0")
	}
	return model, nil
}

func GetIdListInTable() ([]int64, error) {
	idList := []int64{}
	var id int64
	rows, err := mysql.Query(GetWhich(), GetDbName(), sqlSelectIdListInTable())
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&id)
		idList = append(idList, id)
	}
	return idList, nil
}

func GetModelListByIdList(idList []int64) ([]*Model, error) {
	modelList := []*Model{}
	db, err := mysql.GetDB(GetWhich(), GetDbName())
	if err != nil {
		return nil, err
	}
	err = db.Table(GetTableName()).Where(fmt.Sprintf("id in %s AND visible = 1", strings.WithinParenthesesInt64(idList))).Find(&modelList).Error
	if err != nil {
		return nil, err
	}
	return modelList, nil
}

func GetModelById(id int64) (*Model, error) {
	db, err := mysql.GetDB(GetWhich(), GetDbName())
	if err != nil {
		return nil, err
	}
	model := &Model{}
	err = db.Table(GetTableName()).Where("id = ? AND visible = 1", id).First(model).Error
	if err != nil {
		return nil, err
	}

	return model, nil
}

func GetModelListByName(name string) ([]*Model, error) {
	modelList := []*Model{}
	db, err := mysql.GetDB(GetWhich(), GetDbName())
	if err != nil {
		return nil, err
	}
	err = db.Table(GetTableName()).Where("name = ? AND visible = 1", name).Find(&modelList).Error
	if err != nil {
		return nil, err
	}
	return modelList, nil
}

func UpdateVisibleById(id int64, visible int) error {
	db, err := mysql.GetDB(GetWhich(), GetDbName())
	if err != nil {
		return err
	}
	return db.Table(GetTableName()).Where("id = ?", id).Updates(map[string]interface{}{"visible": visible}).Error
}

func UpdateVisibleByIdList(idList []int64, visible int) error {
	db, err := mysql.GetDB(GetWhich(), GetDbName())
	if err != nil {
		return err
	}
	return db.Table(GetTableName()).Where(fmt.Sprintf("id in %s", strings.WithinParenthesesInt64(idList))).Updates(map[string]interface{}{"visible": visible}).Error
}

func UpdateFieldListById(id int64, fieldList map[string]interface{}) error {
	db, err := mysql.GetDB(GetWhich(), GetDbName())
	if err != nil {
		return err
	}
	return db.Table(GetTableName()).Where("id = ?", id).Updates(fieldList).Error
}
