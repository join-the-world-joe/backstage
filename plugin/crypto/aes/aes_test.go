package aes

import (
	"backstage/common/crypto"
	"backstage/utils/convert"
	"testing"
)

func TestAESEncrypt(t *testing.T) {
	plainText := "123456"
	cto := NewCrypto(
		WithKey(crypto.Key),
		WithIV(crypto.IV),
		WithPadding(PKCS7),
	)

	cipherText, err := cto.Encrypt([]byte(plainText))
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(convert.Bytes2StringArray(cipherText))
}

func TestAESDecrypt(t *testing.T) {
	cto := NewCrypto(
		WithKey(crypto.Key),
		WithIV(crypto.IV),
		WithPadding(PKCS7),
	)

	cipherText := []byte{115, 95, 42, 5, 150, 253, 116, 174, 140, 237, 86, 198, 251, 217, 193, 24}

	plainText, err := cto.Decrypt(cipherText)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(string(plainText))
}
