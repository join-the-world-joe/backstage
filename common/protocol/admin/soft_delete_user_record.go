package admin

type SoftDeleteUserRecordReq struct {
	UserId     int64
	UserIdList []int64 `json:"user_id_list"`
}

type SoftDeleteUserRecordRsp struct {
	Code int `json:"code"`
}
