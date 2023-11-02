package admin

type SoftDeleteRecordsOfAdvertisementReq struct {
	UserId              int64
	AdvertisementIdList []int64 `json:"advertisement_id_list"`
}

type SoftDeleteRecordsOfAdvertisementRsp struct {
	Code int `json:"code"`
}
