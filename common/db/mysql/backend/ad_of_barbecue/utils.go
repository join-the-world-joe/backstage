package ad_of_barbecue

import (
	"backstage/global/mysql"
	"fmt"
)

func InsertModel(model *Model) (*Model, error) {
	temp, err := mysql.Insert(GetWhich(), GetDbName(), GetTableName(), model)
	if err != nil {
		return nil, err
	}
	model, ok := temp.(*Model)
	if !ok {
		return nil, fmt.Errorf("ad_of_bbq_products.insert failure, convert to Model fail")
	}
	if model.Id == 0 {
		return nil, fmt.Errorf("ad_of_bbq_products.insert failure, model.Id == 0")
	}
	return model, nil
}

func GetMaxId() (int64, error) {
	var id int64
	rows, err := mysql.Query(GetWhich(), GetDbName(), sqlMaxId())
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&id)
	}
	return id, nil
}

func GetLatestVersionModel() (*Model, error) {
	m := &Model{}
	db, err := mysql.GetDB(GetWhich(), GetDbName())
	if err != nil {
		return nil, err
	}
	err = db.Raw(sqlSelectModelWithMaxId()).Scan(&m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func GetModelByVersion(version int64) (*Model, error) {
	m := &Model{}
	db, err := mysql.GetDB(GetWhich(), GetDbName())
	if err != nil {
		return nil, err
	}
	err = db.Raw(sqlSelectModelWithVersion(version)).Scan(&m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}
