package admin

import "encoding/json"

type InsertRecordOfADOfDealsReq struct {
	UserId              int64
	AdvertisementIdList []int64 `json:"advertisement_id_list"`
	Description         string  `json:"description"`
}

type InsertRecordOfADOfDealsRsp struct {
	Code int             `json:"code"`
	Body json.RawMessage `json:"body"`
}
