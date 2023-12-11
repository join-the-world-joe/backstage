package advertisement

import "encoding/json"

type FetchVersionOfADOfHotsReq struct {
	UserId int64
}

type FetchVersionOfADOfHotsRsp struct {
	Code int             `json:"code"`
	Body json.RawMessage `json:"body"`
}
