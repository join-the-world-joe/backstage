package base64

import (
	"backstage/utils/convert"
	"testing"
)

func TestEncode(t *testing.T) {
	src := []byte{171, 205, 28, 130, 28, 35, 176, 121, 112, 169, 215, 212, 45, 230, 171, 164, 9, 47, 2, 106, 134, 145, 242, 21, 32, 80, 172, 58, 208, 191, 25, 4, 178, 125, 198, 196, 25, 60, 136, 246, 196, 195, 237, 137, 25, 216, 18, 90, 212, 50, 212, 113, 120, 158, 255, 23, 180, 136, 116, 1, 130, 172, 251, 100}
	t.Log(Encode(src))
}

func TestDecode(t *testing.T) {
	any := "q80cghwjsHlwqdfULearpAkvAmqGkfIVIFCsOtC/GQRHGgW4ij7C/r4wWsLKDLJOfub3CyttGKgzvo+OWUdaE9BkhGp7U1+cTAbQVMlg/ErFpVkw2qBusqftQW3Ike5pChjOd0pTSbzzjVjOHmZnbA=="
	bytes, err := Decode(any)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("bytes", convert.Bytes2StringArray(bytes))
}
