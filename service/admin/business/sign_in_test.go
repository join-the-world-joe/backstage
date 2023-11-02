package business

import (
	"backstage/common/cache/string/token"
	"backstage/common/protocol/admin"
	"backstage/diagnostic"
	"backstage/global/crypto"
	google_authentictor "backstage/utils/google_authenticator"
	"encoding/json"
	"github.com/google/uuid"
	"testing"
)

var request = `
{"email":"","member_id":"","account":"admin","behavior":4,"password":[153,9,250,71,84,24,188,191,231,54,218,60,252,161,132,139,79,95,169,121,204,28,96,234,119,152,207,143,115,227,37,36,19,112,222,78,119,197,17,158,237,156,36,88,95,249,203,66,222,157,92,166,199,20,68,160,38,224,116,19,129,69,192,242,69,30,171,124,252,209,24,118,77,62,227,173,114,101,1,22,232,60,113,140,227,205,43,218,204,114,70,168,67,229,83,110,22,104,228,255,114,213,150,210,131,223,223,129,230,58,32,202,215,242,90,34,13,215,99,159,27,137,152,246,196,52,112,6],"phone_number":"","country_code":"","verification_code":0,"user_id":0}
`

func TestDecodePasswordFromRawRequest(t *testing.T) {
	req := &admin.SignInReq{}
	err := json.Unmarshal([]byte(request), req)
	if err != nil {
		t.Fatal(err)
	}
	//t.Log("req: ", req)
	bytes, err := crypto.RSADecrypt(req.Password)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("password: ", string(bytes))
}

func TestEncryptToken(t *testing.T) {
	token, secret, err := encryptToken(1)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("token: ", token)
	t.Log("secret: ", secret)
}

//func TestDecryptToken(t *testing.T) {
//	token, secret, err := encryptToken(1)
//	if err != nil {
//		t.Fatal(err)
//	}
//	uid, secret, err := decryptToken(token)
//	if err != nil {
//		t.Fatal(err)
//	}
//	t.Log("Id: ", uid)
//	t.Log("secret: ", secret)
//}

func TestCompareToken(t *testing.T) {
	// d481c866-0644-4d05-8176-73e0624d75e6
	// 5CAA77BBAA6G6DAB
	userId := int64(1)
	unique := "d481c866-0644-4d05-8176-73e0624d75e6"
	secret := "5CAA77BBAA6G6DAB"
	code, _ := google_authentictor.GetGoogleCode(secret)
	diagnostic.SetupLogger()
	diagnostic.SetupCache()
	b, err := compareToken(userId, unique, code)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("b: ", b)
}

func TestCreateToken(t *testing.T) {
	diagnostic.SetupLogger()
	diagnostic.SetupCache()
	userId := int64(1)
	tokenValue, secret, err := encryptToken(userId)
	if err != nil {
		t.Fatal(err)
	} else {
		tokenUnique := uuid.New().String()
		err = token.Create(tokenUnique, tokenValue)
		if err != nil {
			t.Fatal(err)
		}
		t.Log("token unique: ", tokenUnique)
		t.Log("token secret: ", secret)
	}
}
