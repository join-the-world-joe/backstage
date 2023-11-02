package sms

type SendVerificationCodeReq struct {
	Behavior    string `json:"Behavior"`
	CountryCode string `json:"country_code"`
	PhoneNumber string `json:"phone_number"`
}

type SendVerificationCodeRsp struct {
	Code int `json:"code"`
}
