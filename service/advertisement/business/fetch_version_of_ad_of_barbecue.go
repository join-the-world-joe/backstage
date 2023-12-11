package business

import (
	"backstage/common/code"
	"backstage/common/db/mysql/backend/version_of_ad_of_barbecue"
	"backstage/common/protocol/advertisement"
	"context"
	"encoding/json"
)

type OutputOfVersionOfADOfBarbecue struct {
	VersionOfADOfBarbecue int64 `json:"version_of_ad_of_barbecue"`
}

func FetchVersionOfADOfBarbecue(ctx context.Context, req *advertisement.FetchVersionOfADOfBarbecueReq, rsp *advertisement.FetchVersionOfADOfBarbecueRsp) error {
	id, err := version_of_ad_of_barbecue.GetMaxId()
	if err != nil {
		rsp.Code = code.DatabaseFailure
		return nil
	}

	output := &OutputOfVersionOfADOfBarbecue{
		VersionOfADOfBarbecue: id,
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
