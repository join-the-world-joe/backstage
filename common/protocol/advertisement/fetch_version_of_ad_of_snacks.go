package advertisement

import "encoding/json"

type FetchVersionOfADOfSnacksReq struct {
	UserId int64
}

type FetchVersionOfADOfSnacksRsp struct {
	Code int             `json:"code"`
	Body json.RawMessage `json:"body"`
}
