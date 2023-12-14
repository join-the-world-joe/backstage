package advertisement

import (
	"backstage/utils/strings"
	"fmt"
)

func sqlSelectIdListInTable() string {
	return fmt.Sprintf("SELECT id FROM %v WHERE visible = 1", GetTableName())
}

func sqlSelectIdListByIDListWithStatus(idList []int64) string {
	return fmt.Sprintf("SELECT id FROM %v WHERE visible = 1 AND id in %s AND status = 1", GetTableName(), strings.WithinParenthesesInt64(idList))
}

func sqlSelectIdListByIDListWithoutStatus(idList []int64) string {
	return fmt.Sprintf("SELECT id FROM %v WHERE visible = 1 AND id in %s", GetTableName(), strings.WithinParenthesesInt64(idList))
}
