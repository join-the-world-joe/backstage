package aes

import (
	"backstage/common/crypto"
	"backstage/utils/convert"
	"testing"
)

func TestAESEncrypt(t *testing.T) {
	plainText := `{"header":{"major":"4","minor":"2"},"body":{"code":-35}}`
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

	//t.Log("Raw: ", string(cipherText))
	t.Log(convert.Bytes2StringArray(cipherText))
}

func TestAESDecrypt(t *testing.T) {
	cto := NewCrypto(
		WithKey(crypto.Key),
		WithIV(crypto.IV),
		WithPadding(PKCS7),
	)

	cipherText := []byte{171, 205, 28, 130, 28, 35, 176, 121, 112, 169, 215, 212, 45, 230, 171, 164, 9, 47, 2, 106, 134, 145, 242, 21, 32, 80, 172, 58, 208, 191, 25, 4, 178, 125, 198, 196, 25, 60, 136, 246, 196, 195, 237, 137, 25, 216, 18, 90, 212, 50, 212, 113, 120, 158, 255, 23, 180, 136, 116, 1, 130, 172, 251, 100}

	plainText, err := cto.Decrypt(cipherText)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(string(plainText))
}
