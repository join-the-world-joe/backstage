package ad_of_bbq_products

import "fmt"

func sqlMaxId() string {
	return fmt.Sprintf("SELECT MAX(id) FROM %s ", GetTableName())
}

func sqlSelectModelWithMaxId() string {
	return fmt.Sprintf("SELECT id, ad_id_list, created_at,description FROM %v WHERE id = (SELECT MAX(id) FROM %v)", GetTableName(), GetTableName())
}
