package version_of_ad_of_carousel

import "fmt"

func sqlMaxId() string {
	return fmt.Sprintf("SELECT MAX(id) FROM %s ", GetTableName())
}
