package gateway

import "backstage/common/protocol/inform"

type ForceOfflineReq struct {
	UserId       int64 `json:"user_id"`
	Notification *inform.Notification
}

type ForceOfflineRsp struct {
}
