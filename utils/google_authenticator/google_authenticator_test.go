package google_authentictor

import "testing"

func TestCreateCode(t *testing.T) {
	secret := CreateGoogleSecret("666666")
	t.Log(secret)
}

func TestGetGoogleCode(t *testing.T) {
	secret := "CBFHB5CB7GB766CA"
	code, remain := GetGoogleCode(secret)
	t.Log("Code: ", code)
	t.Log("Remain: ", remain)
}
