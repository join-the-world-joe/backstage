package advertisement

import "encoding/json"

type FetchRecordsOfADOfDealsReq struct {
	UserId              int64
	AdvertisementIdList []int64 `json:"advertisement_id_list"`
}

type FetchRecordsOfADOfDealsRsp struct {
	Code int             `json:"code"`
	Body json.RawMessage `json:"body"`
}
