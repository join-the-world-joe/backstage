package google_authentictor

import (
	"fmt"
	"github.com/spf13/cast"
	"testing"
	"time"
)

func TestCreateCode(t *testing.T) {
	secret := CreateGoogleSecret("666666")
	t.Log(secret)
}

func TestGetGoogleCode(t *testing.T) {
	secret := "CBFHB5CB7GB766CA"
	code, remain := GetGoogleCode(secret)
	t.Log("Code: ", code)
	t.Log("Remain: ", remain)
}

func genToken(t *testing.T) {

}

func TestEncryptToken(t *testing.T) {
	userId := 1
	t.Log(fmt.Sprintf("%v-%v", userId, CreateGoogleSecret(cast.ToString(time.Now().Unix()))))
}
