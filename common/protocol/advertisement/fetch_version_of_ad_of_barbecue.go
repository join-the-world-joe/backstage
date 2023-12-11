package advertisement

import "encoding/json"

type FetchVersionOfADOfBarbecueReq struct {
	UserId int64
}

type FetchVersionOfADOfBarbecueRsp struct {
	Code int             `json:"code"`
	Body json.RawMessage `json:"body"`
}
