package business

import (
	"backstage/common/code"
	"backstage/common/db/mysql/backend/ad_of_hots"
	advertisement2 "backstage/common/db/mysql/backend/advertisement"
	"backstage/common/db/mysql/backend/version_of_ad_of_hots"
	"backstage/common/protocol/advertisement"
	"backstage/global/log"
	"context"
	"encoding/json"
)

type OutputOfIDListOfADOfHots struct {
	VersionOfADOfHots int64   `json:"version_of_ad_of_hots"`
	IdListOfADOfHots  []int64 `json:"id_list_of_ad_of_hots"`
}

func FetchIdListOfADOfHots(ctx context.Context, req *advertisement.FetchIdListOfADOfHotsReq, rsp *advertisement.FetchIdListOfADOfHotsRsp) error {
	version, err := version_of_ad_of_hots.GetMaxId()
	if err != nil {
		log.Error("version_of_ad_of_hots.GetMaxId failure, err: ", err)
		rsp.Code = code.DatabaseFailure
		return nil
	}
	model, err := ad_of_hots.GetModelByVersion(version)
	if err != nil {
		log.Error("ad_of_hots.GetLatestVersionModel failure, err: ", err)
		rsp.Code = code.DatabaseFailure
		return nil
	}

	output := &OutputOfIDListOfADOfHots{
		VersionOfADOfHots: model.Id,
		IdListOfADOfHots:  []int64{},
	}

	if len(model.AdvertisementIdList) > 0 {
		idList := []int64{}
		err = json.Unmarshal([]byte(model.AdvertisementIdList), &idList)
		if err != nil {
			log.Error("json.Unmarshal failure, err: ", err)
			rsp.Code = code.InternalError
			return nil
		}
		if req.Behavior > 0 {
			idList, err = advertisement2.GetIDListByIDListWithoutStatus(idList)
			if err != nil {
				log.Error("advertisement2.GetIDListByIDListWithoutStatus failure, err: ", err)
				rsp.Code = code.DatabaseFailure
				return nil
			}
		} else {
			idList, err = advertisement2.GetIDListByIDListWithStatus(idList)
			if err != nil {
				log.Error("advertisement2.GetIDListByIDListWithStatus failure, err: ", err)
				rsp.Code = code.DatabaseFailure
				return nil
			}
		}
		output.IdListOfADOfHots = idList
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
