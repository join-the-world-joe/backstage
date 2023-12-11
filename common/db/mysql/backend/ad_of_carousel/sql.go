package ad_of_carousel

import "fmt"

func sqlDeleteOutdatedRecords(version int64) string {
	return fmt.Sprintf("DELETE FROM %s WHERE version < %v", GetTableName(), version)
}

func sqlSelectModelWithVersion(version int64) string {
	return fmt.Sprintf("SELECT * FROM %s WHERE version = %v", GetTableName(), version)
}
