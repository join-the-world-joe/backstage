package advertisement

import "encoding/json"

type FetchVersionOfADOfBBQProductsReq struct {
	UserId int64
}

type FetchVersionOfADOfBBQProductsRsp struct {
	Code    int             `json:"code"`
	Version int             `json:"version"`
	Body    json.RawMessage `json:"body"`
}
