package oss

import "encoding/json"

type RemoveListOfObjectFileReq struct {
	UserId           int64
	ListOfObjectFile []string `json:"list_of_object_file"`
}

type RemoveListOfObjectFileRsp struct {
	Code int             `json:"code"`
	Body json.RawMessage `json:"body"`
}
