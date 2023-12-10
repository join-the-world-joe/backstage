package advertisement

import "encoding/json"

type FetchRecordsOfADOfCarouselReq struct {
	UserId              int64
	AdvertisementIdList []int64 `json:"advertisement_id_list"`
}

type FetchRecordsOfADOfCarouselRsp struct {
	Code int             `json:"code"`
	Body json.RawMessage `json:"body"`
}
