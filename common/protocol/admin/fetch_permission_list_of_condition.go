package admin

import "encoding/json"

type FetchPermissionListOfConditionReq struct {
	Id       int64    `json:"id"`
	Name     string   `json:"name"` // the name of permission
	UserId   int64    `json:"user_id"`
	Behavior int      `json:"behavior"`
	RoleList []string `json:"role_list"`
}

type FetchPermissionListOfConditionRsp struct {
	Code int             `json:"code"`
	Body json.RawMessage `json:"body"`
}
