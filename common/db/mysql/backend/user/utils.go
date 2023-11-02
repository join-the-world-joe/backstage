package user

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
		return nil, fmt.Errorf("user.InsertModel failure, convert to Model fail")
	}
	if model.Id == 0 {
		return nil, fmt.Errorf("user.InsertModel failure, model.Id == 0")
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

func GetModelByEmail(email string) (*Model, error) {
	db, err := mysql.GetDB(GetWhich(), GetDbName())
	if err != nil {
		return nil, err
	}
	model := &Model{}
	err = db.Table(GetTableName()).Where("email = ?", email).First(model).Error
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

func GetModelByAccount(account string) (*Model, error) {
	db, err := mysql.GetDB(GetWhich(), GetDbName())
	if err != nil {
		return nil, err
	}
	model := &Model{}
	err = db.Table(GetTableName()).Where("account = ?", account).First(model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func GetModelByPhoneNumber(phoneNumber string) (*Model, error) {
	db, err := mysql.GetDB(GetWhich(), GetDbName())
	if err != nil {
		return nil, err
	}
	model := &Model{}
	err = db.Table(GetTableName()).Where("phone_number = ?", phoneNumber).First(model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func GetModelByName(name string) (*Model, error) {
	db, err := mysql.GetDB(GetWhich(), GetDbName())
	if err != nil {
		return nil, err
	}
	model := &Model{}
	err = db.Table(GetTableName()).Where("name = ?", name).First(model).Error
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

func GetModelListByPhoneNumber(phoneNumber string) ([]*Model, error) {
	modelList := []*Model{}
	db, err := mysql.GetDB(GetWhich(), GetDbName())
	if err != nil {
		return nil, err
	}
	err = db.Table(GetTableName()).Where("phone_number = ?", phoneNumber).Find(&modelList).Error
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

func UpdateFieldListById(id int64, fieldList map[string]interface{}) error {
	db, err := mysql.GetDB(GetWhich(), GetDbName())
	if err != nil {
		return err
	}
	return db.Table(GetTableName()).Where("id = ?", id).Updates(fieldList).Error
}

func GetUserIdListNotInUserIdList(userIdList []int64) []int64 {
	uidList := []int64{}
	var userId int64
	rows, err := mysql.Query(GetWhich(), GetDbName(), sqlQueryUserIdListNotInUserIdList(userIdList))
	if err != nil {
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&userId)
		uidList = append(uidList, userId)
	}
	return uidList
}

func GetNameById(id int64) string {
	var name string
	rows, err := mysql.Query(GetWhich(), GetDbName(), sqlQueryUserNameById(id))
	if err != nil {
		return ""
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&name)
	}
	return name
}
