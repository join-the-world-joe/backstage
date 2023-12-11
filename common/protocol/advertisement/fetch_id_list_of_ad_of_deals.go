package advertisement

import "encoding/json"

type FetchIdListOfADOfDealsReq struct {
	UserId int64
}

type FetchIdListOfADOfDealsRsp struct {
	Code int             `json:"code"`
	Body json.RawMessage `json:"body"`
}
