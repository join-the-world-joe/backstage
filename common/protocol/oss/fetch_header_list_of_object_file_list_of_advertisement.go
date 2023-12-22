package oss

import "encoding/json"

type FetchHeaderListOfObjectFileListOfAdvertisementReq struct {
	UserId          int64
	AdvertisementId int64    `json:"advertisement_id"`
	NameListOfFile  []string `json:"name_list_of_file"`
}

type FetchHeaderListOfObjectFileListOfAdvertisementRsp struct {
	Code int             `json:"code"`
	Body json.RawMessage `json:"body"`
}
