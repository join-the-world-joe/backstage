package sms

import (
	"backstage/diagnostic"
	"context"
	"testing"
)

func TestSendVerificationCode(t *testing.T) {
	diagnostic.SetupLogger()
	diagnostic.SetupRegistry()
	req := &SendVerificationCodeReq{CountryCode: "86", PhoneNumber: "18629300170", Behavior: "SignIn"}
	rsp := &SendVerificationCodeRsp{}
	err := SendVerificationCode(context.Background(), req, rsp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("rsp: ", rsp)
}
