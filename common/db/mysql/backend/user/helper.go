package user

func GetWhich() string {
	return "test"
}

func GetDbName() string {
	return "backend"
}

func GetTableName() string {
	return "user"
}

func GetTableNameList() []string {
	return []string{GetTableName()}
}
