package aes

import (
	"backstage/abstract/crypto"
	"fmt"
	"github.com/forgoer/openssl"
)

const (
	Name  = "AES Encryption"
	PKCS7 = "PKCS7"
	PKCS5 = "PKCS5"
)

var padding = map[string]struct{}{"ZERO": {}, PKCS5: {}, PKCS7: {}}

type _crypto struct {
	opts *Options
}

func NewCrypto(opts ...Option) crypto.Crypto {
	options := Options{padding: PKCS7}

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

func (p *_crypto) Encrypt(secret []byte) (cipherText []byte, err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("painc, secret: %v", secret)
			if e, ok := p.(error); ok {
				err = e
			}
		}
	}()
	return openssl.AesCBCEncrypt(secret, []byte(p.opts.key), []byte(p.opts.iv), p.opts.padding)
}

func (p *_crypto) Decrypt(cipherText []byte) (plainText []byte, err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("painc, cipherText: %v", cipherText)
			if e, ok := p.(error); ok {
				err = e
			}
		}
	}()
	return openssl.AesCBCDecrypt(cipherText, []byte(p.opts.key), []byte(p.opts.iv), p.opts.padding)
}

//
//func Echo(conn *websocket.Conn) {
//	major := "1"
//	minor := "1"
//	message := "hello, I'm John!"
//
//	packet := &payload.PacketClient{
//		Header: &payload.Header{
//			Major: major,
//			Minor: minor,
//		},
//		Body: []byte(message),
//	}
//
//	bytes, err := packet.ToBytes()
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	err = conn.WriteMessage(websocket.BinaryMessage, bytes)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//}
