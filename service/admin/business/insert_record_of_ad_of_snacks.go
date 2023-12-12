package business

import (
	"backstage/common/code"
	"backstage/common/db/mysql/backend/ad_of_snacks"
	"backstage/common/db/mysql/backend/version_of_ad_of_snacks"
	"backstage/common/major"
	"backstage/common/protocol/admin"
	"backstage/global/log"
	"context"
	"encoding/json"
	"github.com/spf13/cast"
)

func InsertRecordOfADOfSnacks(ctx context.Context, req *admin.InsertRecordOfADOfSnacksReq, rsp *admin.InsertRecordOfADOfSnacksRsp) error {
	if !hasPermission(
		cast.ToInt(major.Admin),
		cast.ToInt(admin.InsertRecordOfADOfSnacksReq_),
		req.UserId,
	) {
		rsp.Code = code.AccessDenied
		return nil
	}

	version, err := version_of_ad_of_snacks.InsertModel(&version_of_ad_of_snacks.Model{})
	if err != nil {
		log.Error("version_of_ad_of_snacks.InsertModel fail, err: ", err)
		rsp.Code = code.DatabaseFailure
		return nil
	}

	bytes, err := json.Marshal(&req.AdvertisementIdList)
	if err != nil {
		log.Error("json.Marshal fail, err: ", err)
		rsp.Code = code.InternalError
		return nil
	}

	_, err = ad_of_snacks.InsertModel(&ad_of_snacks.Model{
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
