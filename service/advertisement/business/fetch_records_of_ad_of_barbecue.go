package business

import (
	"backstage/common/code"
	advertisement2 "backstage/common/db/mysql/backend/advertisement"
	"backstage/common/db/mysql/backend/product"
	"backstage/common/db/mysql/backend/selling_point_of_advertisement"
	"backstage/common/protocol/advertisement"
	"backstage/global/log"
	"context"
	"encoding/json"
	"golang.org/x/exp/slices"
)

type OutputOfRecordsOfADOfBarbecue struct {
	RecordsOfADOfBarbecue []*Item `json:"records_of_ad_of_barbecue"`
}

func FetchRecordsOfADOfBarbecue(ctx context.Context, req *advertisement.FetchRecordsOfADOfBarbecueReq, rsp *advertisement.FetchRecordsOfADOfBarbecueRsp) error {
	output := &OutputOfRecordsOfADOfBarbecue{}

	if len(req.AdvertisementIdList) > 0 {
		aml, err := advertisement2.GetModelListByIdList(req.AdvertisementIdList)
		if err != nil {
			log.Error("advertisement2.GetModelListByIdList failure, err: ", err)
			rsp.Code = code.DatabaseFailure
			return nil
		}

		if len(aml) > 0 {
			spml, err := selling_point_of_advertisement.GetModelListByAdvertisementIdList(req.AdvertisementIdList)
			if err != nil {
				log.Error("selling_point_of_advertisement.GetModelListByAdvertisementIdList failure, err: ", err)
				rsp.Code = code.DatabaseFailure
				return nil
			}
			itemHash := map[int64][]*Item{} // key: product_id
			productIdList := []int64{}
			for _, ad := range aml {
				item := &Item{
					Title:             ad.Title,
					Stock:             ad.Stock,
					SellingPrice:      ad.SellingPrice,
					ProductId:         ad.ProductId,
					CoverImage:        ad.CoverImage,
					FirstImage:        ad.FirstImage,
					SecondImage:       ad.SecondImage,
					ThirdImage:        ad.ThirdImage,
					FourthImage:       ad.FourthImage,
					FifthImage:        ad.FifthImage,
					PlaceOfOrigin:     ad.PlaceOFOrigin,
					OSSPath:           ad.OSSPath,
					Status:            ad.Status,
					AdvertisementId:   ad.Id,
					AdvertisementName: ad.Name,
					SellingPoints:     getSellingPointByAdvertisementId(ad.Id, spml),
				}
				itemHash[ad.ProductId] = append(itemHash[ad.ProductId], item)
				output.RecordsOfADOfBarbecue = append(output.RecordsOfADOfBarbecue, item)
				if !slices.Contains(productIdList, ad.Id) {
					productIdList = append(productIdList, ad.ProductId)
				}
			}
			pml, err := product.GetModelListByIdList(productIdList)
			if err != nil {
				log.Error("product.GetModelListByIdList failure, err: ", err)
				rsp.Code = code.InternalError
				return nil
			}
			for _, v := range pml {
				if items, exist := itemHash[v.Id]; exist {
					for _, item := range items {
						item.ProductName = v.Name
					}
				}
			}
		}
	}

	bytes, err := json.Marshal(output)
	if err != nil {
		log.Error("json.Marshal failure, err: ", err)
		rsp.Code = code.InvalidData
		return nil
	}
	rsp.Body = bytes
	rsp.Code = code.Success
	return nil
}
