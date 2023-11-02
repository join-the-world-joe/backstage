package business

import (
	"backstage/common/code"
	"backstage/common/db/mysql/backend/version_of_ad_of_carousel"
	"backstage/common/protocol/advertisement"
	"context"
	"github.com/spf13/cast"
)

func FetchVersionOfADOfCarousel(ctx context.Context, req *advertisement.FetchVersionOfADOfCarouselReq, rsp *advertisement.FetchVersionOfADOfCarouselRsp) error {
	id, err := version_of_ad_of_carousel.GetMaxId()
	if err != nil {
		rsp.Code = code.DatabaseFailure
		return nil
	}
	rsp.Code = code.Success
	rsp.Version = cast.ToInt(id)
	return nil
}
