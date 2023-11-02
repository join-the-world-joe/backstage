package rsa

import (
	"backstage/common/crypto"
	"testing"
)

func TestEncryptPassword(t *testing.T) {
	password := "hello, world"
	separator := 8
	cipherText, err := EncryptPassword(password, separator, crypto.PublicKey)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("cipherText: ", cipherText)

	pass, err := DecryptPassword(cipherText, separator, crypto.PrivateKey)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("pass: ", pass)
}

func TestVerifyKeyPair(t *testing.T) {
	privateKey, publicKey, err := GenKeyPair()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("public key: ")
	t.Log(publicKey)
	t.Log("private key: ")
	t.Log(privateKey)

	err = VerifyKeyPair(privateKey, publicKey)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("verify done!")
}
