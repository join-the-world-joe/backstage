package validator

import (
	"github.com/nyaruka/phonenumbers"
)

func IsMobileValid(countryCode, phoneNumber string) bool {
	mobile := "+" + countryCode + phoneNumber
	temp, err := phonenumbers.Parse(mobile, "")
	if err != nil {
		return false
	}
	return phonenumbers.IsValidNumber(temp)
}
