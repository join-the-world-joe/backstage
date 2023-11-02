package base64

import "encoding/base64"

func Encode(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}

func Decode(src string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(src)
}
