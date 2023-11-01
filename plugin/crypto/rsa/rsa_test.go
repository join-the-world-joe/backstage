package rsa

import (
	"backstage/common/crypto"
	"backstage/utils/convert"
	"testing"
)

func TestRSAEncrypt(t *testing.T) {
	plainText := []byte("Hello, I am Joe!")
	cto := NewCrypto(
		WithPublicKey([]byte(crypto.PublicKey)),
	)

	cipherText, err := cto.Encrypt(plainText)
	if err != nil {
		t.Error(err)
		return
	}

	//t.Log(string(cipherText))

	t.Log(convert.Bytes2StringArray(cipherText))
}

func TestRSADecrypt(t *testing.T) {
	cipherText := []byte{86, 32, 109, 247, 2, 237, 47, 242, 169, 143, 215, 205, 75, 89, 20, 125, 141, 108, 45, 53, 91, 63, 170, 27, 207, 65, 220, 93, 156, 120, 42, 221, 211, 217, 122, 121, 27, 95, 28, 181, 216, 98, 67, 52, 177, 248, 208, 98, 52, 45, 203, 179, 33, 224, 94, 194, 164, 153, 180, 103, 121, 207, 190, 0, 62, 133, 0, 210, 77, 249, 134, 24, 80, 213, 73, 116, 22, 69, 157, 255, 29, 255, 71, 239, 41, 150, 202, 226, 92, 251, 234, 58, 50, 184, 177, 181, 93, 243, 65, 214, 181, 50, 225, 198, 203, 185, 177, 102, 249, 99, 168, 179, 149, 110, 58, 96, 228, 217, 221, 177, 64, 46, 197, 203, 109, 94, 86, 87}
	cto := NewCrypto(
		WithPrivateKey([]byte(crypto.PrivateKey)),
	)

	plainText, err := cto.Decrypt(cipherText)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(string(plainText))
}
