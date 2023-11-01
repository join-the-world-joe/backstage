package template

import "strconv"

func GetWhich(id int) string {
	return "HongKong" + strconv.Itoa((id%Mod)+1)
}

func GetDbName() string {
	return "test"
}

func GetTableName() string {
	return "template"
}

func GetTableNameList() []string {
	return []string{GetTableName()}
}
