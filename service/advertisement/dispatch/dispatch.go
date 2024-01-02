package dispatch

import (
	"backstage/common/payload"
	"backstage/common/protocol/advertisement"
	"backstage/global/log"
)

func Dispatch(packet *payload.PacketInternal) {
	switch packet.GetRequest().GetHeader().GetMinor() {
	case advertisement.FetchVersionOfADOfCarouselReq_:
		fetchVersionOfADOfCarousel(packet)
	case advertisement.FetchIdListOfADOfCarouselReq_:
		fetchIdListOfADOfCarousel(packet)
	case advertisement.FetchRecordsOfADOfCarouselReq_:
		fetchRecordsOfADOfCarousel(packet)
	case advertisement.FetchVersionOfADOfDealsReq_:
		fetchVersionOfADOfDeals(packet)
	case advertisement.FetchIdListOfADOfDealsReq_:
		fetchIdListOfADOfDeals(packet)
	case advertisement.FetchRecordsOfADOfDealsReq_:
		fetchRecordsOfADOfDeals(packet)
	case advertisement.FetchVersionOfADOfCampingReq_:
		fetchVersionOfADOfCamping(packet)
	case advertisement.FetchIdListOfADOfCampingReq_:
		fetchIdListOfADOfCamping(packet)
	case advertisement.FetchRecordsOfADOfCampingReq_:
		fetchRecordsOfADOfCamping(packet)
	case advertisement.FetchVersionOfADOfBarbecueReq_:
		fetchVersionOfADOfBarbecue(packet)
	case advertisement.FetchIdListOfADOfBarbecueReq_:
		fetchIdListOfADOfBarbecue(packet)
	case advertisement.FetchRecordsOfADOfBarbecueReq_:
		fetchRecordsOfADOfBarbecue(packet)
	case advertisement.FetchVersionOfADOfSnacksReq_:
		fetchVersionOfADOfSnacks(packet)
	case advertisement.FetchIdListOfADOfSnacksReq_:
		fetchIdListOfADOfSnacks(packet)
	case advertisement.FetchRecordsOfADOfSnacksReq_:
		fetchRecordsOfADOfSnacks(packet)
	case advertisement.FetchIdListOfAdvertisementReq_:
		fetchIdListOfAdvertisement(packet)
	case advertisement.FetchRecordsOfAdvertisementReq_:
		fetchRecordsOfAdvertisement(packet)
	default:
		log.ErrorF("unknown minor [%v]", packet.GetRequest().GetHeader().GetMinor())
	}
}
