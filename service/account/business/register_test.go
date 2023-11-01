package business

import (
	"backstage/common/service/account"
	"backstage/diagnostic"
	"context"
	"testing"
)

func TestRegister(t *testing.T) {
	countryCode := "86"
	phoneNumber := "11337"
	code := "9092"
	diagnostic.SetupLogger()
	diagnostic.SetupRegistry()
	req := &account.RegisterReq{CountryCode: countryCode, PhoneNumber: phoneNumber, VerificationCode: code}
	rsp := &account.RegisterRsp{}
	err := account.Register(context.Background(), req, rsp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("rsp: ", rsp)
}
