package product

func GetWhich() string {
	return "test"
}

func GetDbName() string {
	return "backend"
}

func GetTableName() string {
	return "product"
}

func GetTableNameList() []string {
	return []string{GetTableName()}
}
