package admin

import "encoding/json"

type FetchRoleListOfConditionReq struct {
	Id       int64    `json:"id"`
	UserId   int64    `json:"user_id"`
	Behavior int      `json:"behavior"`
	RoleList [][]byte `json:"role_list"`
}

type FetchRoleListOfConditionRsp struct {
	Code int             `json:"code"`
	Body json.RawMessage `json:"body"`
}
