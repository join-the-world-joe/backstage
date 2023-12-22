package business

import (
	"mime"
	"path"
	"testing"
)

func TestDetectMimeType(t *testing.T) {
	input := "0.webp"
	extension := path.Ext(input)
	mimeType := mime.TypeByExtension(extension)
	t.Log("mime: ", mimeType)
}
