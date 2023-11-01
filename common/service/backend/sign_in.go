package backend

type SignInReq struct {
	Type             string `json:"type"`
	Account          string `json:"account"`
	Email            string `json:"email"`
	PhoneNumber      string `json:"phone_number"`
	CountryCode      string `json:"country_code"`
	Password         []byte `json:"password"` // rsa, bcrypt
	Token            string `json:"token"`
	VerificationCode string `json:"verification_code"`
}

type SignInRsp struct {
	Code   int    `json:"code"`
	UserId int64  `json:"user_id"`
	Role   string `json:"role"`
}
