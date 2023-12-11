package business

import (
	"backstage/common/code"
	"backstage/common/db/mysql/backend/version_of_ad_of_snacks"
	"backstage/common/protocol/advertisement"
	"context"
	"encoding/json"
)

type OutputOfVersionOfADOfHots struct {
	VersionOfADOfHots int64 `json:"version_of_ad_of_hots"`
}

func FetchVersionOfADOfHots(ctx context.Context, req *advertisement.FetchVersionOfADOfHotsReq, rsp *advertisement.FetchVersionOfADOfHotsRsp) error {
	id, err := version_of_ad_of_snacks.GetMaxId()
	if err != nil {
		rsp.Code = code.DatabaseFailure
		return nil
	}

	output := &OutputOfVersionOfADOfHots{
		VersionOfADOfHots: id,
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
