package user

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
		return nil, fmt.Errorf("user.insert failure, convert to Model fail")
	}
	if model.Id == 0 {
		return nil, fmt.Errorf("user.insert failure, model.Id == 0")
	}
	return model, nil
}

func GetModelById(id int64) (*Model, error) {
	db, err := mysql.GetDB(GetWhich(), GetDbName())
	if err != nil {
		return nil, err
	}
	model := &Model{}
	err = db.Table(GetTableName()).Where("id = ?", id).First(model).Error
	if err != nil {
		return nil, err
	}

	return model, nil
}

func GetModelByMobile(countryCode, phoneNumber string) (*Model, error) {
	db, err := mysql.GetDB(GetWhich(), GetDbName())
	if err != nil {
		return nil, err
	}
	model := &Model{}
	err = db.Table(GetTableName()).Where("country_code = ? AND phone_number = ?", countryCode, phoneNumber).First(model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}
