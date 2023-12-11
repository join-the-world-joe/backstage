package business

import (
	"backstage/common/code"
	"backstage/common/db/mysql/backend/ad_of_deals"
	"backstage/common/protocol/advertisement"
	"backstage/global/log"
	"context"
	"encoding/json"
)

type OutputOfIDListOfADOfDeals struct {
	VersionOfADOfDeals int64   `json:"version_of_ad_of_deals"`
	IdListOfADOfDeals  []int64 `json:"id_list_of_ad_of_deals"`
}

func FetchIdListOfADOfDeals(ctx context.Context, req *advertisement.FetchIdListOfADOfDealsReq, rsp *advertisement.FetchIdListOfADOfDealsRsp) error {
	model, err := ad_of_deals.GetLatestVersionModel()
	if err != nil {
		log.Error("ad_of_deals.GetLatestVersionModel failure, err: ", err)
		rsp.Code = code.DatabaseFailure
		return nil
	}
	idList := []int64{}
	err = json.Unmarshal([]byte(model.AdvertisementIdList), &idList)
	if err != nil {
		log.Error("json.Unmarshal failure, err: ", err)
		rsp.Code = code.InternalError
		return nil
	}
	output := &OutputOfIDListOfADOfDeals{
		VersionOfADOfDeals: model.Id,
		IdListOfADOfDeals:  idList,
	}

	bytes, err := json.Marshal(output)
	if err != nil {
		log.Error("json.Marshal failure, err: ", err)
		rsp.Code = code.InvalidData
		return nil
	}
	rsp.Body = bytes
	rsp.Code = code.Success
	return nil
}
