package user

import (
	"backstage/utils/strings"
	"fmt"
)

func sqlDeleteAnyById(id int64) string {
	return fmt.Sprintf("DELETE FROM %s WHERE id = %v", GetTableName(), id)
}

func sqlQueryUserIdListNotInUserIdList(userIdList []int64) string {
	return fmt.Sprintf("SELECT id FROM %v WHERE id NOT IN %v AND visible = 1", GetTableName(), strings.WithinParenthesesInt64(userIdList))
}

func sqlQueryUserNameById(id int64) string {
	return fmt.Sprintf("SELECT name FROM %v WHERE visible = 1 AND id = %v", GetTableName(), id)
}
