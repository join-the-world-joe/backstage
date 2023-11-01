package rsa

import (
	"fmt"
	"go-micro-framework/plugin/crypto/rsa"
	"strconv"
	"strings"
	"time"
)

func EncryptPassword(password string, separator int, publicKey string) ([]byte, error) {
	_rsa := rsa.NewCrypto(
		rsa.WithPrivateKey([]byte("")),
		rsa.WithPublicKey([]byte(publicKey)),
	)
	plainText := fmt.Sprintf("%v%v%v", password, fmt.Sprintf("%c", separator), time.Now().Unix())
	return _rsa.Encrypt([]byte(plainText))
}

func DecryptPassword(cipherText []byte, separator int, privateKey string) (string, error) {
	_rsa := rsa.NewCrypto(
		rsa.WithPrivateKey([]byte(privateKey)),
		rsa.WithPublicKey([]byte("")),
	)
	plainText, err := _rsa.Decrypt(cipherText)
	if err != nil {
		return "", err
	}
	list := strings.Split(string(plainText), fmt.Sprintf("%c", separator))
	if len(list) < 2 {
		return "", fmt.Errorf("len(list) < 2")
	}

	password := list[0]
	expire, err := strconv.Atoi(list[1])
	if err != nil {
		return "", err
	}

	// 验证密码是否过期
	if time.Now().Unix() > int64(expire)+10 {
		return "", fmt.Errorf("expired")
	}
	return password, nil
}
