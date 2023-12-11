package advertisement

import "encoding/json"

type FetchVersionOfADOfCarouselReq struct {
	UserId int64
}

type FetchVersionOfADOfCarouselRsp struct {
	Code int             `json:"code"`
	Body json.RawMessage `json:"body"`
}
