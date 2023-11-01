package backend

import (
	"context"
	"go-micro-framework/diagnostic"
	"go-micro-framework/plugin/crypto/rsa"
	"go-micro-framework/utils/bcrypt"
	"testing"
)

func TestVerificationCodeSignIn(t *testing.T) {
	diagnostic.SetupLogger()
	diagnostic.SetupRegistry()
	req := &SignInReq{
		CountryCode:      "86",
		PhoneNumber:      "18629300170",
		VerificationCode: "4053",
	}
	rsp := &SignInRsp{}
	err := SignIn(context.Background(), req, rsp)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("rsp: ", rsp)
}

func TestPasswordSignIn(t *testing.T) {
	PublicKey := `
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDZsfv1qscqYdy4vY+P4e3cAtmv
ppXQcRvrF1cB4drkv0haU24Y7m5qYtT52Kr539RdbKKdLAM6s20lWy7+5C0Dgacd
wYWd/7PeCELyEipZJL07Vro7Ate8Bfjya+wltGK9+XNUIHiumUKULW4KDx21+1NL
AUeJ6PeW+DAkmJWF6QIDAQAB
-----END PUBLIC KEY-----
`
	passwordPlainText := "1234561"
	bPassword, err := bcrypt.PasswordHash(passwordPlainText)
	if err != nil {
		t.Fatal(err)
	}
	diagnostic.SetupLogger()
	diagnostic.SetupRegistry()
	crypto := rsa.NewCrypto(
		rsa.WithPublicKey([]byte(PublicKey)),
	)
	encryptedPasswordBytes, err := crypto.Encrypt([]byte(bPassword))
	if err != nil {
		t.Fatal(err)
	}

	req := &SignInReq{
		CountryCode: "86",
		PhoneNumber: "18629300170",
		Password:    encryptedPasswordBytes,
	}
	rsp := &SignInRsp{}
	err = SignIn(context.Background(), req, rsp)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("rsp: ", rsp)
	t.Log("Password Plain Text: ", passwordPlainText)
	t.Log("Bcrypt Password: ", []byte(bPassword))
	t.Log("Encrypted Password: ", encryptedPasswordBytes)
}
