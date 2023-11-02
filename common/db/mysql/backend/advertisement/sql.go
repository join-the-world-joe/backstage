package advertisement

import "fmt"

func sqlSelectIdListInTable() string {
	return fmt.Sprintf("SELECT id FROM %v WHERE visible = 1", GetTableName())
}
