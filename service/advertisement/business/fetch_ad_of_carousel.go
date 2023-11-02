package business

import (
	"backstage/common/code"
	"backstage/common/db/mysql/backend/ad_of_carousel"
	"backstage/common/protocol/advertisement"
	"context"
	"encoding/json"
	"github.com/spf13/cast"
)

type OutputOfADOfCarousel struct {
	ImagePathList string `json:"image_path_list"`
}

func FetchADOfCarousel(ctx context.Context, req *advertisement.FetchADOfCarouselReq, rsp *advertisement.FetchADOfCarouselRsp) error {
	model, err := ad_of_carousel.GetLatestVersionModel()
	if err != nil {
		rsp.Code = code.DatabaseFailure
		return nil
	}

	output := &OutputOfADOfCarousel{
		ImagePathList: model.ImagePathList,
	}
	bytes, err := json.Marshal(output)
	if err != nil {
		rsp.Code = code.InvalidData
		return nil
	}

	rsp.Body = bytes
	rsp.Code = code.Success
	rsp.Version = cast.ToInt(model.Id)
	return nil
}
