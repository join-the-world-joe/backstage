package admin

import "encoding/json"

type CheckPermissionReq struct {
	UserId int64 `json:"user_id"`
	Major  int   `json:"major"`
	Minor  int   `json:"minor"`
}

type CheckPermissionRsp struct {
	Code int             `json:"code"`
	Body json.RawMessage `json:"body"`
}
