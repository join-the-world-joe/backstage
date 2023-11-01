package business

import (
	"backstage/common/service/account"
	"backstage/diagnostic"
	"context"
	"testing"
)

func TestLogin1(t *testing.T) {
	countryCode := "86"
	phoneNumber := "11337"
	code := "9092"
	diagnostic.SetupLogger()
	diagnostic.SetupRegistry()
	req := &account.LoginReq{CountryCode: countryCode, PhoneNumber: phoneNumber, VerificationCode: code}
	rsp := &account.LoginRsp{}
	err := account.Login(context.Background(), req, rsp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("rsp: ", rsp)
}

func TestLogin2(t *testing.T) {
	countryCode := "86"
	phoneNumber := "11337"
	token := "34281e99-8350-46b2-82ca-974994e6bc07"
	diagnostic.SetupLogger()
	diagnostic.SetupRegistry()
	req := &account.LoginReq{CountryCode: countryCode, PhoneNumber: phoneNumber, Token: token}
	rsp := &account.LoginRsp{}
	err := account.Login(context.Background(), req, rsp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("rsp: ", rsp)
}
