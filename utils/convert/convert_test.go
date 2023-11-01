package convert

import (
	"testing"
)

func TestBytes2StringArray(t *testing.T) {
	t.Log("str: ", Bytes2StringArray([]byte(`123456`)))
}
