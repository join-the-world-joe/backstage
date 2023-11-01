package rsa

import (
	"backstage/abstract/crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

const (
	Name = "RSA Encryption"
)

type _crypto struct {
	opts *Options
}

func NewCrypto(opts ...Option) crypto.Crypto {
	options := Options{}

	for _, o := range opts {
		o(&options)
	}

	return &_crypto{
		opts: &options,
	}
}

func (p *_crypto) Name() string {
	return Name
}

func (p *_crypto) Encrypt(plainText []byte) ([]byte, error) {
	block, _ := pem.Decode(p.opts.publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, plainText)
}

func (p *_crypto) Decrypt(cipherText []byte) ([]byte, error) {
	block, _ := pem.Decode(p.opts.privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, cipherText)
}
