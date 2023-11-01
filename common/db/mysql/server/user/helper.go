package user

func GetWhich() string {
	return "test"
}

func GetDbName() string {
	return "server"
}

func GetTableName() string {
	return "user"
}

func GetTableNameList() []string {
	return []string{GetTableName()}
}
