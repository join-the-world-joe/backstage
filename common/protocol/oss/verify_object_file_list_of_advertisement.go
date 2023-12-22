package oss

import "encoding/json"

type VerifyObjectFileListOfAdvertisementReq struct {
	UserId               int64
	AdvertisementId      int64    `json:"advertisement_id"`
	NameListOfObjectFile []string `json:"name_list_of_object_file"`
}

type VerifyObjectFileListOfAdvertisementRsp struct {
	Code int             `json:"code"`
	Body json.RawMessage `json:"body"`
}
