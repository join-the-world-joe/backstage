package advertisement

import "encoding/json"

type FetchADOfHotDealsReq struct {
	UserId int64
}

type FetchADOfHotDealsRsp struct {
	Code    int             `json:"code"`
	Version int             `json:"version"`
	Body    json.RawMessage `json:"body"`
}
