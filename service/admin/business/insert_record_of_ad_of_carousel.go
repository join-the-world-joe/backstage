package business

import (
	"backstage/common/code"
	"backstage/common/db/mysql/backend/ad_of_carousel"
	"backstage/common/db/mysql/backend/version_of_ad_of_carousel"
	"backstage/common/major"
	"backstage/common/protocol/admin"
	"backstage/global/log"
	"context"
	"encoding/json"
	"github.com/spf13/cast"
)

func InsertRecordOfADOfCarousel(ctx context.Context, req *admin.InsertRecordOfADOfCarouselReq, rsp *admin.InsertRecordOfADOfCarouselRsp) error {
	if !hasPermission(
		cast.ToInt(major.Admin),
		cast.ToInt(admin.InsertRecordOfADOfCarouselReq_),
		req.UserId,
	) {
		rsp.Code = code.AccessDenied
		return nil
	}

	version, err := version_of_ad_of_carousel.InsertModel(&version_of_ad_of_carousel.Model{})
	if err != nil {
		log.Error("version_of_ad_of_carousel.InsertModel fail, err: ", err)
		rsp.Code = code.DatabaseFailure
		return nil
	}

	bytes, err := json.Marshal(&req.AdvertisementIdList)
	if err != nil {
		log.Error("json.Marshal fail, err: ", err)
		rsp.Code = code.InternalError
		return nil
	}

	_, err = ad_of_carousel.InsertModel(&ad_of_carousel.Model{
		Version:             version.Id,
		AdvertisementIdList: string(bytes),
	})
	if err != nil {
		rsp.Code = code.DatabaseFailure
		return nil
	}

	rsp.Code = code.Success
	return nil
}
