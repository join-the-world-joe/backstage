package advertisement

import "encoding/json"

type FetchRecordsOfAdvertisementReq struct {
	AdvertisementIdList []int64 `json:"advertisement_id_list"`
}

type FetchRecordsOfAdvertisementRsp struct {
	Code int             `json:"code"`
	Body json.RawMessage `json:"body"`
}
