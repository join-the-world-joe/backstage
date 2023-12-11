package business

import (
	"backstage/common/code"
	"backstage/common/db/mysql/backend/ad_of_barbecue"
	"backstage/common/protocol/advertisement"
	"backstage/global/log"
	"context"
	"encoding/json"
)

type OutputOfIDListOfADOfBarbecue struct {
	VersionOfADOfBarbecue int64   `json:"version_of_ad_of_barbecue"`
	IdListOfADOfBarbecue  []int64 `json:"id_list_of_ad_of_barbecue"`
}

func FetchIdListOfADOfBarbecue(ctx context.Context, req *advertisement.FetchIdListOfADOfBarbecueReq, rsp *advertisement.FetchIdListOfADOfBarbecueRsp) error {
	model, err := ad_of_barbecue.GetLatestVersionModel()
	if err != nil {
		log.Error("ad_of_hots.GetLatestVersionModel failure, err: ", err)
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
	output := &OutputOfIDListOfADOfBarbecue{
		VersionOfADOfBarbecue: model.Id,
		IdListOfADOfBarbecue:  idList,
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
