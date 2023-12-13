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

type OutputOfRecordsOfADOfHots struct {
	RecordsOfADOfHots []*Item `json:"records_of_ad_of_hots"`
}

func FetchRecordsOfADOfHots(ctx context.Context, req *advertisement.FetchRecordsOfADOfHotsReq, rsp *advertisement.FetchRecordsOfADOfHotsRsp) error {
	output := &OutputOfRecordsOfADOfHots{}

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
			for _, temp := range aml {
				item := &Item{
					Title:             temp.Title,
					Stock:             temp.Stock,
					SellingPrice:      temp.SellingPrice,
					ProductId:         temp.ProductId,
					Image:             temp.Image,
					PlaceOfOrigin:     temp.PlaceOFOrigin,
					AdvertisementId:   temp.Id,
					AdvertisementName: temp.Name,
					SellingPoints:     getSellingPointByAdvertisementId(temp.Id, spml),
				}
				itemHash[temp.ProductId] = append(itemHash[temp.ProductId], item)
				output.RecordsOfADOfHots = append(output.RecordsOfADOfHots, item)
				if !slices.Contains(productIdList, temp.Id) {
					productIdList = append(productIdList, temp.ProductId)
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
