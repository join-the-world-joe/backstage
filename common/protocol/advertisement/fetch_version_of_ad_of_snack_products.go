package advertisement

import "encoding/json"

type FetchVersionOfADOfSnackProductsReq struct {
	UserId int64
}

type FetchVersionOfADOfSnackProductsRsp struct {
	Code    int             `json:"code"`
	Version int             `json:"version"`
	Body    json.RawMessage `json:"body"`
}
