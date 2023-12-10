package advertisement

import "encoding/json"

type FetchIdListOfADOfCarouselReq struct {
	UserId int64
}

type FetchIdListOfADOfCarouselRsp struct {
	Code int             `json:"code"`
	Body json.RawMessage `json:"body"`
}
