package ad_of_carousel

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
		return nil, fmt.Errorf("carousel_advertisement.insert failure, convert to Model fail")
	}
	if model.Id == 0 {
		return nil, fmt.Errorf("carousel_advertisement.insert failure, model.Id == 0")
	}
	return model, nil
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

func RemoveOutdatedRecordsOfAdCarousel(version int64) error {
	return mysql.Delete(GetWhich(), GetDbName(), sqlDeleteOutdatedRecords(version))
}
