package business

import (
	"backstage/common/code"
	"backstage/common/db/mysql/backend/version_of_ad_of_camping"
	"backstage/common/protocol/advertisement"
	"context"
	"encoding/json"
)

type OutputOfVersionOfADOfCamping struct {
	VersionOfADOfCamping int64 `json:"version_of_ad_of_camping"`
}

func FetchVersionOfADOfCamping(ctx context.Context, req *advertisement.FetchVersionOfADOfCampingReq, rsp *advertisement.FetchVersionOfADOfCampingRsp) error {
	id, err := version_of_ad_of_camping.GetMaxId()
	if err != nil {
		rsp.Code = code.DatabaseFailure
		return nil
	}

	output := &OutputOfVersionOfADOfCamping{
		VersionOfADOfCamping: id,
	}

	bytes, err := json.Marshal(output)
	if err != nil {
		rsp.Code = code.InternalError
		return nil
	}

	rsp.Body = bytes
	rsp.Code = code.Success
	return nil
}
