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
	case admin.InsertRecordOfUserReq_:
		insertRecordOfUser(packet)
	case admin.SoftDeleteRecordOfUserReq_:
		softDeleteRecordOfUser(packet)
	case admin.UpdateRecordOfUserReq_:
		updateRecordOfUser(packet)
	case admin.FetchFieldListOfConditionReq_:
		fetchFieldListOfCondition(packet)
	case admin.FetchTrackListOfConditionReq_:
		fetchTrackListOfCondition(packet)
	case admin.CheckPermissionReq_:
		checkPermission(packet)
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
	case admin.InsertRecordOfADOfDealsReq_:
		insertRecordOfADOfDeals(packet)
	case admin.RemoveOutdatedRecordsOfADOfDealsReq_:
		removeOutdatedRecordsOfADOfDeals(packet)
	case admin.InsertRecordOfADOfBarbecueReq_:
		insertRecordOfADOfBarbecue(packet)
	case admin.RemoveOutdatedRecordsOfADOfBarbecueReq_:
		removeOutdatedRecordsOfADOfBarbecue(packet)
	case admin.InsertRecordOfADOfCampingReq_:
		insertRecordOfADOfCamping(packet)
	case admin.RemoveOutdatedRecordsOfADOfCampingReq_:
		removeOutdatedRecordsOfADOfCamping(packet)
	case admin.InsertRecordOfADOfSnacksReq_:
		insertRecordOfADOfSnacks(packet)
	case admin.RemoveOutdatedRecordsOfADOfSnacksReq_:
		removeOutdatedRecordsOfADOfSnacks(packet)
	case admin.InsertRecordOfProductReq_:
		insertRecordOfProduct(packet)
	case admin.SoftDeleteRecordsOfProductReq_:
		softDeleteRecordsOfProduct(packet)
	case admin.UpdateRecordOfProductReq_:
		updateRecordOfProduct(packet)
	default:
		log.ErrorF("unknown minor [%v]", packet.GetRequest().GetHeader().GetMinor())
	}
}
