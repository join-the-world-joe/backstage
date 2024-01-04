package oss

import "encoding/json"

type VerifyObjectFileListReq struct {
	UserId               int64
	OSSFolder            string   `json:"oss_folder"`
	NameListOfObjectFile []string `json:"name_list_of_object_file"`
}

type VerifyObjectFileListRsp struct {
	Code int             `json:"code"`
	Body json.RawMessage `json:"body"`
}
