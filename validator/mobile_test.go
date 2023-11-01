package validator

import "testing"

func TestIsMobileValid(t *testing.T) {
	countryCode := "86"
	phoneNumber := "184009226371"
	t.Log(IsMobileValid(countryCode, phoneNumber))
}
