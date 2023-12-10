package business

import (
	"backstage/common/code"
	"backstage/common/db/mysql/backend/ad_of_carousel"
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
	model, err := ad_of_carousel.GetLatestVersionModel()
	if err != nil {
		log.Error("FetchIdListOfADOfCarousel.ad_of_carousel.GetLatestVersionModel failure, err: ", err)
		rsp.Code = code.DatabaseFailure
		return nil
	}
	idList := []int64{}
	err = json.Unmarshal([]byte(model.AdvertisementIdList), &idList)
	if err != nil {
		log.Error("FetchIdListOfADOfCarousel.json.Unmarshal failure, err: ", err)
		rsp.Code = code.InternalError
		return nil
	}
	output := &OutputOfIDListOfADOfCarousel{
		VersionOfADOfCarousel: model.Id,
		IdListOfADOfCarousel:  idList,
	}

	bytes, err := json.Marshal(output)
	if err != nil {
		log.Error("FetchIdListOfADOfCarousel.json.Marshal failure, err: ", err)
		rsp.Code = code.InvalidData
		return nil
	}
	rsp.Body = bytes
	rsp.Code = code.Success
	return nil
}
