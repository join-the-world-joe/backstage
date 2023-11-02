package admin

import "encoding/json"

type FetchTrackListOfConditionReq struct {
	UserId     int64  `json:"user_id"`
	Operator   []byte `json:"operator"`
	Behavior   int    `json:"behavior"`
	Permission string `json:"permission"`
	Major      string `json:"major"`
	Minor      string `json:"minor"`
	Begin      int64  `json:"begin"`
	End        int64  `json:"end"`
}

type FetchTrackListOfConditionRsp struct {
	Code int             `json:"code"`
	Body json.RawMessage `json:"body"`
}
