package aes

import (
	"crypto/rand"
	"fmt"
	"go-micro-framework/plugin/crypto/aes"
	"strings"
)

func GenKeyIVPair() ([]byte, []byte, error) { // key, iv, error
	key := make([]byte, 32)
	iv := make([]byte, 16)
	_, err := rand.Read(key)
	if err != nil {
		return nil, nil, err
	}
	_, err = rand.Read(iv)
	if err != nil {
		return nil, nil, err
	}
	return key, iv, nil
}

func VerifyKeyPair(key, iv []byte) error {
	plainText := []byte("Hello, world!")
	cto := aes.NewCrypto(
		aes.WithKey(string(key)),
		aes.WithIV(string(iv)),
		aes.WithPadding(aes.PKCS7),
	)

	cipherText, err := cto.Encrypt(plainText)
	if err != nil {
		return err
	}

	genPlainText, err := cto.Decrypt(cipherText)
	if err != nil {
		return err
	}

	if strings.Compare(string(plainText), string(genPlainText)) != 0 {
		return fmt.Errorf("plainText not equal to genPlainText")
	}

	return nil
}
