package user

import (
	"backstage/diagnostic"
	"backstage/global/mysql"
	"testing"
)

func TestAutoMigrate(t *testing.T) {
	diagnostic.SetupMySQL()
	err := mysql.AutoMigrate(GetWhich(), GetDbName(), GetTableName(), &Model{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestInsert(t *testing.T) {
	countryCode := "86"
	phoneNumber := "11111"
	diagnostic.SetupMySQL()
	temp, err := Insert(&Model{CountryCode: countryCode, PhoneNumber: phoneNumber})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(temp)
}
