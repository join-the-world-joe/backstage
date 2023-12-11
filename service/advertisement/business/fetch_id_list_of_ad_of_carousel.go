package business

import (
	"backstage/common/code"
	"backstage/common/db/mysql/backend/ad_of_carousel"
	"backstage/common/db/mysql/backend/version_of_ad_of_carousel"
	"backstage/common/protocol/advertisement"
	"backstage/global/log"
	"context"
	"encoding/json"
)

type OutputOfIDListOfADOfCarousel struct {
	VersionOfADOfCarousel int64   `json:"version_of_ad_of_carousel"`
	IdListOfADOfCarousel  []int64 `json:"id_list_of_ad_of_carousel"`
}

func FetchIdListOfADOfCarousel(ctx context.Context, req *advertisement.FetchIdListOfADOfCarouselReq, rsp *advertisement.FetchIdListOfADOfCarouselRsp) error {
	version, err := version_of_ad_of_carousel.GetMaxId()
	if err != nil {
		log.Error("version_of_ad_of_carousel.GetMaxId failure, err: ", err)
		rsp.Code = code.DatabaseFailure
		return nil
	}

	model, err := ad_of_carousel.GetModelByVersion(version)
	if err != nil {
		log.Error("ad_of_carousel.GetModelByVersion failure, err: ", err)
		rsp.Code = code.DatabaseFailure
		return nil
	}
	output := &OutputOfIDListOfADOfCarousel{
		VersionOfADOfCarousel: model.Version,
		IdListOfADOfCarousel:  []int64{},
	}

	if len(model.AdvertisementIdList) > 0 {
		idList := []int64{}
		err = json.Unmarshal([]byte(model.AdvertisementIdList), &idList)
		if err != nil {
			log.Error("json.Unmarshal failure, err: ", err)
			rsp.Code = code.InternalError
			return nil
		}
		output.IdListOfADOfCarousel = idList
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
