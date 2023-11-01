package template

import "strconv"

func GetWhich(id int) string {
	return "Backend" + strconv.Itoa((id%Mod)+1)
}

func GetDBName() string {
	return "test"
}

func GetTableName() string {
	return "template"
}
