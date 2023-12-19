package advertisement

import "encoding/json"

type FetchVersionOfADOfCampingReq struct {
	UserId int64
}

type FetchVersionOfADOfCampingRsp struct {
	Code int             `json:"code"`
	Body json.RawMessage `json:"body"`
}
