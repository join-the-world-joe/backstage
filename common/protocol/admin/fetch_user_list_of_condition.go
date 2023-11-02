package admin

import "encoding/json"

type FetchUserListOfConditionReq struct {
	Id          int64  `json:"id"`
	UserId      int64  `json:"user_id"`
	Name        []byte `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Behavior    int    `json:"behavior"`
}

type FetchUserListOfConditionRsp struct {
	Code int             `json:"code"`
	Body json.RawMessage `json:"body"`
}
