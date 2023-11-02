package user_role

import "fmt"

func sqlQueryRoleListByUserId(userId int64) string {
	return fmt.Sprintf("SELECT role FROM %v WHERE user_id = %v AND visible = 1", GetTableName(), userId)
}

func sqlQueryUserIdListByRole(role string) string {
	return fmt.Sprintf("SELECT user_id FROM %v WHERE role = '%v' AND visible = 1", GetTableName(), role)
}

func sqlQueryUserIdList() string {
	return fmt.Sprintf("SELECT user_id FROM %v WHERE visible = 1", GetTableName())
}

func sqlQueryIdByRole(role string, visible int) string {
	return fmt.Sprintf("SELECT id FROM %v WHERE role = '%v' AND visible = %v", GetTableName(), role, visible)
}
