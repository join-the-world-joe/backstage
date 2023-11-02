package gateway

type PingReq struct {
	Message string `json:"message"`
}

type PongRsp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
