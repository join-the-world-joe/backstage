package business

import (
	"backstage/common/code"
	"backstage/common/db/mysql/backend/version_of_ad_of_carousel"
	"backstage/common/protocol/advertisement"
	"context"
	"encoding/json"
)

type OutputOfVersionOfADOfCarousel struct {
	VersionOfADOfCarousel int64 `json:"version_of_ad_of_carousel"`
}

func FetchVersionOfADOfCarousel(ctx context.Context, req *advertisement.FetchVersionOfADOfCarouselReq, rsp *advertisement.FetchVersionOfADOfCarouselRsp) error {
	version, err := version_of_ad_of_carousel.GetMaxId()
	if err != nil {
		rsp.Code = code.DatabaseFailure
		return nil
	}

	output := &OutputOfVersionOfADOfCarousel{
		VersionOfADOfCarousel: version,
	}

	bytes, err := json.Marshal(output)
	if err != nil {
		rsp.Code = code.InternalError
		return nil
	}

	rsp.Body = bytes
	rsp.Code = code.Success
	return nil
}
