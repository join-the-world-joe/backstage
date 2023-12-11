package advertisement

import "encoding/json"

type FetchRecordsOfADOfHotsReq struct {
	UserId              int64
	AdvertisementIdList []int64 `json:"advertisement_id_list"`
}

type FetchRecordsOfADOfHotsRsp struct {
	Code int             `json:"code"`
	Body json.RawMessage `json:"body"`
}
