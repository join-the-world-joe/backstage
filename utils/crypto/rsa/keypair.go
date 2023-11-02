package rsa

import (
	rsa2 "backstage/plugin/crypto/rsa"
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"strings"
)

func GenKeyPair() (string, string, error) { // privateKey, publicKey, error
	privatekey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return "", "", err
	}
	publickey := &privatekey.PublicKey

	privateKeyBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privatekey),
	}

	privatePem := new(bytes.Buffer)
	err = pem.Encode(privatePem, privateKeyBlock)
	if err != nil {
		return "", "", err
	}

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publickey)
	if err != nil {
		return "", "", err
	}
	publicKeyBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}
	publicPem := new(bytes.Buffer)
	err = pem.Encode(publicPem, publicKeyBlock)
	if err != nil {
		return "", "", err
	}

	return privatePem.String(), publicPem.String(), nil
}

func VerifyKeyPair(privatePem, publicPem string) error {
	plainText := []byte("Hello, world!")
	cto := rsa2.NewCrypto(
		rsa2.WithPublicKey([]byte(publicPem)),
	)

	cipherText, err := cto.Encrypt(plainText)
	if err != nil {
		return err
	}

	cto = rsa2.NewCrypto(
		rsa2.WithPrivateKey([]byte(privatePem)),
	)
	genPlainText, err := cto.Decrypt(cipherText)
	if err != nil {
		return err
	}

	if strings.Compare(string(plainText), string(genPlainText)) != 0 {
		return fmt.Errorf("plainText not equal to genPlainText")
	}

	return nil
}
