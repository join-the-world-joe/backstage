package ad_of_carousel

import "fmt"

func sqlMaxId() string {
	return fmt.Sprintf("SELECT MAX(id) FROM %s ", GetTableName())
}

func sqlSelectModelWithMaxId1() string {
	return fmt.Sprintf("SELECT id,image_path_list,created_at FROM %v WHERE id = (SELECT MAX(id) FROM %s) ", GetTableName(), GetTableName())
}

func sqlSelectModelWithMaxId() string {
	return fmt.Sprintf("SELECT id, advertisement_id_list, created_at, description FROM %v WHERE id = (SELECT MAX(id) FROM %v)", GetTableName(), GetTableName())
}

func sqlDeleteOutdatedRecords(version int64) string {
	return fmt.Sprintf("DELETE FROM %s WHERE version < %v", GetTableName(), version)
}
