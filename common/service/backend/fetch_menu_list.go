package backend

import "encoding/json"

type FetchMenuListReq struct {
	Role string `json:"role"`
}

type FetchMenuListRsp struct {
	Code int             `json:"code"`
	Body json.RawMessage `json:"body"`
}
