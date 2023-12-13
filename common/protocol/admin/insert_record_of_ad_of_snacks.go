package admin

import "encoding/json"

type InsertRecordOfADOfSnacksReq struct {
	UserId              int64
	AdvertisementIdList []int64 `json:"advertisement_id_list"`
}

type InsertRecordOfADOfSnacksRsp struct {
	Code int             `json:"code"`
	Body json.RawMessage `json:"body"`
}
