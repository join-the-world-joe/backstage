package advertisement

import "encoding/json"

type FetchIdListOfADOfHotsReq struct {
	UserId int64
}

type FetchIdListOfADOfHotsRsp struct {
	Code int             `json:"code"`
	Body json.RawMessage `json:"body"`
}
