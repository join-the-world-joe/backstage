package business

import (
	"backstage/common/code"
	"backstage/common/db/mysql/backend/version_of_ad_of_deals"
	"backstage/common/protocol/advertisement"
	"context"
	"encoding/json"
)

type OutputOfVersionOfADOfDeals struct {
	VersionOfADOfDeals int64 `json:"version_of_ad_of_deals"`
}

func FetchVersionOfADOfDeals(ctx context.Context, req *advertisement.FetchVersionOfADOfDealsReq, rsp *advertisement.FetchVersionOfADOfDealsRsp) error {
	id, err := version_of_ad_of_deals.GetMaxId()
	if err != nil {
		rsp.Code = code.DatabaseFailure
		return nil
	}

	output := &OutputOfVersionOfADOfDeals{
		VersionOfADOfDeals: id,
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
