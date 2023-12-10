package business

import (
	"backstage/common/code"
	"backstage/common/db/mysql/backend/ad_of_carousel"
	"backstage/common/major"
	"backstage/common/protocol/admin"
	"context"
	"github.com/spf13/cast"
)

func RemoveOutdatedRecordsOfADOfCarousel(ctx context.Context, req *admin.RemoveOutdatedRecordsOfADOfCarouselReq, rsp *admin.RemoveOutdatedRecordsOfADOfCarouselRsp) error {
	if !hasPermission(
		cast.ToInt(major.Admin),
		cast.ToInt(admin.RemoveOutdatedRecordsOfADOfCarouselReq_),
		req.UserId,
	) {
		rsp.Code = code.AccessDenied
		return nil
	}

	ad_of_carousel.RemoveOutdatedRecordsOfAdCarousel()

	rsp.Code = code.Success
	return nil
}
