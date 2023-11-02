package account

type LogoutReq struct {
	Token string `json:"token"`
}

type LogoutRsp struct {
	Code int `json:"code"`
}
