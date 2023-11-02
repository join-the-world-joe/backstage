package admin

import "encoding/json"

type FetchRecordsOfAdvertisementReq struct {
	UserId              int64
	AdvertisementIdList []int64 `json:"advertisement_id_list"`
}

type FetchRecordsOfAdvertisementRsp struct {
	Code int             `json:"code"`
	Body json.RawMessage `json:"body"`
}
