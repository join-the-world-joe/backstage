package advertisement

import "encoding/json"

type FetchADOfBBQProductsReq struct {
	UserId int64
}

type FetchADOfBBQProductsRsp struct {
	Code    int             `json:"code"`
	Version int             `json:"version"`
	Body    json.RawMessage `json:"body"`
}
