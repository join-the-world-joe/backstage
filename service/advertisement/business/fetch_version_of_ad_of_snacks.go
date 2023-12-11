package business

import (
	"backstage/common/code"
	"backstage/common/db/mysql/backend/version_of_ad_of_deals"
	"backstage/common/protocol/advertisement"
	"context"
	"encoding/json"
)

type OutputOfVersionOfADOfSnacks struct {
	VersionOfADOfSnacks int64 `json:"version_of_ad_of_snacks"`
}

func FetchVersionOfADOfSnacks(ctx context.Context, req *advertisement.FetchVersionOfADOfSnacksReq, rsp *advertisement.FetchVersionOfADOfSnacksRsp) error {
	id, err := version_of_ad_of_deals.GetMaxId()
	if err != nil {
		rsp.Code = code.DatabaseFailure
		return nil
	}

	output := &OutputOfVersionOfADOfSnacks{
		VersionOfADOfSnacks: id,
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
