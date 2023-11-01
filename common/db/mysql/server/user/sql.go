package user

import "fmt"

func sql1(countryCode, phoneNumber string) string {
	return fmt.Sprintf("SELECT * FROM %v WHERE country_code = %v AND phone_number = %v", GetTableName(), countryCode, phoneNumber)
}
