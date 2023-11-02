package advertisement

import "encoding/json"

type FetchADOfCarouselReq struct {
	UserId int64
}

type FetchADOfCarouselRsp struct {
	Code    int             `json:"code"`
	Version int             `json:"version"`
	Body    json.RawMessage `json:"body"`
}
