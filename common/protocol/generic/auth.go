package generic

type AuthenticateReq struct {
	Email string `json:"email"`
}

type AuthenticateRsp struct {
	UserId int64 `json:"user_id"`
	Code   int   `json:"code"`
}
