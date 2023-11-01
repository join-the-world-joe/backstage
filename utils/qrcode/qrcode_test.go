package qrcode

import (
	"io/ioutil"
	"testing"
)

func TestEncode(t *testing.T) {
	fileName := "test.png"
	png, err := Encode("http://www.google.com", 256)
	if err != nil {
		t.Error(err)
		return
	}

	err = ioutil.WriteFile(fileName, png, 0666)
	if err != nil {
		t.Error(err)
		return
	}
}
