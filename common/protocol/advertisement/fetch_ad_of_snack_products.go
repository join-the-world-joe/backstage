package advertisement

import "encoding/json"

type FetchADOfSnackProductsReq struct {
	UserId int64
}

type FetchADOfSnackProductsRsp struct {
	Code    int             `json:"code"`
	Version int             `json:"version"`
	Body    json.RawMessage `json:"body"`
}
