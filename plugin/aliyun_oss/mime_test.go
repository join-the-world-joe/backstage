package aliyun_oss

import (
	"mime"
	"testing"
)

func TestMIME(t *testing.T) {
	file := "1.mp4"
	contentType := mime.TypeByExtension(file)
	t.Log("Content Type: ", contentType)
}
