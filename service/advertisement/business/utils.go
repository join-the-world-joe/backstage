package business

import "backstage/common/db/mysql/backend/selling_point_of_advertisement"

func getSellingPointByAdvertisementId(advertisementId int64, spl []*selling_point_of_advertisement.Model) [][]byte {
	spList := [][]byte{}
	for _, v := range spl {
		if v.AdvertisementId == advertisementId {
			spList = append(spList, []byte(v.SellingPoint))
		}
	}
	return spList
}
