package backend

import "encoding/json"

type FetchAttributeListOfRoleReq struct {
	Role  string `json:"role"`
	Query string `json:"query"`
}

type FetchAttributeListOfRoleRsp struct {
	Code int             `json:"code"`
	Body json.RawMessage `json:"body"`
}
