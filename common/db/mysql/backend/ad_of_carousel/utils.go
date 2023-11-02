package ad_of_carousel

import (
	"backstage/global/mysql"
	"fmt"
)

func Insert(model *Model) (*Model, error) {
	temp, err := mysql.Insert(GetWhich(), GetDbName(), GetTableName(), model)
	if err != nil {
		return nil, err
	}
	model, ok := temp.(*Model)
	if !ok {
		return nil, fmt.Errorf("carousel_advertisement.insert failure, convert to Model fail")
	}
	if model.Id == 0 {
		return nil, fmt.Errorf("carousel_advertisement.insert failure, model.Id == 0")
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
