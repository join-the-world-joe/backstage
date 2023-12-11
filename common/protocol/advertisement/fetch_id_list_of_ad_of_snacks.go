package advertisement

import "encoding/json"

type FetchIdListOfADOfSnacksReq struct {
	UserId int64
}

type FetchIdListOfADOfSnacksRsp struct {
	Code int             `json:"code"`
	Body json.RawMessage `json:"body"`
}
