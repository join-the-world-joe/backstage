package ad_of_snacks

import "fmt"

func sqlMaxId() string {
	return fmt.Sprintf("SELECT MAX(id) FROM %s ", GetTableName())
}

func sqlSelectModelWithMaxId() string {
	return fmt.Sprintf("SELECT id, advertisement_id_list, created_at,description FROM %v WHERE id = (SELECT MAX(id) FROM %v)", GetTableName(), GetTableName())
}
