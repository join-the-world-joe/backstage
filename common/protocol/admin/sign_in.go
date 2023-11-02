package admin

type SignInReq struct {
	Behavior         int    `json:"behavior"`
	Account          string `json:"account"`
	Email            string `json:"email"`
	PhoneNumber      string `json:"phone_number"`
	CountryCode      string `json:"country_code"`
	Password         []byte `json:"password"` // rsa, bcrypt
	VerificationCode int32  `json:"verification_code"`
	UserId           int64  `json:"user_id"`
	MemberId         string `json:"member_id"`
}

type SignInRsp struct {
	Code     int    `json:"code"`
	UserId   int64  `json:"user_id"`
	Name     string `json:"name"`
	Secret   string `json:"secret"`
	MemberId string `json:"member_id"`
}
