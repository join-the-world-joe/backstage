package dispatch

import (
	"backstage/common/payload"
	"backstage/common/protocol/product"
	"backstage/global/log"
)

func Dispatch(packet *payload.PacketInternal) {
	switch packet.GetRequest().GetHeader().GetMinor() {
	case product.FetchIdListOfProductReq_:
		fetchIdListOfProduct(packet)
	case product.FetchRecordsOfProductReq_:
		fetchRecordsOfProduct(packet)
	default:
		log.ErrorF("unknown minor [%v]", packet.GetRequest().GetHeader().GetMinor())
	}
}
