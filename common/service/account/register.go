package account

type RegisterReq struct {
	Account          string `json:"account"`
	Email            string `json:"email"`
	PhoneNumber      string `json:"phone_number"`
	CountryCode      string `json:"country_code"`
	Password         []byte `json:"password"`
	InvitationCode   string `json:"invitation_code"`
	VerificationCode string `json:"verification_code"`
}

type RegisterRsp struct {
	Code int `json:"code"`
}
