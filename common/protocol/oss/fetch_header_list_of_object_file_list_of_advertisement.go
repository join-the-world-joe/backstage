package oss

import "encoding/json"

type FetchHeaderListOfObjectFileListOfAdvertisementReq struct {
	UserId         int64
	OSSFolder      string   `json:"oss_folder"`
	NameListOfFile []string `json:"name_list_of_file"`
}

type FetchHeaderListOfObjectFileListOfAdvertisementRsp struct {
	Code int             `json:"code"`
	Body json.RawMessage `json:"body"`
}
