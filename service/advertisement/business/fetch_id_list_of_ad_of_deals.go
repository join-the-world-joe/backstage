package business

import (
	"backstage/common/code"
	"backstage/common/db/mysql/backend/ad_of_deals"
	"backstage/common/db/mysql/backend/version_of_ad_of_deals"
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
	version, err := version_of_ad_of_deals.GetMaxId()
	if err != nil {
		log.Error("version_of_ad_of_deals.GetMaxId failure, err: ", err)
		rsp.Code = code.DatabaseFailure
		return nil
	}

	model, err := ad_of_deals.GetModelByVersion(version)
	if err != nil {
		log.Error("ad_of_deals.GetLatestVersionModel failure, err: ", err)
		rsp.Code = code.DatabaseFailure
		return nil
	}

	output := &OutputOfIDListOfADOfDeals{
		VersionOfADOfDeals: model.Version,
		IdListOfADOfDeals:  []int64{},
	}

	if len(model.AdvertisementIdList) > 0 {
		idList := []int64{}
		err = json.Unmarshal([]byte(model.AdvertisementIdList), &idList)
		if err != nil {
			log.Error("json.Unmarshal failure, err: ", err)
			rsp.Code = code.InternalError
			return nil
		}
		output.IdListOfADOfDeals = idList
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
