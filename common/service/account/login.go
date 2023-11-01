package account

type LoginReq struct {
	Type             string `json:"type"`
	Account          string `json:"account"`
	Email            string `json:"email"`
	PhoneNumber      string `json:"phone_number"`
	CountryCode      string `json:"country_code"`
	Password         []byte `json:"password"`
	Token            string `json:"token"`
	VerificationCode string `json:"verification_code"`
}

type LoginRsp struct {
	Code   int    `json:"code"`
	Token  string `json:"token"`
	UserId int64  `json:"user_id"`
}
