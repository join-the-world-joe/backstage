package selling_point_of_advertisement

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
		return nil, fmt.Errorf("selling_point_of_advertisement.InsertModel failure, convert to Model fail")
	}
	if model.Id == 0 {
		return nil, fmt.Errorf("selling_point_of_advertisement.InsertModel failure, model.Id == 0")
	}
	return model, nil
}

func GetModelListByAdvertisementId(advertisementId int64) ([]*Model, error) {
	modelList := []*Model{}
	db, err := mysql.GetDB(GetWhich(), GetDbName())
	if err != nil {
		return nil, err
	}
	err = db.Table(GetTableName()).Where("advertisement_id = ? AND visible = 1", advertisementId).Find(&modelList).Error
	if err != nil {
		return nil, err
	}
	return modelList, nil
}

func UpdateVisibleByIdList(idList []int64, visible int) error {
	db, err := mysql.GetDB(GetWhich(), GetDbName())
	if err != nil {
		return err
	}
	return db.Table(GetTableName()).Where(fmt.Sprintf("id in %s AND visible = 1", strings.WithinParenthesesInt64(idList))).Updates(map[string]interface{}{"visible": visible}).Error
}

func UpdateVisibleByAdvertisementIdAndSellingPoint(advertisementId int64, sellingPoint string, visible int) error {
	db, err := mysql.GetDB(GetWhich(), GetDbName())
	if err != nil {
		return err
	}
	return db.Table(GetTableName()).Where(fmt.Sprintf("advertisement_id = %v AND selling_point = '%v' AND visible = 1", advertisementId, sellingPoint)).Updates(map[string]interface{}{"visible": visible}).Error
}
