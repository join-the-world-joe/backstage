package user_role

import (
	"backstage/global/mysql"
	"fmt"
	"golang.org/x/exp/slices"
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

func GetModelByUserId(userId int64) (*Model, error) {
	db, err := mysql.GetDB(GetWhich(), GetDbName())
	if err != nil {
		return nil, err
	}
	model := &Model{}
	err = db.Table(GetTableName()).Where("user_id = ?", userId).First(model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func GetRoleListByUserId(userId int64) []string {
	var role string
	rows, err := mysql.Query(GetWhich(), GetDbName(), sqlQueryRoleListByUserId(userId))
	if err != nil {
		return nil
	}
	roleList := []string{}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&role)
		roleList = append(roleList, role)
	}
	return roleList
}

func GetUserIdListByRole(role string) []int64 {
	var userId int64
	rows, err := mysql.Query(GetWhich(), GetDbName(), sqlQueryUserIdListByRole(role))
	if err != nil {
		return nil
	}
	userList := []int64{}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&userId)
		userList = append(userList, userId)
	}
	return userList
}

func GetUserIdListByRoleList(roleList []string) []int64 {
	userIdList := []int64{}
	for _, roleName := range roleList {
		uidList := GetUserIdListByRole(roleName)
		for _, uid := range uidList {
			if !slices.Contains(userIdList, uid) {
				userIdList = append(userIdList, uid)
			}
		}
	}
	return userIdList
}

func UpdateVisibleByUserId(id int64, visible int) error {
	db, err := mysql.GetDB(GetWhich(), GetDbName())
	if err != nil {
		return err
	}
	return db.Table(GetTableName()).Where("user_id = ?", id).Updates(map[string]interface{}{"visible": visible}).Error
}

func UpdateVisibleByRole(role string, visible int) error {
	db, err := mysql.GetDB(GetWhich(), GetDbName())
	if err != nil {
		return err
	}

	return db.Table(GetTableName()).Where("role = ?", role).Updates(map[string]interface{}{"visible": visible}).Error
}

func GetIdByRoleAndVisible(role string, visible int) int64 {
	var id int64
	rows, err := mysql.Query(GetWhich(), GetDbName(), sqlQueryIdByRole(role, visible))
	if err != nil {
		return 0
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&id)
	}
	return id
}

func GetUserIdList() []int64 {
	var userId int64
	rows, err := mysql.Query(GetWhich(), GetDbName(), sqlQueryUserIdList())
	if err != nil {
		return nil
	}
	userList := []int64{}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&userId)
		userList = append(userList, userId)
	}
	return userList
}
