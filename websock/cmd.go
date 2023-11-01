package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"go-micro-framework/common/major"
	"go-micro-framework/common/service/account"
	"go-micro-framework/common/service/generic"
	"go-micro-framework/common/service/sms"
	"time"
)

var countryCode = "86"
var phoneNumber = "223344"
var behavior = "Register"
var code = "1111"
var token_ = "da62c3a3-f2a2-4b72-8592-7309c105aead"

func auth(conn *websocket.Conn) {
	ma := major.Generic
	mi := generic.AuthReq
	req := &generic.AuthenticateReq{
		Email: "wenxing.cn@gmail.com",
	}
	content, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	send(conn, ma, mi, string(content))
}

func ping(conn *websocket.Conn) {
	count++
	ma := major.Generic
	mi := generic.Ping
	message := fmt.Sprintf("Hello, count %v", count)
	send(conn, ma, mi, message)
}

func repeat(conn *websocket.Conn) {
	for i := 0; i <= 1000000; i++ {
		time.Sleep(time.Millisecond * 500)
		ping(conn)
	}
}

func register(conn *websocket.Conn) {
	ma := major.Account
	mi := account.RegisterReq_
	body := account.RegisterReq{
		CountryCode:      countryCode,
		PhoneNumber:      phoneNumber,
		VerificationCode: code,
	}
	message, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}
	send(conn, ma, mi, string(message))
}

func sendVerificationCodeRegister(conn *websocket.Conn) {
	ma := major.SMS
	mi := sms.SendVerificationCodeReq_
	body := sms.SendVerificationCodeReq{
		CountryCode: countryCode,
		PhoneNumber: phoneNumber,
		Behavior:    "Register",
	}
	message, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}
	send(conn, ma, mi, string(message))
}

func sendVerificationCodeLogin(conn *websocket.Conn) {
	ma := major.SMS
	mi := sms.SendVerificationCodeReq_
	body := sms.SendVerificationCodeReq{
		CountryCode: countryCode,
		PhoneNumber: phoneNumber,
		Behavior:    "Login",
	}
	message, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}
	send(conn, ma, mi, string(message))
}

func login(conn *websocket.Conn) {
	ma := major.Account
	mi := account.LoginReq_
	req := &account.LoginReq{
		CountryCode:      countryCode,
		PhoneNumber:      phoneNumber,
		VerificationCode: code,
	}
	content, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	send(conn, ma, mi, string(content))
}

func token(conn *websocket.Conn) {
	ma := major.Account
	mi := account.LoginReq_
	req := &account.LoginReq{
		CountryCode: countryCode,
		PhoneNumber: phoneNumber,
		Token:       token_,
	}
	content, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	send(conn, ma, mi, string(content))
}
