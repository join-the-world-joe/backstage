package admin

type SoftDeleteUserRecordReq struct {
	Id         int64   `json:"id"`
	UserId     int64   `json:"user_id"`
	UserIdList []int64 `json:"user_id_list"`
}

type SoftDeleteUserRecordRsp struct {
	Code int `json:"code"`
}
