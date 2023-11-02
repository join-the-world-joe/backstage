package admin

type CheckPermissionReq struct {
	Id    int64 `json:"id"`
	Major int   `json:"major"`
	Minor int   `json:"minor"`
}

type CheckPermissionRsp struct {
	Code  int `json:"code"`
	Major int `json:"major"`
	Minor int `json:"minor"`
}
