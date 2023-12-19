package advertisement

import "encoding/json"

type FetchIdListOfADOfCarouselReq struct {
	Behavior int `json:"behavior"` // 0 - with status concern(ignore zero status records), for client side; 1 - without status concern for backend
	UserId   int64
}

type FetchIdListOfADOfCarouselRsp struct {
	Code int             `json:"code"`
	Body json.RawMessage `json:"body"`
}
