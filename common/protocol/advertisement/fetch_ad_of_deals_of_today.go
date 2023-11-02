package advertisement

import "encoding/json"

type FetchADOfDealsOfTodayReq struct {
	UserId int64
}

type FetchADOfDealsOfTodayRsp struct {
	Code    int             `json:"code"`
	Version int             `json:"version"`
	Body    json.RawMessage `json:"body"`
}
