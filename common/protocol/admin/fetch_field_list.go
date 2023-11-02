package admin

import "encoding/json"

type FetchFieldListReq struct {
	RoleList []string `json:"role_list"`
}

type FetchFieldListRsp struct {
	Code int             `json:"code"`
	Body json.RawMessage `json:"body"`
}
