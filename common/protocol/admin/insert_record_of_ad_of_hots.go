package admin

import "encoding/json"

type InsertRecordOfADOfHotsReq struct {
	UserId              int64
	AdvertisementIdList []int64 `json:"advertisement_id_list"`
	Description         string  `json:"description"`
}

type InsertRecordOfADOfHotsRsp struct {
	Code int             `json:"code"`
	Body json.RawMessage `json:"body"`
}