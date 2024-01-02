package business

import (
	"backstage/common/code"
	"backstage/common/db/mysql/backend/advertisement"
	"backstage/common/db/mysql/backend/selling_point_of_advertisement"
	"backstage/common/macro/timestamp"
	advertisement2 "backstage/common/protocol/advertisement"
	"backstage/global/log"
	"context"
	"encoding/json"
)

type RecordOfAdvertisement struct {
	Id            int64    `json:"id"`
	Name          string   `json:"name"`
	Title         string   `json:"title"`
	SellingPrice  int      `json:"selling_price"`
	PlaceOfOrigin string   `json:"place_of_origin"`
	SellingPoints [][]byte `json:"selling_points"`
	CoverImage    string   `json:"cover_image"`
	FirstImage    string   `json:"first_image"`
	SecondImage   string   `json:"second_image"`
	ThirdImage    string   `json:"third_image"`
	FourthImage   string   `json:"fourth_image"`
	FifthImage    string   `json:"fifth_image"`
	Stock         int      `json:"stock"`
	Status        int      `json:"status"`
	ProductId     int64    `json:"product_id"`
	BuyingPrice   int      `json:"buying_price"`
	CreatedAt     string   `json:"created_at"`
	UpdatedAt     string   `json:"updated_at"`
}

type OutputOfRecordsOfAdvertisement struct {
	RecordsOfAdvertisement map[int64]*RecordOfAdvertisement `json:"records_of_advertisement"`
}

func FetchRecordsOfAdvertisement(ctx context.Context, req *advertisement2.FetchRecordsOfAdvertisementReq, rsp *advertisement2.FetchRecordsOfAdvertisementRsp) error {
	output := &OutputOfRecordsOfAdvertisement{
		RecordsOfAdvertisement: map[int64]*RecordOfAdvertisement{},
	}

	ml, err := advertisement.GetModelListByIdList(req.AdvertisementIdList)
	if err != nil {
		rsp.Code = code.DatabaseFailure
		return nil
	}
	for _, m := range ml {
		pml, err := selling_point_of_advertisement.GetModelListByAdvertisementId(m.Id)
		if err != nil {
			log.Error("business.selling_point_of_advertisement.GetModelListByAdvertisementId failure, err: ", err.Error())
			continue
		}
		points := [][]byte{}
		for _, v := range pml {
			//fmt.Println("id: ", m.Id, ", selling_point: ", v.SellingPoint)
			points = append(points, []byte(v.SellingPoint))
		}
		output.RecordsOfAdvertisement[m.Id] = &RecordOfAdvertisement{
			Id:            m.Id,
			Name:          m.Name,
			Title:         m.Title,
			SellingPrice:  m.SellingPrice,
			PlaceOfOrigin: m.PlaceOFOrigin,
			SellingPoints: points,
			CoverImage:    m.CoverImage,
			FirstImage:    m.FirstImage,
			SecondImage:   m.SecondImage,
			ThirdImage:    m.ThirdImage,
			FourthImage:   m.FourthImage,
			FifthImage:    m.FifthImage,
			Stock:         m.Stock,
			Status:        m.Status,
			ProductId:     m.ProductId,
			CreatedAt:     m.CreatedAt.Format(timestamp.YYMDHMS),
			UpdatedAt:     m.UpdatedAt.Format(timestamp.YYMDHMS),
		}
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
