package token

import (
	"backstage/diagnostic"
	"github.com/google/uuid"
	"testing"
)

func TestCreate(t *testing.T) {
	countryCode := "86"
	phoneNumber := "110"
	userId := "3"
	token := uuid.New().String()
	diagnostic.SetupCache()
	diagnostic.SetupLogger()
	err := Create(countryCode, phoneNumber, userId, token)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGet(t *testing.T) {
	countryCode := "86"
	phoneNumber := "113355"
	token := "da62c3a3-f2a2-4b72-8592-7309c105aead"
	diagnostic.SetupCache()
	diagnostic.SetupLogger()
	temp, err := Get(countryCode, phoneNumber, token)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("temp:", temp)
}
