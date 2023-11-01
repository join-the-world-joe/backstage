package aes

import (
	"testing"
)

func TestVerifyKeyPair(t *testing.T) {
	key, iv, err := GenKeyIVPair()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Key: ", key, "Key String: ")
	t.Log("IV: ", iv)

	err = VerifyKeyPair(key, iv)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("done!")
}
