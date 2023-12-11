package advertisement

import "encoding/json"

type FetchRecordsOfADOfBarbecueReq struct {
	UserId              int64
	AdvertisementIdList []int64 `json:"advertisement_id_list"`
}

type FetchRecordsOfADOfBarbecueRsp struct {
	Code int             `json:"code"`
	Body json.RawMessage `json:"body"`
}
