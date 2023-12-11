package business

import (
	"backstage/common/code"
	"backstage/common/db/mysql/backend/ad_of_hots"
	"backstage/common/db/mysql/backend/version_of_ad_of_hots"
	"backstage/common/major"
	"backstage/common/protocol/admin"
	"backstage/global/log"
	"context"
	"encoding/json"
	"github.com/spf13/cast"
)

func InsertRecordOfADOfHots(ctx context.Context, req *admin.InsertRecordOfADOfHotsReq, rsp *admin.InsertRecordOfADOfHotsRsp) error {
	if !hasPermission(
		cast.ToInt(major.Admin),
		cast.ToInt(admin.InsertRecordOfADOfHotsReq_),
		req.UserId,
	) {
		rsp.Code = code.AccessDenied
		return nil
	}

	version, err := version_of_ad_of_hots.InsertModel(&version_of_ad_of_hots.Model{})
	if err != nil {
		log.Error("version_of_ad_of_carousel.GetMaxId fail, err: ", err)
		rsp.Code = code.DatabaseFailure
		return nil
	}

	bytes, err := json.Marshal(&req.AdvertisementIdList)
	if err != nil {
		log.Error("json.Marshal fail, err: ", err)
		rsp.Code = code.InternalError
		return nil
	}

	_, err = ad_of_hots.InsertModel(&ad_of_hots.Model{
		Version:             version.Id,
		AdvertisementIdList: string(bytes),
		Description:         req.Description,
	})
	if err != nil {
		rsp.Code = code.DatabaseFailure
		return nil
	}

	rsp.Code = code.Success
	return nil
}
