package template

import "strconv"

func GetWhich() string {
	return "HongKong"
}

func GetDbName() string {
	return "test"
}

func GetTableName(id int) string {
	return "template" + strconv.Itoa((id%Mod)+1)
}

func GetTableNameList() []string {
	nameList := make([]string, 0)
	for i := 1; i <= Mod; i++ {
		nameList = append(nameList, GetTableName(i))
	}
	return nameList
}
