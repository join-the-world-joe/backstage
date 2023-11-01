package backend

import "encoding/json"

type FetchAttributeListReq struct {
	Role string `json:"role"`
}

type FetchAttributeListRsp struct {
	Code int             `json:"code"`
	Body json.RawMessage `json:"body"`
}
