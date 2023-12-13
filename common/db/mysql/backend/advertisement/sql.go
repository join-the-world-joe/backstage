package advertisement

import (
	"backstage/utils/strings"
	"fmt"
)

func sqlSelectIdListInTable() string {
	return fmt.Sprintf("SELECT id FROM %v WHERE visible = 1", GetTableName())
}

func sqlSelectIdListByIDList(idList []int64) string {
	return fmt.Sprintf("SELECT id FROM %v WHERE visible = 1 AND id in %s", GetTableName(), strings.WithinParenthesesInt64(idList))
}
