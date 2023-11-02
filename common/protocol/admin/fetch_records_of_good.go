package admin

import "encoding/json"

type FetchRecordsOfGoodReq struct {
	UserId        int64
	ProductIdList []int64 `json:"product_id_list"`
}

type FetchRecordsOfGoodRsp struct {
	Code int             `json:"code"`
	Body json.RawMessage `json:"body"`
}
