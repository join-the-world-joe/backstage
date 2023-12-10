package dispatch

//
//import (
//	"backstage/common/major"
//	"backstage/common/payload"
//	"backstage/common/protocol/advertisement"
//	"backstage/common/route"
//	"backstage/global/config"
//	"backstage/global/log"
//	"backstage/service/advertisement/business"
//	"context"
//	"encoding/json"
//)
//
//func fetchADOfHotDeals(packet *payload.PacketInternal) {
//	req := &advertisement.FetchADOfHotDealsReq{}
//	rsp := &advertisement.FetchADOfHotDealsRsp{}
//
//	err := json.Unmarshal(packet.GetRequest().GetBody(), req)
//	if err != nil {
//		log.Error("Dispatch.fetchADOfHotDeals.json.Unmarshal failure, err: ", err.Error())
//		return
//	}
//
//	req.UserId = packet.GetSession().GetUserId()
//
//	err = business.FetchADOfHotDeals(context.Background(), req, rsp)
//	if err != nil {
//		log.Error("Dispatch.fetchADOfHotDeals.business.FetchADOfHotDeals failure, err: ", err.Error())
//		return
//	}
//
//	bytes, err := json.Marshal(rsp)
//	if err != nil {
//		log.Error("Dispatch.fetchADOfHotDeals.json.Marshal failure, err: ", err.Error())
//		return
//	}
//
//	packet.Response = &payload.PacketClient{
//		Header: &payload.Header{
//			Major: major.Advertisement,
//			Minor: advertisement.FetchADOfHotDealsRsp_,
//		},
//		Body: bytes,
//	}
//
//	err = route.Downstream(
//		config.DownstreamProtocol(),
//		packet,
//	)
//	if err != nil {
//		log.Error("Dispatch.fetchADOfDealsOfToday.route.Downstream failure, err: ", err.Error())
//		return
//	}
//}
