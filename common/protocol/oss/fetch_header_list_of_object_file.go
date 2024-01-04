package oss

import "encoding/json"

type FetchHeaderListOfObjectFileListReq struct {
	UserId         int64
	NameListOfFile []string `json:"name_list_of_file"`
}

type FetchHeaderListOfObjectFileListRsp struct {
	Code int             `json:"code"`
	Body json.RawMessage `json:"body"`
}
