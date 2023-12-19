package admin

import "encoding/json"

type InsertRecordOfADOfCampingReq struct {
	UserId              int64
	AdvertisementIdList []int64 `json:"advertisement_id_list"`
}

type InsertRecordOfADOfCampingRsp struct {
	Code int             `json:"code"`
	Body json.RawMessage `json:"body"`
}
