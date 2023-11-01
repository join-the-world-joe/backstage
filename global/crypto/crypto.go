package crypto

import (
	"backstage/abstract/crypto"
	crypto2 "backstage/common/crypto"
	"backstage/plugin/crypto/aes"
	"backstage/plugin/crypto/rsa"
)

var _rsa crypto.Crypto
var _aes crypto.Crypto

func init() {
	_rsa = rsa.NewCrypto(
		rsa.WithPrivateKey([]byte(crypto2.PrivateKey)),
		rsa.WithPublicKey([]byte(crypto2.PublicKey)),
	)
	_aes = aes.NewCrypto(
		aes.WithKey(crypto2.Key),
		aes.WithIV(crypto2.IV),
		aes.WithPadding(aes.PKCS7),
	)
}

func RSAEncrypt(plainText []byte) ([]byte, error) {
	return _rsa.Encrypt(plainText)
}

func RSADecrypt(cipherText []byte) ([]byte, error) {
	return _rsa.Decrypt(cipherText)
}

func AESEncrypt(plainText []byte) ([]byte, error) {
	return _aes.Encrypt(plainText)
}

func AESDecrypt(cipherText []byte) ([]byte, error) {
	return _aes.Decrypt(cipherText)
}
