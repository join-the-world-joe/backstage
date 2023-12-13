package admin

import "encoding/json"

type InsertRecordOfADOfCarouselReq struct {
	UserId              int64
	AdvertisementIdList []int64 `json:"advertisement_id_list"`
}

type InsertRecordOfADOfCarouselRsp struct {
	Code int             `json:"code"`
	Body json.RawMessage `json:"body"`
}
