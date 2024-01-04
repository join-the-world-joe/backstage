package admin

type SoftDeleteRecordOfUserReq struct {
	UserId     int64
	UserIdList []int64 `json:"user_id_list"`
}

type SoftDeleteRecordOfUserRsp struct {
	Code int `json:"code"`
}
