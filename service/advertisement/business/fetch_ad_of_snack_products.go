package business

import (
	"backstage/common/code"
	"backstage/common/db/mysql/backend/ad_of_snack_products"
	advertisement2 "backstage/common/db/mysql/backend/advertisement"
	"backstage/common/protocol/advertisement"
	"context"
	"encoding/json"
	"github.com/spf13/cast"
)

type OutputOfADOfSnackProducts struct {
	SnackProducts []*Item `json:"snack_products"`
}

func FetchADOfSnackProducts(ctx context.Context, req *advertisement.FetchADOfSnackProductsReq, rsp *advertisement.FetchADOfSnackProductsRsp) error {
	model, err := ad_of_snack_products.GetLatestVersionModel()
	if err != nil {
		rsp.Code = code.DatabaseFailure
		return nil
	}
	idList := []int64{}
	err = json.Unmarshal([]byte(model.ADIdList), &idList)
	if err != nil {
		rsp.Code = code.InternalError
		return nil
	}
	output := &OutputOfADOfSnackProducts{}
	for _, v := range idList {
		temp, err := advertisement2.GetModelById(v)
		if err != nil {
			continue
		}
		output.SnackProducts = append(output.SnackProducts, &Item{
			Title: temp.Title,
			Stock: temp.Stock,
			//Price:     temp.SellingPrice,
			ProductId: temp.ProductId,
			ImagePath: temp.Url,
			//SellingPoint:  temp.,
			PlaceOfOrigin: temp.PlaceOFOrigin,
		})
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
