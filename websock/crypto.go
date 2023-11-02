package main

import (
	"backstage/abstract/crypto"
	crypto2 "backstage/common/crypto"
	"backstage/plugin/crypto/aes"
)

var _crypto crypto.Crypto

func init() {
	_crypto = aes.NewCrypto(
		aes.WithKey(crypto2.Key),
		aes.WithIV(crypto2.IV),
		aes.WithPadding(aes.PKCS7),
	)
}

func encrypt(plainText []byte) ([]byte, error) {
	return _crypto.Encrypt(plainText)
}

func decrypt(cipherText []byte) ([]byte, error) {
	return _crypto.Decrypt(cipherText)
}
