package advertisement

import "encoding/json"

type FetchRecordsOfADOfCampingReq struct {
	UserId              int64
	AdvertisementIdList []int64 `json:"advertisement_id_list"`
}

type FetchRecordsOfADOfCampingRsp struct {
	Code int             `json:"code"`
	Body json.RawMessage `json:"body"`
}
