package admin

import "encoding/json"

type FetchFieldListOfConditionReq struct {
	Id       int64  `json:"id"`
	Table    string `json:"table"`
	Field    string `json:"field"` // the name of field
	Behavior int    `json:"behavior"`
	Role     string `json:"role"`
}

type FetchFieldListOfConditionRsp struct {
	Code int             `json:"code"`
	Body json.RawMessage `json:"body"`
}
