package advertisement

import "encoding/json"

type FetchIdListOfADOfBarbecueReq struct {
	UserId int64
}

type FetchIdListOfADOfBarbecueRsp struct {
	Code int             `json:"code"`
	Body json.RawMessage `json:"body"`
}
