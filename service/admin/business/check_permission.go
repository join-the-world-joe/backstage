package business

import (
	"backstage/common/code"
	"backstage/common/protocol/admin"
	"context"
	"encoding/json"
)

type OutputOfCheckPermission struct {
	Major int `json:"major"`
	Minor int `json:"minor"`
}

func CheckPermission(ctx context.Context, req *admin.CheckPermissionReq, rsp *admin.CheckPermissionRsp) error {
	var err error
	rsp.Body, err = json.Marshal(&OutputOfCheckPermission{
		Major: req.Major,
		Minor: req.Minor,
	})
	if err != nil {
		rsp.Code = code.InternalError
		return nil
	}
	if !hasPermission(req.Major, req.Minor, req.UserId) {
		rsp.Code = code.AccessDenied
		return nil
	}

	rsp.Code = code.Success
	return nil
}
