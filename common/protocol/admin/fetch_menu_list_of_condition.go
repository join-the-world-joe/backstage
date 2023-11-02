package admin

import "encoding/json"

type FetchMenuListOfConditionReq struct {
	Id       int64    `json:"id"`
	RoleList []string `json:"role_list"`
	UserId   int64    `json:"user_id"`
	Menu     string   `json:"menu"`
	Parent   string   `json:"parent"`
	Behavior int      `json:"behavior"`
}

type FetchMenuListOfConditionRsp struct {
	Code int             `json:"code"`
	Body json.RawMessage `json:"body"`
}
