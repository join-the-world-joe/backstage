package qrcode

import "github.com/skip2/go-qrcode"

func Encode(any string, size int) ([]byte, error) {
	return qrcode.Encode(any, qrcode.Medium, size)
}
