package version_of_ad_of_hot_deals

import "fmt"

func sqlMaxId() string {
	return fmt.Sprintf("SELECT MAX(id) FROM %s ", GetTableName())
}
