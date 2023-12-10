package dispatch

import (
	"backstage/common/payload"
	"backstage/common/protocol/admin"
	"backstage/global/log"
)

func Dispatch(packet *payload.PacketInternal) {
	switch packet.GetRequest().GetHeader().GetMinor() {
	case admin.FetchMenuListOfConditionReq_:
		fetchMenuListOfCondition(packet)
	case admin.FetchUserListOfConditionReq_:
		fetchUserListOfCondition(packet)
	case admin.FetchRoleListOfConditionReq_:
		fetchRoleListOfCondition(packet)
	case admin.FetchPermissionListOfConditionReq_:
		fetchPermissionListOfCondition(packet)
	case admin.InsertUserRecordReq_:
		insertUserRecord(packet)
	case admin.SoftDeleteUserRecordReq_:
		softDeleteUserRecord(packet)
	case admin.UpdateUserRecordReq_:
		updateUserRecord(packet)
	case admin.FetchFieldListOfConditionReq_:
		fetchFieldListOfCondition(packet)
	case admin.FetchTrackListOfConditionReq_:
		fetchTrackListOfCondition(packet)
	case admin.CheckPermissionReq_:
		checkPermission(packet)
	case admin.FetchIdListOfGoodReq_:
		fetchIdListOfGood(packet)
	case admin.FetchRecordsOfGoodReq_:
		fetchRecordsOfGood(packet)
	case admin.InsertRecordOfGoodReq_:
		insertRecordOfGood(packet)
	case admin.SoftDeleteRecordsOfGoodReq_:
		softDeleteRecordsOfGood(packet)
	case admin.UpdateRecordOfGoodReq_:
		updateRecordOfGood(packet)
	case admin.FetchIdListOfAdvertisementReq_:
		fetchIdListOfAdvertisement(packet)
	case admin.FetchRecordsOfAdvertisementReq_:
		fetchRecordsOfAdvertisement(packet)
	case admin.InsertRecordOfAdvertisementReq_:
		insertRecordOfAdvertisement(packet)
	case admin.SoftDeleteRecordsOfAdvertisementReq_:
		softDeleteRecordsOfAdvertisement(packet)
	case admin.UpdateRecordOfAdvertisementReq_:
		updateRecordOfAdvertisement(packet)
	case admin.InsertRecordOfADOfCarouselReq_:
		insertRecordOfADOfCarousel(packet)
	case admin.RemoveOutdatedRecordsOfADOfCarouselReq_:
		removeOutdatedRecordsOfADOfCarousel(packet)
	default:
		log.ErrorF("unknown minor [%v]", packet.GetRequest().GetHeader().GetMinor())
	}
}
