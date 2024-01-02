package advertisement

import "encoding/json"

type FetchIdListOfAdvertisementReq struct {
	Behavior          int    `json:"behavior"`
	AdvertisementName []byte `json:"advertisement_name"`
}

type FetchIdListOfAdvertisementRsp struct {
	Code int             `json:"code"`
	Body json.RawMessage `json:"body"`
}
