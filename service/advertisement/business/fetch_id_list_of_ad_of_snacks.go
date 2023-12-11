package business

import (
	"backstage/common/code"
	"backstage/common/db/mysql/backend/ad_of_snacks"
	"backstage/common/protocol/advertisement"
	"backstage/global/log"
	"context"
	"encoding/json"
)

type OutputOfIDListOfADOfSnacks struct {
	VersionOfADOfSnacks int64   `json:"version_of_ad_of_snacks"`
	IdListOfADOfSnacks  []int64 `json:"id_list_of_ad_of_snacks"`
}

func FetchIdListOfADOfSnacks(ctx context.Context, req *advertisement.FetchIdListOfADOfSnacksReq, rsp *advertisement.FetchIdListOfADOfSnacksRsp) error {
	model, err := ad_of_snacks.GetLatestVersionModel()
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
	output := &OutputOfIDListOfADOfHots{
		VersionOfADOfHots: model.Id,
		IdListOfADOfHots:  idList,
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
