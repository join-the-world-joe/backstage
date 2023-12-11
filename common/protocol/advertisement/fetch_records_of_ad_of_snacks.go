package advertisement

import "encoding/json"

type FetchRecordsOfADOfSnacksReq struct {
	UserId              int64
	AdvertisementIdList []int64 `json:"advertisement_id_list"`
}

type FetchRecordsOfADOfSnacksRsp struct {
	Code int             `json:"code"`
	Body json.RawMessage `json:"body"`
}
