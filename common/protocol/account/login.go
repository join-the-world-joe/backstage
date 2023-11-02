package account

type LoginReq struct {
	Behavior         int    `json:"behavior"`
	Account          string `json:"account"`
	Email            string `json:"email"`
	PhoneNumber      string `json:"phone_number"`
	CountryCode      string `json:"country_code"`
	Password         []byte `json:"password"` // rsa, bcrypt
	VerificationCode int32  `json:"verification_code,string"`
	UserId           int64  `json:"user_id"`
	MemberId         string `json:"member_id"`
}

type LoginRsp struct {
	Code     int    `json:"code"`
	MemberId string `json:"member_id"`
	UserId   int64  `json:"user_id"`
	Secret   string `json:"secret"`
}
