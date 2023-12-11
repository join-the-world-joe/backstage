package advertisement

import "encoding/json"

type FetchVersionOfADOfDealsReq struct {
	UserId int64
}

type FetchVersionOfADOfDealsRsp struct {
	Code int             `json:"code"`
	Body json.RawMessage `json:"body"`
}
