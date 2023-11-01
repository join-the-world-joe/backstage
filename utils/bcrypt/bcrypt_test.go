package bcrypt

import (
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {
	password := "123456"
	hash, err := PasswordHash(password)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Hash: ", hash) // 固化至数据
}

func TestVerity(t *testing.T) {
	password := "hello"                                                    // from the output of other Encrypt Algorithm
	hash := "$2a$10$OtEBwB37ZXCgWgs5lU22vONQOjw4Wqq2e81IhTPwUhIHqbZwICD4K" // from database
	fmt.Println(PasswordVerify(password, hash))
}
