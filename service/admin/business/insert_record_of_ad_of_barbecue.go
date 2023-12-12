package business

import (
	"backstage/common/code"
	"backstage/common/db/mysql/backend/ad_of_barbecue"
	"backstage/common/db/mysql/backend/version_of_ad_of_barbecue"
	"backstage/common/major"
	"backstage/common/protocol/admin"
	"backstage/global/log"
	"context"
	"encoding/json"
	"github.com/spf13/cast"
)

func InsertRecordOfADOfBarbecue(ctx context.Context, req *admin.InsertRecordOfADOfBarbecueReq, rsp *admin.InsertRecordOfADOfBarbecueRsp) error {
	if !hasPermission(
		cast.ToInt(major.Admin),
		cast.ToInt(admin.InsertRecordOfADOfBarbecueReq_),
		req.UserId,
	) {
		rsp.Code = code.AccessDenied
		return nil
	}

	version, err := version_of_ad_of_barbecue.InsertModel(&version_of_ad_of_barbecue.Model{})
	if err != nil {
		log.Error("version_of_ad_of_barbecue.InsertModel fail, err: ", err)
		rsp.Code = code.DatabaseFailure
		return nil
	}

	bytes, err := json.Marshal(&req.AdvertisementIdList)
	if err != nil {
		log.Error("json.Marshal fail, err: ", err)
		rsp.Code = code.InternalError
		return nil
	}

	_, err = ad_of_barbecue.InsertModel(&ad_of_barbecue.Model{
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
