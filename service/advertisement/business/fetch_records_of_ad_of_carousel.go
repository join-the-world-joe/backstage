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

type OutputOfRecordsOfADOfCarousel struct {
	RecordsOfADOfCarousel []*Item `json:"records_of_ad_of_carousel"`
}

func FetchRecordsOfADOfCarousel(ctx context.Context, req *advertisement.FetchRecordsOfADOfCarouselReq, rsp *advertisement.FetchRecordsOfADOfCarouselRsp) error {
	output := &OutputOfRecordsOfADOfCarousel{}

	if len(req.AdvertisementIdList) > 0 {
		aml, err := advertisement2.GetModelListByIdList(req.AdvertisementIdList)
		if err != nil {
			log.Error("FetchRecordsOfADOfCarousel.advertisement2.GetModelListByIdList failure, err: ", err)
			rsp.Code = code.DatabaseFailure
			return nil
		}

		if len(aml) > 0 {
			spml, err := selling_point_of_advertisement.GetModelListByAdvertisementIdList(req.AdvertisementIdList)
			if err != nil {
				log.Error("FetchRecordsOfADOfCarousel.selling_point_of_advertisement.GetModelListByAdvertisementIdList failure, err: ", err)
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
					ImagePath:         temp.Url,
					PlaceOfOrigin:     temp.PlaceOFOrigin,
					Description:       temp.Description,
					AdvertisementId:   temp.Id,
					AdvertisementName: temp.Name,
					SellingPoints:     getSellingPointByAdvertisementId(temp.Id, spml),
				}
				itemHash[temp.ProductId] = append(itemHash[temp.ProductId], item)
				output.RecordsOfADOfCarousel = append(output.RecordsOfADOfCarousel, item)
				if !slices.Contains(productIdList, temp.Id) {
					productIdList = append(productIdList, temp.ProductId)
				}
			}
			pml, err := product.GetModelListByIdList(productIdList)
			if err != nil {
				log.Error("FetchRecordsOfADOfCarousel.product.GetModelListByIdList failure, err: ", err)
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
		log.Error("FetchRecordsOfADOfCarousel.json.Marshal failure, err: ", err)
		rsp.Code = code.InvalidData
		return nil
	}
	rsp.Body = bytes
	rsp.Code = code.Success
	return nil
}

func getSellingPointByAdvertisementId(advertisementId int64, spl []*selling_point_of_advertisement.Model) [][]byte {
	spList := [][]byte{}
	for _, v := range spl {
		if v.AdvertisementId == advertisementId {
			spList = append(spList, []byte(v.SellingPoint))
		}
	}
	return spList
}
